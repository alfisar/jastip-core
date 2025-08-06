package service

import (
	"context"
	"fmt"
	repositoryProduct "jastip-core/application/products/repository"
	"jastip-core/application/products_travel/repository"
	repositoryTravel "jastip-core/application/travel_schedule/repository"
	"log"
	"strings"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type key struct {
	ProductID int
	TravelID  int
}

func checkData(products []int, travel []int, userID int, repoProducts repositoryProduct.ProductsRepositoryContract, repoTravel repositoryTravel.TravelSchRepositoryContract, conn *gorm.DB) (result domain.ProductsTravelRequest, err domain.ErrorData) {
	var (
		wg                          sync.WaitGroup
		travelResult, productResult []int
	)
	errChan := make(chan domain.ErrorData, 2)
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// check data product
	go func() {
		defer wg.Done()

		productsExist, errData := repoProducts.Gets(ctx, conn, map[string]any{
			"user_id": userID,
		})

		if errData != nil {
			log.Printf("failed Get Data product  on func checkData : %s \n", errData.Error())

			errChan <- errorhandler.ErrGetData(errData)
			cancel()
			return
		}

		mapProduct := make(map[int]bool)
		if len(productsExist) > 0 {
			for _, v := range productsExist {
				mapProduct[v.ID] = true
			}

			for _, v := range products {
				if mapProduct[v] {
					productResult = append(productResult, v)
				}
			}
		}

	}()

	// check data travel
	go func() {
		defer wg.Done()

		travelExist, errData := repoTravel.Gets(ctx, conn, map[string]any{
			"user_id": userID,
		})

		if errData != nil {
			log.Printf("failed Get Data travel on func checkData : %s \n", errData.Error())

			errChan <- errorhandler.ErrGetData(errData)
			cancel()
			return
		}

		mapTravel := make(map[int]bool)
		if len(travelExist) > 0 {
			for _, v := range travelExist {
				mapTravel[v.ID] = true
			}

			for _, v := range travel {
				if mapTravel[v] {
					travelResult = append(travelResult, v)
				}
			}
		}

	}()

	wg.Wait()
	close(errChan)
	for v := range errChan {
		if v.Code != 0 {
			err = v
			return
		}
	}

	result.ProductID = productResult
	result.TravelID = travelResult
	return
}

func filteredData(ctx context.Context, conn *gorm.DB, repo repository.ProductTravelRepositoryContract, data []domain.ProductsTravel) (filterData []domain.ProductsTravel, err domain.ErrorData) {
	dataUnfilter := []string{}
	value := []any{}
	for _, v := range data {
		dataUnfilter = append(dataUnfilter, "(?,?)")
		value = append(value, v.ProductID, v.TravelID)
	}

	stringExpr := fmt.Sprintf("(product_id, traveler_schedule_id) IN (%s)", strings.Join(dataUnfilter, ","))
	where := gorm.Expr(stringExpr, value...)
	result, errData := repo.GetExpr(ctx, conn, where)
	if errData != nil {
		log.Printf("failed Get Data product Travel on func filteredData : %s \n", errData.Error())

		err = errorhandler.ErrGetData(errData)
		return
	}

	exists := make(map[key]bool)
	for _, v := range result {
		exists[key{v.ProductID, v.TravelID}] = true
	}

	for _, v := range data {
		if !exists[key{v.ProductID, v.TravelID}] {
			filterData = append(filterData, v)
		}
	}

	return
}
