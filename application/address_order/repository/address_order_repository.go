package repository

import (
	"context"
	"fmt"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type addressOrderRepository struct{}

func NewAddressOrderRepository() *addressOrderRepository {
	return &addressOrderRepository{}
}

func (r *addressOrderRepository) Create(ctx context.Context, conn *gorm.DB, data domain.AddressOrder) (id int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}
	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("address_order").Save(&data).Error
	id = data.ID
	return
}
