package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type TraveklSchServiceContract interface {
	AddSchedule(ctx context.Context, poolData *domain.Config, data domain.TravelSchRequest) (id int, err domain.ErrorData)
	GetList(ctx context.Context, poolData *domain.Config, param domain.Params) (totalPage int, currentPage int, total int64, result []domain.TravelSchResponse, err domain.ErrorData)
	GetDetails(ctx context.Context, poolData *domain.Config, id int, userID int) (result domain.TravelSchResponse, err domain.ErrorData)
	Update(ctx context.Context, poolData *domain.Config, id int, userID int, update map[string]any) (err domain.ErrorData)
	Delete(ctx context.Context, poolData *domain.Config, id int, userID int) (err domain.ErrorData)
}
