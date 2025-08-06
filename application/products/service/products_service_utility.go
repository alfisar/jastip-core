package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"jastip-core/application/products/repository"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/nfnt/resize"
)

func isImageFile(fileHeader *multipart.FileHeader) bool {
	file, err := fileHeader.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}

	fileType := http.DetectContentType(buffer)
	return strings.HasPrefix(fileType, "image/")
}
func fileHeaderHeaderToBytes(fh *multipart.FileHeader) []byte {
	file, _ := fh.Open()
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.Bytes()
}
func compressImage(file multipart.File) (result bytes.Buffer, err domain.ErrorData) {
	var compressedImgBuffer bytes.Buffer

	img, _, errData := image.Decode(file)
	if errData != nil {
		message := fmt.Sprintf("error decode image on func compressImage : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		return
	}

	resizeImg := resize.Resize(1000, 0, img, resize.NearestNeighbor)
	finalQuality := 0

	errData = imaging.Encode(&compressedImgBuffer, resizeImg, imaging.JPEG)
	if errData != nil {
		message := fmt.Sprintf("error Encode image on func compressImage : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		return
	}

	compressedSize := int(compressedImgBuffer.Len())
	if int64(compressedSize) > 1024*1024 {
		for quality := 100; quality > 10; quality -= 10 {
			var compressedImgBuffers bytes.Buffer
			errData = imaging.Encode(&compressedImgBuffers, resizeImg, imaging.JPEG, imaging.JPEGQuality(quality))
			if errData != nil {
				message := fmt.Sprintf("error Encode image on func compressImage : %s", errData.Error())
				log.Println(message)

				err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
				return
			}

			compressedSize := int(compressedImgBuffers.Len())
			if int64(compressedSize) <= 1024*1024 {
				if finalQuality < quality {
					compressedImgBuffer = compressedImgBuffers
					finalQuality = quality
				}
				break
			}
		}

		if compressedImgBuffer.Len() == 0 {
			message := fmt.Sprintf("error compress image on func compressImage : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrInvalidDataImage, message)
			return
		}
	}
	result = compressedImgBuffer
	return
}

func validationImage(r *multipart.Form) (fileHeader *multipart.FileHeader, compressBuffer bytes.Buffer, err domain.ErrorData) {

	fileHeader = r.File["image"][0]

	file, errData := fileHeader.Open()
	if errData != nil {
		message := fmt.Sprintf("Error open image on func validationImage : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		return
	}

	defer file.Close()
	fileSize := fileHeader.Size

	if fileSize > 10*1024*1024 {
		message := fmt.Sprintf("invalid size image on func validationImage : %s", errorhandler.ErrInvalidDataImage)
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrInvalidDataImage, message)
		return
	} else {

		compressBuffers, errData := compressImage(file)
		if errData.Code != 0 {
			err = errData
			return
		}
		compressBuffer = compressBuffers

	}

	switch fileHeader.Header.Values("Content-Type")[0] {

	case consts.ImageJPG, consts.ImageJPEG, consts.ImagePNG:
	default:
		message := fmt.Sprintf("invalid content image on func validationImage : %s", errorhandler.ErrInvalidDataImage)
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrInvalidDataImage, message)
		return
	}

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(fileHeader.Filename)
	validExtension := false

	for _, allowedExt := range allowedExtensions {
		if strings.EqualFold(ext, allowedExt) {
			validExtension = true
			break
		}
	}

	if !validExtension {
		message := fmt.Sprintf("invalid extention image on func validationImage : %s", errorhandler.ErrInvalidDataImage)
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrInvalidDataImage, message)
		return
	}

	isImage := isImageFile(fileHeader)

	if !isImage {
		message := fmt.Sprintf("invalid data image on func validationImage : %s", errorhandler.ErrInvalidDataImage)
		log.Println(message)

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, errorhandler.ErrInvalidDataImage, message)
		return
	}

	return
}

func saveToMinio(ctx context.Context, poolData *domain.Config, fileHeader multipart.FileHeader, compressBuffer bytes.Buffer, path string, pattern string) (name string, err domain.ErrorData) {
	var reader *bytes.Reader

	file, errData := fileHeader.Open()
	if errData != nil {
		message := fmt.Sprintf("Error open image on func validationImage : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		return
	}

	defer file.Close()

	if compressBuffer.Len() > 0 {
		reader = bytes.NewReader(compressBuffer.Bytes())
	} else {
		// Reset file reader karena sudah dibaca sebelumnya
		_, errData := file.Seek(0, 0)
		if errData != nil {
			message := fmt.Sprintf("invalid seek data image on func validationImage : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
			return
		}

		reader = bytes.NewReader(fileHeaderHeaderToBytes(&fileHeader))
	}

	_, errData = poolData.Minio.Client.PutObject(ctx, poolData.Minio.BucketName, path+pattern, reader, int64(compressBuffer.Len()), minio.PutObjectOptions{ContentType: fileHeader.Header.Values("Content-Type")[0]})

	if errData != nil {
		message := fmt.Sprintf("failed save data to minio on func savetominio : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInternal(errorhandler.ErrCodeInternalServer, errData)
		return
	}
	name = path + pattern
	return
}

func saveProduct(ctx context.Context, poolData *domain.Config, fileHeader multipart.FileHeader, compressBuffer bytes.Buffer, data domain.ProductData, repo repository.ProductsRepositoryContract) (err domain.ErrorData) {
	removeSpace := strings.ReplaceAll(fileHeader.Filename, " ", "")
	timestamp := time.Now().Format("20060102150405")
	pattern := fmt.Sprintf("%s_%s", timestamp, removeSpace)

	data.Image = "products/" + pattern
	conn := poolData.DBSql.Begin()

	data.Status = 1
	errData := repo.Create(conn, data)
	if errData != nil {
		message := fmt.Sprintf("failed createndata product on func saveProduct : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInsertData(errData)
		return
	}

	_, err = saveToMinio(ctx, poolData, fileHeader, compressBuffer, "products/", pattern)
	if err.Code != 0 {
		return
	}

	conn.Commit()
	return
}

func getList(poolData *domain.Config, userID int, param domain.Params, repo repository.ProductsRepositoryContract) (result []domain.ProductResp, currentPage int, limits int, total int64, err domain.ErrorData) {
	var errData error
	pages, offset, limit := handler.CalculateOffsetAndLimit(param.Page, param.Limit)

	where := map[string]any{
		"status":  param.Status,
		"user_id": userID,
	}

	result, total, errData = repo.GetList(poolData.DBSql, param, where, offset, limit)
	if errData != nil {
		message := fmt.Sprintf("Error Get List data on func getList : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			err = errorhandler.ErrGetData(errData)
		}
		return
	}
	currentPage = pages
	limits = limit
	return
}

func getListProductTravel(poolData *domain.Config, userID int, travelID int, param domain.Params, repo repository.ProductsRepositoryContract) (result []domain.ProductResp, currentPage int, limits int, total int64, err domain.ErrorData) {
	var errData error
	pages, offset, limit := handler.CalculateOffsetAndLimit(param.Page, param.Limit)

	where := map[string]any{
		"status":                              param.Status,
		"user_id":                             userID,
		"product_travel.traveler_schedule_id": travelID,
	}

	result, total, errData = repo.GetListProductTravel(poolData.DBSql, param, where, offset, limit)
	if errData != nil {
		message := fmt.Sprintf("Error Get List data on func getList : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			err = errorhandler.ErrGetData(errData)
		}
		return
	}
	currentPage = pages
	limits = limit
	return
}

func updateProducts(ctx context.Context, poolData *domain.Config, repo repository.ProductsRepositoryContract, id int, userID int, updates map[string]any, fileHeader *multipart.FileHeader, compressBuffer bytes.Buffer) (err domain.ErrorData) {
	where := map[string]any{
		"id":      id,
		"user_id": userID,
	}

	dataProduct, errData := repo.Get(poolData.DBSql, where)
	if errData != nil {
		message := fmt.Sprintf("Error update data on func updateProducts : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrGetData(fmt.Errorf(message))
		}
		return
	}

	conn := poolData.DBSql.Begin()
	errData = repo.Update(conn, updates, where)
	if errData != nil {
		message := fmt.Sprintf("Error update data on func updateProducts : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrUpdateData(fmt.Errorf(message))
		}

		conn.Rollback()
		return
	}

	split := strings.Split(dataProduct.Image, "products/")

	_, err = saveToMinio(ctx, poolData, *fileHeader, compressBuffer, "products/", split[1])
	if err.Code != 0 {
		conn.Rollback()
		return
	}

	conn.Commit()
	return
}

func deleteProducts(poolData *domain.Config, repo repository.ProductsRepositoryContract, id int, userID int) (err domain.ErrorData) {

	where := map[string]any{
		"id":      id,
		"user_id": userID,
	}

	errData := repo.Delete(poolData.DBSql, where)
	if errData != nil {
		message := fmt.Sprintf("Error delete data on func deleteProducts : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrUpdateData(fmt.Errorf(message))
		}
	}

	return
}
