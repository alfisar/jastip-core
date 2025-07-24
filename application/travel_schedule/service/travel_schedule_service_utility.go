package service

import (
	"fmt"
	repoCountries "jastip-core/application/countries/repository"
	"jastip-core/application/travel_schedule/repository"
	"log"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/handler"
	"github.com/alfisar/jastip-import/helpers/helper"

	"gorm.io/gorm"
)

func validatedTravelTime(poolData *domain.Config, userID int, startDate string, endDate string, repo repository.TravelSchRepositoryContract) (err domain.ErrorData) {
	result, errData := repo.GetByTimeBetween(poolData.DBSql, userID, startDate, endDate)
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

	startDates, errData := helper.GenerateDateTime(startDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}

	endDates, errData := helper.GenerateDateTime(endDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}

	if endDates.Before(startDates) {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : End Dates tidak boleh lebih kecil dari start date ")
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}
	return
}

func validatedUpdateTravelTime(poolData *domain.Config, id int, userID int, startDate string, endDate string, repo repository.TravelSchRepositoryContract) (err domain.ErrorData) {
	result, errData := repo.GetByTimeBetween(poolData.DBSql, userID, startDate, endDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else if errData.Error() != "get traveler schedule error : "+gorm.ErrRecordNotFound.Error() {
			err = errorhandler.ErrGetData(errData)
		}

	} else if result.ID != 0 {
		if result.ID != id {
			message := fmt.Sprintf("Error validate data on func validatedTravelTime : Schedule sudah ada!!")
			log.Println(message)
			err = errorhandler.ErrValidation(fmt.Errorf(message))
		}

	}

	startDates, errData := helper.GenerateDateTime(startDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}

	endDates, errData := helper.GenerateDateTime(endDate)
	if errData != nil {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : %s", errData.Error())
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}

	if endDates.Before(startDates) {
		message := fmt.Sprintf("Error validate data on func validatedTravelTime : End Dates tidak boleh lebih kecil dari start date ")
		log.Println(message)
		err = errorhandler.ErrValidation(fmt.Errorf(message))
	}
	return
}

func validateCountries(poolData *domain.Config, countriesID int, repo repoCountries.CountriesRepositoryContract) (err domain.ErrorData) {
	where := map[string]any{
		"id": countriesID,
	}
	_, errData := repo.Get(poolData.DBSql, where)
	if errData != nil {
		message := fmt.Sprintf("Error validate data countries on func validateCountries : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else if errData.Error() != "get one countries data error : "+gorm.ErrRecordNotFound.Error() {
			err = errorhandler.ErrGetData(errData)
		}

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
			errorhandler.ErrInsertData(fmt.Errorf(message))
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

func getDetail(poolData *domain.Config, repo repository.TravelSchRepositoryContract, id int, userID int) (result domain.TravelSchResponse, err domain.ErrorData) {
	var errData error

	where := map[string]any{
		"id":      id,
		"user_id": userID,
	}

	result, errData = repo.GetDetail(poolData.DBSql, where)
	if errData != nil {
		message := fmt.Sprintf("Error get detail data on func getDetail : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			err = errorhandler.ErrGetData(fmt.Errorf(message))
		}
	}

	return
}

func updateSchedule(poolData *domain.Config, repo repository.TravelSchRepositoryContract, id int, updates map[string]any) (err domain.ErrorData) {
	where := map[string]any{
		"id": id,
	}
	errData := repo.Update(poolData.DBSql, where, updates)
	if errData != nil {
		message := fmt.Sprintf("Error update data on func updateSchedule : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrUpdateData(fmt.Errorf(message))
		}
	}

	return
}

func deleteSchedule(poolData *domain.Config, repo repository.TravelSchRepositoryContract, id int, userID int) (err domain.ErrorData) {

	where := map[string]any{
		"id":      id,
		"user_id": userID,
	}

	errData := repo.Delete(poolData.DBSql, where)
	if errData != nil {
		message := fmt.Sprintf("Error delete data on func deleteSchedule : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrUpdateData(fmt.Errorf(message))
		}
	}

	return
}

func mappingDataUpdate(data domain.TravelSchResponse, updates map[string]any) (result domain.TravelSchResponse) {
	if v, exist := updates["location"]; exist {
		data.Location = v.(string)
	}

	if v, exist := updates["period_start"]; exist {
		data.PeriodStart = v.(string)
	}

	if v, exist := updates["period_end"]; exist {
		data.PeriodEnd = v.(string)
	}

	result = data
	return
}
