package handler

import (
	"github.com/butga/paketin/src/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) PostUserHandler(context *gin.Context) {

	// Service
	err := h.userService.Create(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Berhasil menambahkan data user",
	})
}


