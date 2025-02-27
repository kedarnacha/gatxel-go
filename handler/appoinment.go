package handler

import (
	"gatxel-appointment/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	repository models.AppointmentRepository
}

func NewAppointmentHandler(repository models.AppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{repository: repository}
}

func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	appointments, err := h.repository.GetAllAppointments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}
	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", appointments))
}

func (h *AppointmentHandler) CreateAppointment(ctx *gin.Context) {
	appointment := &models.Appointment{}
	if err := ctx.ShouldBindJSON(appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}
	appointment, err := h.repository.CreateAppointment(ctx, appointment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create appointment"))
		return
	}
	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfully", appointment))
}

func (h *AppointmentHandler) GetAppointmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	appointment, err := h.repository.GetAppointmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get appointment"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", appointment))
}

func (h *AppointmentHandler) UpdateAppointmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	appointment, err := h.repository.GetAppointmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get appointment"))
		return
	}

	updateData := models.Appointment{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid payload"))
		return
	}

	data := map[string]interface{}{
		"title":       updateData.Title,
		"description": updateData.Description,
		"start_time":  updateData.StartTime,
		"end_time":    updateData.EndTime,
	}
	updatedAppointment, err := h.repository.UpdateAppointmentByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update appointment"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfully", updatedAppointment))
}

func (h *AppointmentHandler) DeleteAppointmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	err = h.repository.DeleteAppointmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete appointment"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfully", nil))
}
