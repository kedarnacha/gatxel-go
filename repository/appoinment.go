package repository

import (
	"context"
	"gatxel-appointment/models"
)

type AppoinmentRepository interface {
	GetAllAppointments(ctx context.Context) ([]*models.Appointment, error)
	GetAppointmentByID(ctx context.Context, id int64) (*models.Appointment, error)
	CreateAppointment(ctx context.Context, appointment *models.Appointment) (*models.Appointment, error)
	UpdateAppointmentByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Appointment, error)
	DeleteAppointmentByID(ctx context.Context, id int64) error
}
