package handler

import (
	"github.com/butga/paketin/src/posisi"
	"github.com/gin-gonic/gin"
)

type posisiHandler struct {
	posisiService posisi.Service
}

func NewPosisiHandler(service posisi.Service) *posisiHandler {
	return &posisiHandler{service}
}

func (h *posisiHandler) PostPosisiHandler(context *gin.Context) {
	h.posisiService.Create(context)
}