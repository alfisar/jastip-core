package service

import "github.com/alfisar/jastip-import/domain"

type CountriesServiceContract interface {
	GetList(poolData *domain.Config, param domain.Params) (result []domain.Countries, currentPage int, count int64, limits int, totalPage int, err domain.ErrorData)
}
