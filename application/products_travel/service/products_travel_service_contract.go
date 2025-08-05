package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type ProductsTravelServiceContract interface {
	Create(ctx context.Context, poolData *domain.Config, userID int, data domain.ProductsTravelRequest) (err domain.ErrorData)
	Delete(ctx context.Context, poolData *domain.Config, data domain.ProductsTravelRequest) (err domain.ErrorData)
}
