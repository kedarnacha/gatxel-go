package handler

import (
	"net/http"
	"strconv"

	"github.com/kedarnacha/gatxel-go/helper"
	"github.com/kedarnacha/gatxel-go/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository models.UserRepository
}

func NewUserHandler(repository models.UserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}
func (h *UserHandler) GetAllUser(c *gin.Context) {
	user, err := h.repository.GetAllUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get data"))
	}
	c.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", user))
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	user, err := h.repository.GetUserByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get user"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Fetch data successfully", user))
}

func (h *UserHandler) UpdateUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}

	_, err = h.repository.GetUserByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helper.ResponseFailed("User not found"))
		return
	}

	var updateData models.User
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid request body"))
		return
	}
	data := map[string]interface{}{
		"username": updateData.Username,
		"email":    updateData.Email,
		"password": updateData.Password,
	}
	updatedUser, err := h.repository.UpdateUserByID(ctx, int64(id), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update user"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Update data successfully", updatedUser))
}

func (h *UserHandler) DeleteUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ResponseFailed("Invalid ID"))
		return
	}
	err = h.repository.DeleteUserByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete user"))
		return
	}
	ctx.JSON(http.StatusOK, helper.ResponseSuccess("Delete data successfully", nil))
}
