package service

import (
	"context"
	repoCountries "jastip-core/application/countries/repository"
	"jastip-core/application/travel_schedule/repository"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/handler"
)

type travelSchService struct {
	repo          repository.TravelSchRepositoryContract
	repoCountries repoCountries.CountriesRepositoryContract
}

func NewTravelSchService(repo repository.TravelSchRepositoryContract, repoCountries repoCountries.CountriesRepositoryContract) *travelSchService {
	return &travelSchService{
		repo:          repo,
		repoCountries: repoCountries,
	}
}

func (s *travelSchService) AddSchedule(ctx context.Context, poolData *domain.Config, data domain.TravelSchRequest) (id int, err domain.ErrorData) {
	var wg sync.WaitGroup

	errChan := make(chan domain.ErrorData, 2)
	wg.Add(2)

	go func() {
		defer wg.Done()
		err = validatedTravelTime(poolData, data.UserID, data.PeriodStart, data.PeriodEnd, s.repo)
		if err.Code != 0 {
			errChan <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = validateCountries(poolData, data.Location, s.repoCountries)
		if err.Code != 0 {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	close(errChan)
	for v := range errChan {
		err = v
		return
	}

	data.Status = 1
	id, err = addSchedule(data, poolData, s.repo)
	return
}

func (s *travelSchService) GetList(ctx context.Context, poolData *domain.Config, param domain.Params) (totalPage int, currentPage int, total int64, result []domain.TravelSchResponse, err domain.ErrorData) {
	result, currentPage, total, err = getList(poolData, param, s.repo)
	if err.Code != 0 {
		return
	}

	totalPage = int(handler.CalculateTotalPages(total, int64(param.Limit)))
	return
}

func (s *travelSchService) GetDetails(ctx context.Context, poolData *domain.Config, id int, userID int) (result domain.TravelSchResponse, err domain.ErrorData) {
	result, err = getDetail(poolData, s.repo, id, userID)
	return
}

func (s *travelSchService) Update(ctx context.Context, poolData *domain.Config, id int, userID int, update map[string]any) (err domain.ErrorData) {
	details, errData := getDetail(poolData, s.repo, id, userID)
	if errData.Code != 0 {
		err = errData
		return
	}

	dataMapping := mappingDataUpdate(details, update)
	err = validatedUpdateTravelTime(poolData, id, userID, dataMapping.PeriodStart, dataMapping.PeriodEnd, s.repo)
	if err.Code != 0 {
		return
	}

	err = updateSchedule(poolData, s.repo, id, update)
	return
}

func (s *travelSchService) Delete(ctx context.Context, poolData *domain.Config, id int, userID int) (err domain.ErrorData) {
	err = deleteSchedule(poolData, s.repo, id, userID)
	return
}
