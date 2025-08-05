package service

import (
	"context"
	repositoryProduct "jastip-core/application/products/repository"
	"jastip-core/application/products_travel/repository"
	repositoryTravel "jastip-core/application/travel_schedule/repository"

	"github.com/alfisar/jastip-import/domain"
)

type productsTravelService struct {
	repo        repository.ProductTravelRepositoryContract
	repoProduct repositoryProduct.ProductsRepositoryContract
	repoTravel  repositoryTravel.TravelSchRepositoryContract
}

func NewProductsTravelService(repo repository.ProductTravelRepositoryContract, repoProduct repositoryProduct.ProductsRepositoryContract, repoTravel repositoryTravel.TravelSchRepositoryContract) *productsTravelService {
	return &productsTravelService{
		repo:        repo,
		repoProduct: repoProduct,
		repoTravel:  repoTravel,
	}
}

func (s *productsTravelService) Create(ctx context.Context, poolData *domain.Config, userID int, data domain.ProductsTravelRequest) (err domain.ErrorData) {

	data, err = checkData(data.ProductID, data.TravelID, userID, s.repoProduct, s.repoTravel, poolData.DBSql)

	connTrx := poolData.DBSql.Begin()

	for _, v := range data.TravelID {

		dataReq := makesData(data.ProductID, v)
		dataFiltered, errData := filteredData(ctx, connTrx, s.repo, dataReq)
		if errData.Code != 0 {
			err = errData
			connTrx.Rollback()
			return
		}
		if len(dataFiltered) > 0 {
			err = createData(ctx, connTrx, s.repo, dataFiltered)
			if err.Code != 0 {
				connTrx.Rollback()
				return
			}
		}
	}
	connTrx.Commit()
	return
}

func (s *productsTravelService) Delete(ctx context.Context, poolData *domain.Config, data domain.ProductsTravelRequest) (err domain.ErrorData) {

	for _, v := range data.TravelID {
		err = deleteData(ctx, poolData.DBSql, s.repo, v, data.ProductID)
		if err.Code != 0 {
			return
		}
	}
	return
}
