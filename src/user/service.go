package user

import (
	"github.com/butga/paketin/src/user/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service interface {
	Create(context *gin.Context) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(context *gin.Context) error {

	user := readBody(context)

	userModel := convertDTOToModel(user)

	err := s.repository.Create(userModel)

	return err
}

func readBody(context *gin.Context) (user dto.UserRequest){
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": "error readBody",
		})
	}

	return
}

func convertDTOToModel(user dto.UserRequest) User{
	return User{
		ID: user.ID,
		NamaKantor: user.NamaKantor,
		NamaAdmin: user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor: user.NomorKantor,
		CreatedAt: time.Now(),
		Deleted: "false",
	}
}
