package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"app_name" : "paketin",
		"version" : "1.0.1",
	})
}

func Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
