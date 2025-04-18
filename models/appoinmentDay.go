package models

import (
	"context"
	"time"
)

type AppoinmentDay struct {
	ID          int64     `json:"id" gorm:"primaryKey"`
	AppinmentID string    `json:"appoinment_ID"`
	DayOfWeek   string    `json:"Day_Of_Week"	`
	startTime   string    `json: "starttime"`
	endTime     string    `json: "endtime"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Status      string    `json: "status"`
}

func (AppoinmentDay) TableName() string {
	return "appoinmentDay"
}

type AppoinmentDayRepository interface {
	GetAllAppoinmentDay(ctx context.Context) ([]*AppoinmentDay, error)
	GetAppoinmentDayByID(ctx context.Context, id int64) (*AppoinmentDay, error)
	CreateAppoinmentDay(ctx context.Context, appoinmentDay *AppoinmentDay) (*AppoinmentDay, error)
	UpdateAppoinmentDayByID(ctx context.Context, id int64, data map[string]interface{}) (*AppoinmentDay, error)
	DeleteAppoinmentDayByID(ctx context.Context, id int64) error
}
