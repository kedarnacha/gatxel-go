package models

import (
	"context"
	"time"
)

type AppoinmentSlot struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	AppoinmentID string    `json:"appoinment_ID"`
	Available    string    `json:"available"`
	StartTime    string    `json:"startTime`
	endTime      string    `json:"endTime`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Status       string    `json: "status"`
}

func (AppoinmentSlot) TableName() string {
	return "appoinmentSlot"
}

type AppoinmentSlotRepository interface {
	GetAllAppoinmentSlot(ctx context.Context) ([]*AppoinmentSlot, error)
	GetAppoinmentSlotByID(ctx context.Context, id int64) (*AppoinmentSlot, error)
	CreateAppoinmentSlotDay(ctx context.Context, appoinmentSlot *AppoinmentSlot) (*AppoinmentSlot, error)
	UpdateAppoinmentSlotByID(ctx context.Context, id int64, data map[string]interface{}) (*AppoinmentSlot, error)
	DeleteAppoinmentSlotByID(ctx context.Context, id int64) error
}
