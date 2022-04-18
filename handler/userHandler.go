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

func (h *userHandler) GetUsersHandler(context *gin.Context) {
	h.userService.FindAll(context)
}

func (h *userHandler) GetUserHandler(context *gin.Context) {
	h.userService.Find(context)
}

func (h *userHandler) DeleteUsersHandler(context *gin.Context) {
	h.userService.Delete(context)
}