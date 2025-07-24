package repository

import (
	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type CountriesRepositoryContract interface {
	Gets(conn *gorm.DB, page int, limit int, where map[string]any, param domain.Params) (result []domain.Countries, count int64, err error)
	Get(conn *gorm.DB, where map[string]any) (result domain.Countries, err error)
}
