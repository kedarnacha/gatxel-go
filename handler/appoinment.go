package handler

import (
	"net/http"
	"strconv"

	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"

	"github.com/gin-gonic/gin"
)

type AppoinmentHandler struct {
	repository models.AppoinmentRepository
}

func NewAppoinmentHandler(repository models.AppoinmentRepository) *AppoinmentHandler {
	return &AppoinmentHandler{repository: repository}
}

func (h *AppoinmentHandler) GetAllAppoinment(c *gin.Context) {
	appoinments, err := h.repository.GetAllAppoinment(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", appoinments))
}

func (h *AppoinmentHandler) CreateAppointment(ctx *gin.Context) {
	appoinment := &models.Appoinment{}
	if err := ctx.ShouldBindJSON(appoinment); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}

	appoinment, err := h.repository.CreateAppoinment(ctx, appoinment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create appointment"))
		return
	}

	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfully", appoinment))
}

func (h *AppoinmentHandler) GetAppointmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}

	appointment, err := h.repository.GetAppoinmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ResponseFailed("Appointment not found"))
		return
	}

	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", appointment))
}

func (h *AppoinmentHandler) UpdateAppointmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}

	_, err = h.repository.GetAppoinmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ResponseFailed("Appointment not found"))
		return
	}

	var updateData models.Appoinment
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request body"))
		return
	}

	data := map[string]interface{}{
		"title":       updateData.Title,
		"description": updateData.Description,
		"start_time":  updateData.StartTime,
		"end_time":    updateData.EndTime,
	}

	updatedAppoinment, err := h.repository.UpdateAppoinmentByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update appointment"))
		return
	}

	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfully", updatedAppoinment))
}

func (h *AppoinmentHandler) DeleteAppoinmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}

	err = h.repository.DeleteAppoinmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ResponseFailed("Appoinment not found"))
		return
	}

	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfully", nil))
}
