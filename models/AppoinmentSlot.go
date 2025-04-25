package models

import (
	"context"
	"time"
)

type AppoinmentSlot struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	AppoinmentID int64     `json:"appointment_id"`
	Available    bool      `json:"available"`
	StartTime    string    `json:"startTime" gorm:"column:start_time"`
	EndTime      string    `json:"endTime" gorm:"column:end_time"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (AppoinmentSlot) TableName() string {
	return "appoinment_slot"
}

type AppointmentSlotRepository interface {
	GetAllAppoinmentSlot(ctx context.Context) ([]*AppoinmentSlot, error)
	GetAppoinmentSlotByID(ctx context.Context, id int64) (*AppoinmentSlot, error)
	CreateAppoinmentSlotDay(ctx context.Context, appointmentSlot *AppoinmentSlot) (*AppoinmentSlot, error)
	UpdateAppoinmentSlotByID(ctx context.Context, id int64, data map[string]interface{}) (*AppoinmentSlot, error)
	DeleteAppoinmentSlotByID(ctx context.Context, id int64) error
}
