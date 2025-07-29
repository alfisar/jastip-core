package repository

import (
	"fmt"
	"strings"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type productsRepository struct{}

func NewProductsRepository() *productsRepository {
	return &productsRepository{}
}

func (r *productsRepository) Create(conn *gorm.DB, data domain.ProductData) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("products").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("create product error : %w", err)
		return
	}

	return

}

func (r *productsRepository) GetList(conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.ProductResp, total int64, err error) {
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
		errData := conn.Debug().Table("products").Where(where).Where("LOWER(name) like ?", "%"+strings.ToLower(param.Search)+"%").Offset(offset).Limit(limit).Find(&result).Order("name ASC").Error
		if errData != nil {
			errData = fmt.Errorf("get products data error : %w", errData)
			errChan <- errData
			return
		}
	}()

	go func() {
		defer wg.Done()
		errData := conn.Debug().Table("products").Where(where).Where("LOWER(name) like ?", "%"+strings.ToLower(param.Search)+"%").Count(&total).Error
		if errData != nil {
			errData = fmt.Errorf("get products count error : %w", errData)
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

func (r *productsRepository) Update(conn *gorm.DB, update map[string]any, where map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("products").Where(where).Updates(&update).Error
	if err != nil {
		err = fmt.Errorf("create product error : %w", err)
		return
	}

	return

}

func (r *productsRepository) Delete(conn *gorm.DB, where map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("products").Where(where).Delete(&domain.ProductData{}).Error
	if err != nil {
		err = fmt.Errorf("create product error : %w", err)
		return
	}

	return

}
