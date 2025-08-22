package service

import (
	"context"
	repositoryAddrOrder "jastip-core/application/address_order/repository"
	"jastip-core/application/order/repository"
	repositoryOrderDetail "jastip-core/application/order_detail/repository"
	repositoryProduct "jastip-core/application/products/repository"
	repositoryProductTravel "jastip-core/application/products_travel/repository"
	repositoryTravel "jastip-core/application/travel_schedule/repository"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"
	clientAuth "github.com/alfisar/jastip-import/proto/auth/client"
)

type orderService struct {
	repo              repository.OrderRepositoryContract
	repoOrderDetail   repositoryOrderDetail.OrderDetailRepositoryContract
	repoAddrOrder     repositoryAddrOrder.AddressOrderRepositoryContract
	repoProduct       repositoryProduct.ProductsRepositoryContract
	repoProductTravel repositoryProductTravel.ProductTravelRepositoryContract
	repoTravel        repositoryTravel.TravelSchRepositoryContract
	authClient        clientAuth.AuthClientContract
}

func NewOrderService(repo repository.OrderRepositoryContract, repoOrderDetail repositoryOrderDetail.OrderDetailRepositoryContract, repoAddrOrder repositoryAddrOrder.AddressOrderRepositoryContract, repoProduct repositoryProduct.ProductsRepositoryContract, repoProductTravel repositoryProductTravel.ProductTravelRepositoryContract, repoTravel repositoryTravel.TravelSchRepositoryContract, authClient clientAuth.AuthClientContract) *orderService {
	return &orderService{
		repo:              repo,
		repoOrderDetail:   repoOrderDetail,
		repoAddrOrder:     repoAddrOrder,
		authClient:        authClient,
		repoProduct:       repoProduct,
		repoProductTravel: repoProductTravel,
		repoTravel:        repoTravel,
	}
}

func (s *orderService) Create(ctx context.Context, poolData *domain.Config, data domain.OrderRequest, userID int) (err domain.ErrorData) {
	price, addr, products, errs := checkData(ctx, *s, poolData, userID, data)
	if errs.Code != 0 {
		err = errs
		return
	}

	if price != data.Price {
		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidInput, "Harga total tidak telah berubah, silahkan coba kembali", "")
		return
	}

	dataOrder := domain.OrderData{
		BuyerID:       data.AddressID,
		Invoice:       helper.GenerateInvoiceNumber(),
		TravelID:      data.TravelID,
		Price:         data.Price,
		Status:        1,
		PaymentStatus: 1,
		PaymentMethod: data.PaymentMethod,
	}

	dataAddr := domain.AddressOrder{
		UserID:        data.BuyerID,
		ReceiverName:  addr.ReceiverName,
		ReceiverPhone: addr.ReceiverPhone,
		Province:      addr.Province,
		Street:        addr.Street,
		City:          addr.City,
		District:      addr.District,
		SUbDistrict:   addr.SUbDistrict,
		PostalCode:    addr.PostalCode,
		Tag:           addr.Tag,
	}

	err = createOrder(ctx, *s, poolData, dataOrder, dataAddr, products)
	return

}

func (s *orderService) GetList(ctx context.Context, poolData *domain.Config, params domain.Params, userID int) (totalPage int, currentPage int, total int64, result []domain.OrderListResponse, err domain.ErrorData) {
	var errData error
	where := map[string]any{
		"buyer_id": userID,
	}

	if params.Status != 0 {
		where["status"] = params.Status
	}

	pages, offset, limit := handler.CalculateOffsetAndLimit(params.Page, params.Limit)
	result, total, errData = s.repo.Getlist(ctx, poolData.DBSql, params, where, offset, limit)
	if errData != nil {
		errorhandler.ErrGetData(errData)
		return
	}

	totalPage = int(handler.CalculateTotalPages(total, int64(params.Limit)))
	currentPage = pages
	return
}

func (s *orderService) GetDetails(ctx context.Context, poolData *domain.Config, orderID int, userID int) (result domain.OrderOneResponse, err domain.ErrorData) {
	var errData error
	where := map[string]any{
		"id":       orderID,
		"buyer_id": userID,
	}

	result, errData = s.repo.Get(ctx, poolData.DBSql, where)
	if errData != nil {
		errorhandler.ErrGetData(errData)
		return
	}
	return
}
