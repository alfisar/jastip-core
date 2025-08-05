package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductTravelRepositoryContract interface {
	CreateBulk(ctx context.Context, conn *gorm.DB, data []domain.ProductsTravel) (err error)
	DeleteBulk(ctx context.Context, conn *gorm.DB, where map[string]any) (err error)
	GetExpr(ctx context.Context, conn *gorm.DB, where clause.Expr) (result []domain.ProductsTravel, err error)
}
