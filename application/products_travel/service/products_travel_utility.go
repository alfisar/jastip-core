package service

import (
	"context"
	"fmt"
	repositoryProduct "jastip-core/application/products/repository"
	"jastip-core/application/products_travel/repository"
	repositoryTravel "jastip-core/application/travel_schedule/repository"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

func makesData(products []int, idTravel int) (result []domain.ProductsTravel) {
	for _, v := range products {
		result = append(result, domain.ProductsTravel{
			ProductID: v,
			TravelID:  idTravel,
		})
	}

	return
}

func checkData(products []int, travel []int, userID int, repoProducts repositoryProduct.ProductsRepositoryContract, repoTravel repositoryTravel.TravelSchRepositoryContract, conn *gorm.DB) (result domain.ProductsTravelRequest, err domain.ErrorData) {
	var (
		wg                          sync.WaitGroup
		travelResult, productResult []int
	)

	wg.Add(2)
	// check data product
	go func() {
		defer wg.Done()
		productsExist, errData := repoProducts.Gets(conn, map[string]any{
			"user_id": userID,
		})

		if errData != nil {
			message := fmt.Sprintf("failed Get Data product  on func checkData : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrGetData(errData)
			return
		}

		mapProduct := make(map[string]bool)
		if len(productsExist) > 0 {
			for _, v := range productsExist {
				mapProduct[strconv.Itoa(v.ID)] = true
			}

			for _, v := range products {
				if mapProduct[strconv.Itoa(v)] {
					productResult = append(productResult, v)
				}
			}
		}
		return
	}()

	// check data travel
	go func() {
		defer wg.Done()
		travelExist, errData := repoTravel.Gets(conn, map[string]any{
			"user_id": userID,
		})

		if errData != nil {
			message := fmt.Sprintf("failed Get Data travel on func checkData : %s", errData.Error())
			log.Println(message)

			err = errorhandler.ErrGetData(errData)
			return
		}

		mapTravel := make(map[string]bool)
		if len(travelExist) > 0 {
			for _, v := range travelExist {
				mapTravel[strconv.Itoa(v.ID)] = true
			}

			for _, v := range travel {
				if mapTravel[strconv.Itoa(v)] {
					travelResult = append(travelResult, v)
				}
			}
		}
		return
	}()

	wg.Wait()
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
		message := fmt.Sprintf("failed Get Data product Travel on func filteredData : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrGetData(errData)
		return
	}

	exists := make(map[string]bool)
	for _, v := range result {
		exists[strconv.Itoa(v.ProductID)+"-"+strconv.Itoa(v.TravelID)] = true
	}

	for _, v := range data {
		if !exists[strconv.Itoa(v.ProductID)+"-"+strconv.Itoa(v.TravelID)] {
			filterData = append(filterData, v)
		}
	}

	return
}

func createData(ctx context.Context, conn *gorm.DB, repo repository.ProductTravelRepositoryContract, data []domain.ProductsTravel) (err domain.ErrorData) {
	errData := repo.CreateBulk(ctx, conn, data)
	if errData != nil {
		message := fmt.Sprintf("failed createData product Travel on func createData : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrInsertData(errData)
		return
	}

	return
}

func deleteData(ctx context.Context, conn *gorm.DB, repo repository.ProductTravelRepositoryContract, travelID int, productsID []int) (err domain.ErrorData) {
	where := map[string]any{
		"traveler_schedule_id": travelID,
		"IN":                   productsID,
	}

	errData := repo.DeleteBulk(ctx, conn, where)
	if errData != nil {
		message := fmt.Sprintf("failed delete data product Travel on func deleteData : %s", errData.Error())
		log.Println(message)

		err = errorhandler.ErrDeleteData(errData)
		return
	}
	return
}
