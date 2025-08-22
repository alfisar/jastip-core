package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type OrderServiceContract interface {
	Create(ctx context.Context, poolData *domain.Config, data domain.OrderRequest, userID int) (err domain.ErrorData)
	GetList(ctx context.Context, poolData *domain.Config, params domain.Params, userID int) (totalPage int, currentPage int, total int64, result []domain.OrderListResponse, err domain.ErrorData)
	GetDetails(ctx context.Context, poolData *domain.Config, orderID int, userID int) (result domain.OrderOneResponse, err domain.ErrorData)
}
