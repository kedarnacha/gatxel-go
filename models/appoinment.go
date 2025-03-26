package models

import (
	"context"
	"time"
)

type Appoinment struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	UserID      int64     `json:"user_id" gorm:"index"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Appoinment) TableName() string {
	return "appoinment"
}

type AppoinmentRepository interface {
	GetAllAppoinment(ctx context.Context) ([]*Appoinment, error)
	GetAppoinmentByID(ctx context.Context, id int64) (*Appoinment, error)
	CreateAppoinment(ctx context.Context, appoinment *Appoinment) (*Appoinment, error)
	UpdateAppoinmentByID(ctx context.Context, id int64, data map[string]interface{}) (*Appoinment, error)
	DeleteAppoinmentByID(ctx context.Context, id int64) error
}
