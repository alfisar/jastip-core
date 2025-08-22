package service

import (
	"context"
	"mime/multipart"

	"github.com/alfisar/jastip-import/domain"
	"github.com/valyala/fasthttp"
)

type ProductsServiceContract interface {
	Create(ctx context.Context, poolData *domain.Config, data domain.ProductData, r *fasthttp.Request) (err domain.ErrorData)
	GetList(ctx context.Context, poolData *domain.Config, userID int, params domain.Params) (totalPage int, currentPage int, total int64, limit int, result []domain.ProductResp, err domain.ErrorData)
	GetListProductTravel(ctx context.Context, poolData *domain.Config, userID int, travelID int, params domain.Params) (totalPage int, currentPage int, total int64, limit int, result []domain.ProductResp, err domain.ErrorData)
	Update(ctx context.Context, poolData *domain.Config, id int, userID int, update map[string]any, file *multipart.Form) (err domain.ErrorData)
	Delete(ctx context.Context, poolData *domain.Config, id int, userId int) (err domain.ErrorData)
}
