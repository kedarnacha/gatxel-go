package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kedarnacha/gatxel-go/models"
	"github.com/kedarnacha/gatxel-go/repository"
)

type AppoinmentSlotHandler struct {
	Repo *repository.AppoinmentSlotRepository
}

func NewAppoinmentSlotHandler(repo *repository.AppoinmentSlotRepository) *AppoinmentSlotHandler {
	return &AppoinmentSlotHandler{Repo: repo}
}

func (h *AppoinmentSlotHandler) GetAllAppoinmentSlot(c *gin.Context) {
	data, err := h.Repo.GetAllAppoinmentSlot(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *AppoinmentSlotHandler) CreateAppoinmentSlot(c *gin.Context) {
	var input models.AppoinmentSlot
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.Repo.CreateAppoinmentSlot(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (h *AppoinmentSlotHandler) GetAppoinmentSlotByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	data, err := h.Repo.GetAppoinmentSlotByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *AppoinmentSlotHandler) UpdateAppoinmentSlotByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if val, ok := data["startTime"]; ok {
		data["start_time"] = val
		delete(data, "startTime")
	}
	if val, ok := data["endTime"]; ok {
		data["end_time"] = val
		delete(data, "endTime")
	}

	result, err := h.Repo.UpdateAppoinmentSlotByID(c.Request.Context(), id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *AppoinmentSlotHandler) DeleteAppoinmentSlotByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := h.Repo.DeleteAppoinmentSlotByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
