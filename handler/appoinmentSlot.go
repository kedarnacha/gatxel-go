package handler

import (
	"context"
	"fmt"

	"github.com/kedarnacha/gatxel-go/models"
	"gorm.io/gorm"
)

type AppoinmentSlotRepository struct {
	db *gorm.DB
}

func NewAppoinmentSlotRepository(db *gorm.DB) *AppoinmentSlotRepository {
	return &AppoinmentSlotRepository{db: db}
}

func (r *AppoinmentSlotRepository) GetAllAppoinmentSlot(ctx context.Context) ([]*models.AppoinmentSlot, error) {
	fmt.Println("Querying table: appoinmentSlot")
	var appoinmentSlot []*models.AppoinmentSlot

	err := r.db.Table("appoinmentSlot").Find(&appoinmentSlot).Error
	if err != nil {
		return nil, err
	}
	return appoinmentSlot, nil
}

func (r *AppoinmentSlotRepository) CreateAppoinmentSlot(ctx context.Context, appoinmentSlot *models.AppoinmentSlot) (*models.AppoinmentSlot, error) {
	if err := r.db.Create(appoinmentSlot).Error; err != nil {
		return nil, err
	}
	return appoinmentSlot, nil
}

func (r *AppoinmentSlotRepository) GetAppoinmentSlotByID(ctx context.Context, id int64) (*models.AppoinmentSlot, error) {
	appoinmentSlot := &models.AppoinmentSlot{}
	if res := r.db.Model(appoinmentSlot).Where("id = ?", id).First(appoinmentSlot); res.Error != nil {
		return nil, res.Error
	}
	return appoinmentSlot, nil
}
func (r *AppoinmentSlotRepository) UpdateAppoinmentSlotByID(ctx context.Context, id int64, data map[string]interface{}) (*models.AppoinmentSlot, error) {
	appoinmentSlot := &models.AppoinmentSlot{}
	res := r.db.Model(&appoinmentSlot).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return appoinmentSlot, nil
}

func (r *AppoinmentSlotRepository) DeleteAppoinmentSlotByID(ctx context.Context, id int64) error {
	appoinmentSlot := &models.AppoinmentSlot{}
	res := r.db.Model(&appoinmentSlot).Where("id = ?", id).Delete(appoinmentSlot)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
