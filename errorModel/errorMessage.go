package errorModel

import (
	"github.com/butga/paketin/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateInternalServerError(context *gin.Context, message string)  {
	context.JSON(http.StatusInternalServerError, gin.H{
		"paketin": src.Payload{
			Code:    500,
			Message: "Terjadi kesalahan sistem, " + message,
			Data:    nil,
		},
	})
}

func GenerateRequestIDError(context *gin.Context, ID string)  {
	context.JSON(http.StatusInternalServerError, gin.H{
		"paketin": src.Payload{
			Code:    500,
			Message: "Terjadi kesalahan sistem, " + ID + " tidak sesuai.",
			Data:    nil,
		},
	})
}

func GenerateDuplicateRegisterUser(context *gin.Context)  {
	context.JSON(http.StatusBadRequest, gin.H{
		"paketin": src.Payload{
			Code:    400,
			Message: "Terjadi kesalahan request, kantor telah terdaftar dalam sistem.",
			Data:    nil,
		},
	})
}

func GenerateFailedAddData(context *gin.Context, message string)  {
	context.JSON(http.StatusBadRequest, gin.H{
		"paketin": src.Payload{
			Code:    400,
			Message: "Terjadi kesalahan request, data " + message + " gagal ditambahkan.",
			Data:    nil,
		},
	})
}

func GenerateFailedParsingRequest(context *gin.Context)  {
	context.JSON(http.StatusBadRequest, gin.H{
		"ajarin": src.Payload{
			Code:    400,
			Message: "Terjadi kesalahan request, gagal parsing request",
			Data:    nil,
		},
	})
}