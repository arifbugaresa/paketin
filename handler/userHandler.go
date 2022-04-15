package handler

import (
	"github.com/butga/paketin/src/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) PostUserHandler(context *gin.Context) {
	h.userService.Create(context)
}