package handler

import (
	"github.com/butga/paketin/src/paket"
	"github.com/gin-gonic/gin"
)

type paketHandler struct {
	paketService paket.Service
}

func NewPaketHandler(service paket.Service) *paketHandler {
	return &paketHandler{service}
}

func (h *paketHandler) PostPaketHandler(context *gin.Context) {
	h.paketService.Create(context)
}