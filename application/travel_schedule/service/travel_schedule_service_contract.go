package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type TraveklSchServiceContract interface {
	AddSchedule(ctx context.Context, poolData *domain.Config, data domain.TravelSchRequest) (id int, err domain.ErrorData)
}
