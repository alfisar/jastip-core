package repository

import (
	"context"
	"fmt"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type orderDetailRepository struct{}

func NewOrderDetailsRepository() *orderDetailRepository {
	return &orderDetailRepository{}
}

func (r *orderDetailRepository) CreateBulk(ctx context.Context, conn *gorm.DB, data []domain.OrderDetail) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("orders_detail").Create(&data).Error
	return
}
