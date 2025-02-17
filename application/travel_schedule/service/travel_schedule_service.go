package service

import (
	"context"
	"jastip-core/application/travel_schedule/repository"

	"github.com/alfisar/jastip-import/domain"
)

type travelSchService struct {
	repo repository.TravelSchRepositoryContract
}

func NewTravelSchService(repo repository.TravelSchRepositoryContract) *travelSchService {
	return &travelSchService{
		repo: repo,
	}
}

func (s *travelSchService) AddSchedule(ctx context.Context, poolData *domain.Config, data domain.TravelSchRequest) (id int, err domain.ErrorData) {
	err = validatedTravelTime(poolData, data.UserID, data.Location, data.PeriodStart, data.PeriodEnd, s.repo)
	if err.Code != 0 {
		return
	}
	data.Status = 1
	id, err = addSchedule(data, poolData, s.repo)
	return
}
