package service

import (
	"context"
	"log"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

func checkData(ctx context.Context, dataOrderService orderService, poolData *domain.Config, userID int, data domain.OrderRequest) (price float32, addrResult domain.AddressResponse, productResult []domain.OrderDetail, err domain.ErrorData) {
	var wg sync.WaitGroup

	errChan := make(chan domain.ErrorData, 3)
	wg.Add(3)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if userID != data.BuyerID {
		log.Printf("invalid data buyer on func checkData order : buyer different with user login \n")
		err = errorhandler.ErrValidation(nil)
		return
	}

	go func() {
		defer wg.Done()
		where := map[string]any{
			"user_id":              userID,
			"traveler_schedule.id": data.TravelID,
		}

		_, errData := dataOrderService.repoTravel.GetDetail(poolData.DBSql, where)
		if errData != nil {
			log.Printf("failed Get Data travel on func checkData order : %s \n", errData.Error())
			err = errorhandler.ErrGetData(errData)
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()

		resultAddress, errData := dataOrderService.authClient.GetAddrByID(ctx, int32(userID), int32(data.AddressID))
		if errData != nil {
			log.Printf("failed Get Data addr on func checkData order : %s \n", errData.Error())
			err = errorhandler.ErrGetData(errData)
			errChan <- err
			return
		}

		addrResult = domain.AddressResponse{
			ReceiverName:  resultAddress.ReceiverName,
			ReceiverPhone: resultAddress.ReceiverPhone,
			Province:      resultAddress.Province,
			Street:        resultAddress.Street,
			City:          resultAddress.City,
			District:      resultAddress.District,
			SUbDistrict:   resultAddress.SubDistrict,
			PostalCode:    resultAddress.PostalCode,
			Tag:           resultAddress.Tag,
		}

	}()

	go func() {
		defer wg.Done()

		for _, v := range data.Product {
			where := map[string]any{
				"id":      v.ID,
				"user_id": userID,
			}
			resultProduct, errData := dataOrderService.repoProduct.Get(ctx, poolData.DBSql, where)
			if errData != nil {
				log.Printf("failed Get Data product on func checkData order : %s \n", errData.Error())
				err = errorhandler.ErrGetData(errData)
				errChan <- err
				return
			}

			whereExpr := gorm.Expr("product_id = ? AND traveler_schedule_id = ? ", v.ID, data.TravelID)

			resultProductTravel, errData := dataOrderService.repoProductTravel.GetExpr(ctx, poolData.DBSql, whereExpr)
			if errData != nil {
				log.Printf("failed Get Data product travel on func checkData order : %s \n", errData.Error())
				err = errorhandler.ErrGetData(errData)
				errChan <- err
				return
			} else if len(resultProductTravel) == 0 {
				log.Printf("failed Get Data product travel on func checkData order : product not found \n")
				err = errorhandler.ErrGetData(nil)
				errChan <- err
				return
			}

			dataOrderDetail := domain.OrderDetail{
				ProductID:       resultProduct.ID,
				ProductName:     resultProduct.Name,
				ProductImage:    resultProduct.Image,
				ProductPrice:    float32(resultProduct.Price),
				ProductQuantity: resultProduct.Quantity,
				Quantity:        v.Quantity,
				Price:           float32(v.Quantity) * float32(resultProduct.Price),
			}

			price += dataOrderDetail.Price
			productResult = append(productResult, dataOrderDetail)
		}

	}()

	wg.Wait()
	close(errChan)
	for v := range errChan {
		if v.Code != 0 {
			err = v
			return
		}
	}

	return
}

func createOrder(ctx context.Context, dataOrderService orderService, poolData *domain.Config, dataOrder domain.OrderData, dataAddrOrder domain.AddressOrder, dataOrderDetail []domain.OrderDetail) (err domain.ErrorData) {
	connTrx := poolData.DBSql.Begin()
	defer func() {
		if err.Code != 0 {
			connTrx.Rollback()
		} else {
			connTrx.Commit()
		}
	}()

	idAddr, errData := dataOrderService.repoAddrOrder.Create(ctx, connTrx, dataAddrOrder)
	if errData != nil {
		log.Printf("failed Create data address order on func createOrder order : %s \n", errData.Error())
		err = errorhandler.ErrInsertData(errData)
		return
	}

	dataOrder.AddressID = idAddr
	idOrder, errData := dataOrderService.repo.Create(ctx, connTrx, dataOrder)
	if errData != nil {
		log.Printf("failed Create data order on func createOrder order : %s \n", errData.Error())
		err = errorhandler.ErrInsertData(errData)
		return
	}

	for i, _ := range dataOrderDetail {
		dataOrderDetail[i].OrderID = idOrder
	}

	errData = dataOrderService.repoOrderDetail.CreateBulk(ctx, connTrx, dataOrderDetail)
	if errData != nil {
		log.Printf("failed Create data order detail on func createOrder order : %s \n", errData.Error())
		err = errorhandler.ErrInsertData(errData)
		return
	}

	return
}
