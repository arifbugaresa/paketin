package user

import (
	"github.com/butga/paketin/src"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service interface {
	Create(context *gin.Context)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(context *gin.Context) {

	user := readBody(context)

	userModel := convertDTOToModel(user)

	UserDB, err := s.repository.FindByUsername(userModel.Username)
	if UserDB.Username == userModel.Username {
		context.JSON(http.StatusBadRequest, gin.H{
			"ajarin": src.Payload{
				Code:    400,
				Message: "Username telah digunakan",
				Data:    nil,
			},
		})
		return
	}

	err = s.repository.Create(userModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ajarin": src.Payload{
				Code:    400,
				Message: "Gagal menambahkan data user",
				Data:    nil,
			},
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ajarin": src.Payload{
			Code:    200,
			Message: "Berhasil menambahkan data user",
			Data:    nil,
		},
	})
	return
}

func readBody(context *gin.Context) (user UserRequest){
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ajarin": src.Payload{
				Code:    400,
				Message: "Gagal parsing request",
				Data:    nil,
			},
		})
	}

	return
}

func convertDTOToModel(user UserRequest) User{
	return User{
		ID: user.ID,
		Username: user.Username,
		Password: user.Password,
		NamaKantor: user.NamaKantor,
		NamaAdmin: user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor: user.NomorKantor,
		CreatedAt: time.Now(),
		Deleted: "false",
	}
}
