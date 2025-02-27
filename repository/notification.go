package repository

import (
	"context"

	"github.com/kedarnacha/gatxel-go/models"
)

type NotificationRepository interface {
	GetAllNotifications(ctx context.Context) ([]*models.Notification, error)
	GetNotificationByID(ctx context.Context, id int64) (*models.Notification, error)
	CreateNotification(ctx context.Context, notification *models.Notification) (*models.Notification, error)
	UpdateNotificationByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Notification, error)
	DeleteNotificationByID(ctx context.Context, id int64) error
}
