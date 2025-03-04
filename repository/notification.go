package repository

import (
	"context"

	"github.com/kedarnacha/gatxel-go/models"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) GetAllNotification(ctx context.Context) ([]*models.Notification, error) {
	var notification []*models.Notification
	err := r.db.Model(&models.Notification{}).Find(&notification).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) CreateNotification(ctx context.Context, notification *models.Notification) (*models.Notification, error) {
	if err := r.db.Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) GetNotificationByID(ctx context.Context, id int64) (*models.Notification, error) {
	notification := &models.Notification{}
	if res := r.db.Model(notification).Where("id = ?", id).First(notification); res.Error != nil {
		return nil, res.Error
	}
	return notification, nil
}

func (r *NotificationRepository) UpdateNotificationByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Notification, error) {
	notification := &models.Notification{}
	res := r.db.Model(&notification).Where("id = ?", id).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return notification, nil
}

func (r *NotificationRepository) DeleteNotificationByID(ctx context.Context, id int64) error {
	notification := &models.Notification{}
	res := r.db.Model(&notification).Where("id = ?", id).Delete(notification)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
