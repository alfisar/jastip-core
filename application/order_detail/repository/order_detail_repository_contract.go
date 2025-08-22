package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type OrderDetailRepositoryContract interface {
	CreateBulk(ctx context.Context, conn *gorm.DB, data []domain.OrderDetail) (err error)
}
