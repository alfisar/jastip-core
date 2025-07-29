package repository

import (
	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type ProductsRepositoryContract interface {
	Create(conn *gorm.DB, data domain.ProductData) (err error)
	GetList(conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.ProductResp, total int64, err error)
	Get(conn *gorm.DB, where map[string]any) (result domain.ProductResp, err error)
	Update(conn *gorm.DB, update map[string]any, where map[string]any) (err error)
	Delete(conn *gorm.DB, where map[string]any) (err error)
}
