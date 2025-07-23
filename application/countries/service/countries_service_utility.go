package service

import (
	"fmt"
	"jastip-core/application/countries/repository"
	"log"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"github.com/alfisar/jastip-import/helpers/handler"
	"gorm.io/gorm"
)

func getList(conn *gorm.DB, param domain.Params, repo repository.CountriesRepositoryContract) (result []domain.Countries, currentPage int, limits int, count int64, err domain.ErrorData) {
	var (
		offset  int
		errData error
	)

	where := map[string]any{}

	currentPage, offset, limits = handler.CalculateOffsetAndLimit(param.Page, param.Limit)
	result, count, errData = repo.Gets(conn, offset, limits, where, param)
	if errData != nil {
		message := fmt.Sprintf("Error Get List data on func getList coutries : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			err = errorhandler.ErrGetData(errData)
		}
		return
	}

	return

}
