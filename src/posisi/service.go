package posisi

import (
	"github.com/butga/paketin/errorModel"
	"github.com/gin-gonic/gin"
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

	posisi := readBodyAndValidate(context)

	posisiModel := convertDTOToModel(posisi)
	posisiModel.CreatedAt = time.Now()

	err := s.repository.Create(posisiModel)
	if err != nil {
		errorModel.GenerateFailedAddData(context, "posisi")
		return
	}

	errorModel.GenerateNonErrorMessage(context, "Berhasil menambahkan data posisi")
	return
}

func readBodyAndValidate(context *gin.Context) (posisi PosisiRequest) {
	err := context.ShouldBindJSON(&posisi)
	if err != nil {
		errorModel.GenerateFailedParsingRequest(context)
		return
	}

	return
}

func convertDTOToModel(posisi PosisiRequest) Posisi {
	return Posisi{
		NomorResi:         posisi.NomorResi,
		LokasiSekarang:    posisi.LokasiSekarang,
		LokasiSelanjutnya: posisi.LokasiSelanjutnya,
	}
}