package repository

import (
	"github.com/alfisar/jastip-import/domain"
	"gorm.io/gorm"
)

type TravelSchRepositoryContract interface {
	Create(conn *gorm.DB, data domain.TravelSchRequest) (id int, err error)
	GetList(conn *gorm.DB, where map[string]any, search string, offet int, limit int) (result []domain.TravelSchResponse, count int64, err error)
	GetByTimeBetween(conn *gorm.DB, id int, locations string, startDate string, endDate string) (result domain.TravelSchResponse, err error)
	Update(conn *gorm.DB, where map[string]any, updates map[string]any) (err error)
	Delete(conn *gorm.DB, ID int) (err error)
}
