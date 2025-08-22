package repository

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type orderRepository struct{}

func NewOrderRepository() *orderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Create(ctx context.Context, conn *gorm.DB, data domain.OrderData) (id int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("orders").Save(&data).Error
	id = data.ID
	return
}

func (r *orderRepository) Get(ctx context.Context, conn *gorm.DB, where map[string]any) (result domain.OrderOneResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("orders").Where(where).Preload("Travel").Preload("Address").Preload("Product").First(&result).Error
	if err != nil {
		err = fmt.Errorf("get order data error : %w", err)
	}
	return
}

func (r *orderRepository) Getlist(ctx context.Context, conn *gorm.DB, param domain.Params, where map[string]any, offset int, limit int) (result []domain.OrderListResponse, total int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

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
		errData := conn.WithContext(ctx).Debug().Table("orders").Where(where).Where("LOWER(invoice) like ?", "%"+strings.ToLower(param.Search)+"%").Offset(offset).Limit(limit).Preload("Travel").Preload("Address").Find(&result).Error
		if errData != nil {
			errData = fmt.Errorf("get orders list data error : %w", errData)
			errChan <- errData
			return
		}
	}()

	go func() {
		defer wg.Done()
		errData := conn.WithContext(ctx).Debug().Table("orders").Where(where).Where("LOWER(invoice) like ?", "%"+strings.ToLower(param.Search)+"%").Count(&total).Error
		if errData != nil {
			errData = fmt.Errorf("get orders list count error : %w", errData)
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
