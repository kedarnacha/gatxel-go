package repository

import (
	"context"
	"fmt"

	"github.com/kedarnacha/gatxel-go/models"
	"gorm.io/gorm"
)

type AppoinmentDayRepository struct {
	db *gorm.DB
}

func NewAppoinmentDayRepository(db *gorm.DB) *AppoinmentDayRepository {
	return &AppoinmentDayRepository{db: db}
}

func (r *AppoinmentDayRepository) GetAllAppoinmentDay(ctx context.Context) ([]*models.AppoinmentDay, error) {
	fmt.Println("Querying table: appoinment_day")
	var appoinmentDay []*models.AppoinmentDay

	err := r.db.Table("appoinment_day").Find(&appoinmentDay).Error
	if err != nil {
		return nil, err
	}
	return appoinmentDay, nil
}

func (r *AppoinmentDayRepository) CreateAppoinmentDay(ctx context.Context, appoinmentDay *models.AppoinmentDay) (*models.AppoinmentDay, error) {
	if err := r.db.Create(appoinmentDay).Error; err != nil {
		return nil, err
	}
	return appoinmentDay, nil
}

func (r *AppoinmentDayRepository) GetAppoinmentDayByID(ctx context.Context, id int64) (*models.AppoinmentDay, error) {
	appoinmentDay := &models.AppoinmentDay{}
	if res := r.db.Model(appoinmentDay).Where("id = ?", id).First(appoinmentDay); res.Error != nil {
		return nil, res.Error
	}
	return appoinmentDay, nil
}

func (r *AppoinmentDayRepository) UpdateAppoinmentDayByID(ctx context.Context, id int64, data map[string]interface{}) (*models.AppoinmentDay, error) {
	appoinmentDay := &models.AppoinmentDay{}
	res := r.db.Model(&appoinmentDay).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return appoinmentDay, nil
}

func (r *AppoinmentDayRepository) DeleteAppoinmentDayByID(ctx context.Context, id int64) error {
	appoinmentDay := &models.AppoinmentDay{}
	res := r.db.Model(&appoinmentDay).Where("id = ?", id).Delete(appoinmentDay)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
