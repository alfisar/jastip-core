package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type OrderRepositoryContract interface {
	Create(ctx context.Context, conn *gorm.DB, data domain.OrderData) (id int, err error)
	Get(ctx context.Context, conn *gorm.DB, where map[string]any) (result domain.OrderOneResponse, err error)
	Getlist(ctx context.Context, conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.OrderListResponse, total int64, err error)
}
