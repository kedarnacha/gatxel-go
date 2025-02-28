package repository

import (
	"context"
	"gatxel-appointment/models"
)

type AppoinmentRepository interface {
	GetAllAppointments(ctx context.Context) ([]*models.Appoinment, error)
	GetAppointmentByID(ctx context.Context, id int64) (*models.Appoinment, error)
	CreateAppointment(ctx context.Context, appoinment *models.Appoinment) (*models.Appoinment, error)
	UpdateAppointmentByID(ctx context.Context, id int64, data map[string]interface{}) (*models.Appoinment, error)
	DeleteAppointmentByID(ctx context.Context, id int64) error
}
