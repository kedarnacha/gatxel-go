package models

import (
	"context"
	"time"
)

type Notification struct {
	ID            int64     `json:"id"`
	UserID        int64     `json:"user_id"`
	AppointmentID int64     `json:"appointment_id"`
	Message       string    `json:"message"`
	CreatedAt     time.Time `json:"created_at"`
	IsSent        bool      `json:"is_sent"`
}

type NotificationRepository interface {
	GetAllNotification(ctx context.Context) ([]*Notification, error)
	GetNotificationByID(ctx context.Context, id int64) (*Notification, error)
	CreateNotification(ctx context.Context, category *Notification) (*Notification, error)
	UpdateNotificationByID(ctx context.Context, id int64, data map[string]interface{}) (*Notification, error)
	DeleteNotificationByID(ctx context.Context, id int64) error
}
