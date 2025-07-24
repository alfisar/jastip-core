package repository

import (
	"fmt"
	"strings"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type countriesRepository struct{}

func NewCountriesRepository() *countriesRepository {
	return &countriesRepository{}
}

func (r *countriesRepository) Gets(conn *gorm.DB, page int, limit int, where map[string]any, param domain.Params) (result []domain.Countries, count int64, err error) {
	var wg sync.WaitGroup

	errChan := make(chan error, 2)
	wg.Add(2)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	go func() {
		defer wg.Done()
		errData := conn.Debug().Table("countries").Where(where).Where("LOWER(name) like ?", "%"+strings.ToLower(param.Search)+"%").Offset(page).Limit(limit).Find(&result).Order("name ASC").Error
		if errData != nil {
			errData = fmt.Errorf("get countries data error : %w", errData)
			errChan <- errData
			return
		}
	}()

	go func() {
		defer wg.Done()
		errData := conn.Debug().Table("countries").Where(where).Where("LOWER(name) like ?", "%"+strings.ToLower(param.Search)+"%").Count(&count).Error
		if errData != nil {
			errData = fmt.Errorf("get countries count error : %w", errData)
			errChan <- errData
			return
		}
	}()

	wg.Wait()
	close(errChan)
	for v := range errChan {
		err = v
		return
	}

	return
}

func (r *countriesRepository) Get(conn *gorm.DB, where map[string]any) (result domain.Countries, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("countries").Where(where).First(&result).Error
	if err != nil {
		err = fmt.Errorf("get one countries data error : %w", err)

		return
	}
	return
}
