package models

import (
	"context"
	"time"
)

type Notification struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	UserID       int64     `json:"user_id" gorm:"index"`
	AppoinmentID int64     `json:"appoinment_id"`
	Message      string    `json:"message"`
	CreatedAt    time.Time `json:"created_at"`
	IsSent       bool      `json:"is_sent"`
}

// disini define table name
func (Notification) TableName() string {
	return "appoinment"
}

type NotificationRepository interface {
	GetAllNotification(ctx context.Context) ([]*Notification, error)
	GetNotificationByID(ctx context.Context, id int64) (*Notification, error)
	CreateNotification(ctx context.Context, category *Notification) (*Notification, error)
	UpdateNotificationByID(ctx context.Context, id int64, data map[string]interface{}) (*Notification, error)
	DeleteNotificationByID(ctx context.Context, id int64) error
}
