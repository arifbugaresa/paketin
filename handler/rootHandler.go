package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
