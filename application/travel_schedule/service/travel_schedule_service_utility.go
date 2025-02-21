package service

import (
	"fmt"
	"jastip-core/application/travel_schedule/repository"
	"log"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/handler"
	"gorm.io/gorm"
)

func validatedTravelTime(poolData *domain.Config, userID int, locations string, startDate string, endDate string, repo repository.TravelSchRepositoryContract) (err domain.ErrorData) {
	result, errData := repo.GetByTimeBetween(poolData.DBSql, userID, locations, startDate, endDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else if errData.Error() != "get traveler schedule error : "+gorm.ErrRecordNotFound.Error() {
			err = errorhandler.ErrGetData(errData)
		}

	} else if result.ID != 0 {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : Schedule sudah ada!!")
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))

	}
	return
}

func addSchedule(data domain.TravelSchRequest, poolData *domain.Config, repo repository.TravelSchRepositoryContract) (id int, err domain.ErrorData) {
	var errData error
	id, errData = repo.Create(poolData.DBSql, data)
	if errData != nil {
		message := fmt.Sprintf("Error insert data on func addSchedule : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrInsertData(nil)
		}
	}
	return
}

func getList(poolData *domain.Config, param domain.Params, repo repository.TravelSchRepositoryContract) (result []domain.TravelSchResponse, currentPage int, total int64, err domain.ErrorData) {
	var errData error
	pages, offset, limit := handler.CalculateOffsetAndLimit(param.Page, param.Limit)

	where := map[string]any{
		"status": param.Status,
	}

	result, total, errData = repo.GetList(poolData.DBSql, where, param.Search, offset, limit)
	if errData != nil {
		message := fmt.Sprintf("Error Get List data on func getList : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			err = errorhandler.ErrGetData(errData)
		}
		return
	}
	currentPage = pages
	return
}
