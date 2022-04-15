package errorModel

import (
	"github.com/butga/paketin/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateNonErrorMessage(context *gin.Context, message string)  {
	context.JSON(http.StatusOK, gin.H{
		"paketin": src.Payload{
			Code:    200,
			Message: message,
			Data:    nil,
		},
	})
}

func GenerateNonErrorMessageWithData(context *gin.Context, message string, data interface{})  {
	context.JSON(http.StatusOK, gin.H{
		"paketin": src.Payload{
			Code:    200,
			Message: message,
			Data:    data,
		},
	})
}