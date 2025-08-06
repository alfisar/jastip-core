package service

import (
	"bytes"
	"context"
	"jastip-core/application/products/repository"
	"mime/multipart"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/valyala/fasthttp"
)

type productsService struct {
	repo repository.ProductsRepositoryContract
}

func NewProductsService(repo repository.ProductsRepositoryContract) *productsService {
	return &productsService{
		repo: repo,
	}
}

func (s *productsService) Create(ctx context.Context, poolData *domain.Config, data domain.ProductData, r *fasthttp.Request) (err domain.ErrorData) {
	file, _ := r.MultipartForm()

	fileHeader, compressImg, errData := validationImage(file)
	if errData.Code != 0 {
		return
	}
	err = saveProduct(ctx, poolData, *fileHeader, compressImg, data, s.repo)
	return
}

func (s *productsService) GetList(poolData *domain.Config, userID int, params domain.Params) (totalPage int, currentPage int, total int64, limit int, result []domain.ProductResp, err domain.ErrorData) {
	result, currentPage, limit, total, err = getList(poolData, userID, params, s.repo)
	if err.Code != 0 {
		return
	}

	totalPage = int(handler.CalculateTotalPages(total, int64(params.Limit)))
	return
}

func (s *productsService) GetListProductTravel(poolData *domain.Config, userID int, travelID int, params domain.Params) (totalPage int, currentPage int, total int64, limit int, result []domain.ProductResp, err domain.ErrorData) {
	result, currentPage, limit, total, err = getListProductTravel(poolData, userID, travelID, params, s.repo)
	if err.Code != 0 {
		return
	}

	totalPage = int(handler.CalculateTotalPages(total, int64(params.Limit)))
	return
}

func (s *productsService) Update(ctx context.Context, poolData *domain.Config, id int, userID int, update map[string]any, file *multipart.Form) (err domain.ErrorData) {
	var (
		fileHeader  *multipart.FileHeader
		compressImg bytes.Buffer
	)
	if file != nil {
		fileHeader, compressImg, err = validationImage(file)
		if err.Code != 0 {
			return
		}
	}

	err = updateProducts(ctx, poolData, s.repo, id, userID, update, fileHeader, compressImg)
	return
}

func (s *productsService) Delete(poolData *domain.Config, id int, userId int) (err domain.ErrorData) {
	err = deleteProducts(poolData, s.repo, id, userId)
	return
}
