package user

import (
	"github.com/butga/paketin/errorModel"
	"github.com/gin-gonic/gin"
	"time"
)

type Service interface {
	Create(context *gin.Context)
	FindAll(context *gin.Context)
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

	UserDB, err := s.repository.FindByNamaKantor(userModel.NamaKantor)
	if UserDB.NamaKantor == userModel.NamaKantor {
		errorModel.GenerateDuplicateRegisterUser(context)
		return
	}

	err = s.repository.Create(userModel)
	if err != nil {
		errorModel.GenerateFailedAddData(context, "user")
		return
	}

	errorModel.GenerateNonErrorMessage(context, "Berhasil menambahkan data user")
	return
}

func (s *service) FindAll(context *gin.Context) {

	usersDB, err := s.repository.FindAll()
	if err != nil {
		errorModel.GenerateInternalServerError(context, "gagal mengambil semua data user")
		return
	}

	var usersResponse []UserResponse
	for _, u := range usersDB {
		user := convertModelToDTOOut(u)
		usersResponse = append(usersResponse, user)
	}

	errorModel.GenerateNonErrorMessageWithData(context, "Berhasil mengambil data user.", usersResponse)
	return
}

func readBody(context *gin.Context) (user UserRequest){
	err := context.ShouldBindJSON(&user)
	if err != nil {
		errorModel.GenerateFailedParsingRequest(context)
		return
	}

	return
}

func convertDTOToModel(user UserRequest) User{
	return User{
		ID: user.ID,
		Password: user.Password,
		NamaKantor: user.NamaKantor,
		NamaAdmin: user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor: user.NomorKantor,
		CreatedAt: time.Now(),
		Deleted: "false",
	}
}

func convertModelToDTOOut(user User) UserResponse {
	return UserResponse{
		NamaKantor:   user.NamaKantor,
		NamaAdmin:    user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor:  user.NomorKantor,
	}
}