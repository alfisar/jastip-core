package repository

import (
	"fmt"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type travelSchRepository struct{}

func NewTravelSchRepository() *travelSchRepository {
	return &travelSchRepository{}
}

func (r *travelSchRepository) Create(conn *gorm.DB, data domain.TravelSchRequest) (id int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("traveler_schedule").Create(&data).Error
	if err != nil {
		err = fmt.Errorf("create traveler schedule error : %w", err)
		return
	}
	id = data.ID
	return
}

func (r *travelSchRepository) GetList(conn *gorm.DB, where map[string]any, search string, offet int, limit int) (result []domain.TravelSchResponse, count int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	if search == "" {
		err = conn.Debug().Table("traveler_schedule").Where(where).Count(&count).Offset(offet).Limit(limit).Find(&result).Error
	} else {
		err = conn.Debug().Table("traveler_schedule").Where(where).Where("locations like ?", "%"+search+"%").Count(&count).Offset(offet).Limit(limit).Find(&result).Error
	}

	if err != nil {
		err = fmt.Errorf("get traveler schedule error : %w", err)
		return
	}
	return
}

func (r *travelSchRepository) GetDetail(conn *gorm.DB, where map[string]any) (result domain.TravelSchResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("traveler_schedule").Where(where).First(&result).Error

	if err != nil {
		err = fmt.Errorf("get traveler schedule error : %w", err)
		return
	}
	return
}

func (r *travelSchRepository) GetByTimeBetween(conn *gorm.DB, id int, locations string, startDate string, endDate string) (result domain.TravelSchResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("traveler_schedule").Where("user_id = ? AND status = 1 and locations = ?", id, locations).Where("(period_start BETWEEN ? AND ?) OR (period_end BETWEEN ? AND ?)", startDate, endDate, startDate, endDate).First(&result).Error
	if err != nil {
		err = fmt.Errorf("get traveler schedule error : %w", err)
		return
	}
	return
}

func (r *travelSchRepository) Update(conn *gorm.DB, where map[string]any, updates map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	querys := conn.Debug().Table("traveler_schedule").Where(where).Updates(&updates)
	if querys.Error != nil {
		err = fmt.Errorf("Update trevel schedule error : %w", err)

	} else if querys.RowsAffected == 0 {
		err = fmt.Errorf("Update Failed : No Rows Affected")
	}

	return
}

func (r *travelSchRepository) Delete(conn *gorm.DB, where map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.Debug().Table("traveler_schedule").Where(where).Delete(domain.TravelSchRequest{}).Error
	if err != nil {
		err = fmt.Errorf("Delete travel schedule error : %w", err)

	}

	return
}
