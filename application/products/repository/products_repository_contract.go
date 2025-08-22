package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type ProductsRepositoryContract interface {
	Create(ctx context.Context, conn *gorm.DB, data domain.ProductData) (err error)
	GetList(ctx context.Context, conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.ProductResp, total int64, err error)
	Get(ctx context.Context, conn *gorm.DB, where map[string]any) (result domain.ProductResp, err error)
	Gets(ctx context.Context, conn *gorm.DB, where map[string]any) (result []domain.ProductResp, err error)
	GetListProductTravel(ctx context.Context, conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.ProductResp, total int64, err error)
	Update(ctx context.Context, conn *gorm.DB, update map[string]any, where map[string]any) (err error)
	Delete(ctx context.Context, conn *gorm.DB, where map[string]any) (err error)
}
