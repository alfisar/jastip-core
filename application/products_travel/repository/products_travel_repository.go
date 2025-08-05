package repository

import (
	"context"
	"fmt"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productTravelRepository struct {
}

func NewProductTravelrepository() *productTravelRepository {
	return &productTravelRepository{}
}

func (r *productTravelRepository) CreateBulk(ctx context.Context, conn *gorm.DB, data []domain.ProductsTravel) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("product_travel").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("create product travel error : %w", err)
		return
	}

	return
}

func (r *productTravelRepository) DeleteBulk(ctx context.Context, conn *gorm.DB, where map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	whereOther := clause.Expr{}
	whereOne := make(map[string]any)
	for key, v := range where {
		if key == "IN" {
			whereOther = gorm.Expr("product_id IN (?)", v)
		} else {
			whereOne[key] = v
		}
	}
	query := conn.WithContext(ctx).Debug().Table("product_travel").Where(whereOne)
	if whereOther.SQL != "" {
		query = query.Where(whereOther)
	}

	err = query.Delete(&domain.ProductsTravel{}).Error
	if err != nil {
		err = fmt.Errorf("delete product travel error : %w", err)
		return
	}

	return
}

func (r *productTravelRepository) GetExpr(ctx context.Context, conn *gorm.DB, where clause.Expr) (result []domain.ProductsTravel, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Debug().Table("product_travel").Where(where).Find(&result).Error
	if err != nil {
		err = fmt.Errorf("get expr product travel error : %w", err)
		return
	}

	return
}
