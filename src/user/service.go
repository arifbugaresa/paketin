package user

import (
	"github.com/butga/paketin/errorModel"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

type Service interface {
	Create(context *gin.Context)
	FindAll(context *gin.Context)
	Delete(context *gin.Context)
	Find(context *gin.Context)
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

	errorModel.GenerateNonErrorMessageWithData(context, "Berhasil mengambil seluruh data user.", usersResponse)
	return
}

func (s *service) Delete(context *gin.Context) {

	idString := context.Param("id")
	id, err := strconv.Atoi(idString)

	// Validation Get User On DB
	usersDB, err := s.repository.FindByID(id)
	if err != nil {
		errorModel.GenerateRequestIDError(context, "user id")
		return
	} else if usersDB.ID < 1 {
		errorModel.GenerateInternalServerError(context, "user id tidak ditemukan")
		return
	} else {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		usersDB.NamaKantor += strconv.Itoa(r1.Intn(100000))
		usersDB.Deleted = "true"

		s.repository.Save(usersDB)
	}

	errorModel.GenerateNonErrorMessage(context, "Berhasil menghapus data")
	return
}

func (s *service) Find(context *gin.Context) {

	idString := context.Param("id")
	id, err := strconv.Atoi(idString)

	// Validation Get User On DB
	usersDB, err := s.repository.FindByID(id)
	if err != nil {
		errorModel.GenerateRequestIDError(context, "user id")
		return
	} else if usersDB.ID < 1 {
		errorModel.GenerateInternalServerError(context, "user id tidak ditemukan")
		return
	}

	user := convertModelToViewDTOOut(usersDB)

	errorModel.GenerateNonErrorMessageWithData(context, "Berhasil mengambil data user", user)
	return
}

func readBody(context *gin.Context) (user UserRequest) {
	err := context.ShouldBindJSON(&user)
	if err != nil {
		errorModel.GenerateFailedParsingRequest(context)
		return
	}

	return
}

func convertDTOToModel(user UserRequest) User {
	return User{
		ID:           user.ID,
		Password:     user.Password,
		NamaKantor:   user.NamaKantor,
		NamaAdmin:    user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor:  user.NomorKantor,
		CreatedAt:    time.Now(),
		Deleted:      "false",
	}
}

func convertModelToDTOOut(user User) UserResponse {
	return UserResponse{
		ID:           user.ID,
		NamaKantor:   user.NamaKantor,
		NomorKantor:  user.NomorKantor,
	}
}

func convertModelToViewDTOOut(user User) DetailUserResponse {
	return DetailUserResponse{
		ID:           user.ID,
		NamaKantor:   user.NamaKantor,
		NamaAdmin:    user.NamaAdmin,
		AlamatKantor: user.AlamatKantor,
		NomorKantor:  user.NomorKantor,
	}
}
