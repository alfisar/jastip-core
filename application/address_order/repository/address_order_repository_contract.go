package repository

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type AddressOrderRepositoryContract interface {
	Create(ctx context.Context, conn *gorm.DB, data domain.AddressOrder) (id int, err error)
}
