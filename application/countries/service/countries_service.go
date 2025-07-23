package service

import (
	"jastip-core/application/countries/repository"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/handler"
)

type contrieService struct {
	repo repository.CountriesRepositoryContract
}

func NewCountriesService(repo repository.CountriesRepositoryContract) *contrieService {
	return &contrieService{
		repo: repo,
	}
}

func (s *contrieService) GetList(poolData *domain.Config, param domain.Params) (result []domain.Countries, currentPage int, count int64, limits int, totalPage int, err domain.ErrorData) {
	result, currentPage, limits, count, err = getList(poolData.DBSql, param, s.repo)
	if err.Code != 0 {
		return
	}

	totalPage = int(handler.CalculateTotalPages(count, int64(param.Limit)))
	return
}
