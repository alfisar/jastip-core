package service

import (
	"context"
	repositoryProduct "jastip-core/application/products/repository"
	"jastip-core/application/products_travel/repository"
	repositoryTravel "jastip-core/application/travel_schedule/repository"
	"log"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
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

	if err.Code != 0 {
		return
	}

	connTrx := poolData.DBSql.Begin()

	for _, v := range data.TravelID {

		dataReq := []domain.ProductsTravel{}
		for _, value := range data.ProductID {
			dataReq = append(dataReq, domain.ProductsTravel{
				ProductID: value,
				TravelID:  v,
			})
		}

		dataFiltered, errData := filteredData(ctx, connTrx, s.repo, dataReq)
		if errData.Code != 0 {
			err = errData
			connTrx.Rollback()
			return
		}
		if len(dataFiltered) > 0 {
			errData := s.repo.CreateBulk(ctx, connTrx, dataFiltered)
			if errData != nil {
				log.Printf("failed createData product Travel on func createData : %s \n", errData.Error())

				err = errorhandler.ErrInsertData(errData)
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
		where := map[string]any{
			"traveler_schedule_id": v,
			"IN":                   data.ProductID,
		}

		errData := s.repo.DeleteBulk(ctx, poolData.DBSql, where)
		if errData != nil {
			log.Printf("failed delete data product Travel on func Delete : %s \n", errData.Error())

			err = errorhandler.ErrDeleteData(errData)
			return
		}
	}
	return
}
