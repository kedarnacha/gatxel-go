package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"
)

type NotificationHandler struct {
	repository models.NotificationRepository
}

func NewNotificationHandler(repository models.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repository: repository}
}

func (h *NotificationHandler) GetAllNotifications(c *gin.Context) {
	notification, err := h.repository.GetAllNotification(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}
	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", notification))
}

func (h *NotificationHandler) CreateNotification(ctx *gin.Context) {
	notification := &models.Notification{}
	if err := ctx.ShouldBindJSON(notification); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Payload invalid"))
		return
	}
	notification, err := h.repository.CreateNotification(ctx, notification)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to create notification"))
		return
	}
	ctx.JSON(http.StatusCreated, helper.ResponseSuccess("Create data successfully", notification))
}

func (h *NotificationHandler) GetNotificationByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	notification, err := h.repository.GetNotificationByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get notification"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", notification))
}

func (h *NotificationHandler) UpdateNotificationByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}

	_, err = h.repository.GetNotificationByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ResponseFailed("Notification not found"))
		return
	}

	var updateData models.Notification
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request body"))
		return
	}
	data := map[string]interface{}{
		"message": updateData.Message,
		"is_sent": updateData.IsSent,
	}
	updatedNotification, err := h.repository.UpdateNotificationByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update notification"))
		return
	}

	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfully", updatedNotification))
}

func (h *NotificationHandler) DeleteNotificationByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	err = h.repository.DeleteNotificationByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete notification"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfully", nil))
}
