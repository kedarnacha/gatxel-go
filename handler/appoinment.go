package handler

import (
	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppoinmentHandler struct {
	repository models.AppoinmentRepository
}

func NewAppoinmentHandler(repository models.AppoinmentRepository) *AppoinmentHandler {
	return &AppoinmentHandler{repository: repository}
}

func (h *AppoinmentHandler) GetAllAppoinments(c *gin.Context) {
	appoinments, err := h.repository.GetAllAppoinments(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
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
	appointment, err := h.repository.GetAppointmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get appointment"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", appointment))
}

func (h *AppoinmentHandler) UpdateAppoinmentByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	appoinment, err := h.repository.GetAppoinmentByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get appoinment"))
		return
	}

	updateData := models.Appoinment{}
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
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete appointment"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfully", nil))
}
