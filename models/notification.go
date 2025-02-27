package models

import (
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
