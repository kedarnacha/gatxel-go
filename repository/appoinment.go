package repository

import (
	"context"
	"fmt"

	"github.com/kedarnacha/gatxel-go/models"
	"gorm.io/gorm"
)

type AppoinmentRepository struct {
	db *gorm.DB
}

func NewAppoinmentRepository(db *gorm.DB) *AppoinmentRepository {
	return &AppoinmentRepository{db: db}
}

func (r *AppoinmentRepository) GetAllAppoinment(ctx context.Context) ([]*models.Appoinment, error) {
	fmt.Println("Querying table: appoinment")
	var appoinment []*models.Appoinment

	err := r.db.Table("appoinment").Find(&appoinment).Error
	if err != nil {
		return nil, err
	}
	return appoinment, nil
}

func (r *AppoinmentRepository) CreateAppoinment(ctx context.Context, appoinment *models.Appoinment) (*models.Appoinment, error) {
	if err := r.db.Create(appoinment).Error; err != nil {
		return nil, err
	}
	return appoinment, nil
}

func (r *AppoinmentRepository) GetAppoinmentByID(ctx context.Context, id int64) (*models.Appoinment, error) {
	appoinment := &models.Appoinment{}
	if res := r.db.Model(appoinment).Where("id = ?", id).First(appoinment); res.Error != nil {
		return nil, res.Error
	}
	return appoinment, nil
}

func (r *AppoinmentRepository) UpdateAppoinmentByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Appoinment, error) {
	if err := r.db.Model(&models.Appoinment{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return nil, err
	}

	appoinment := &models.Appoinment{}
	if err := r.db.First(appoinment, id).Error; err != nil {
		return nil, err
	}

	return appoinment, nil
}

func (r *AppoinmentRepository) DeleteAppoinmentByID(ctx context.Context, id int64) error {
	appoinment := &models.Appoinment{}
	res := r.db.Model(&appoinment).Where("id = ?", id).Delete(appoinment)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
