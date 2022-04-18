package paket

import (
	"github.com/butga/paketin/errorModel"
	"github.com/gin-gonic/gin"
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

func readBody(context *gin.Context) (paket PaketRequest) {
	err := context.ShouldBindJSON(&paket)
	if err != nil {
		errorModel.GenerateFailedParsingRequest(context)
		return
	}

	return
}

func convertDTOToModel(paket PaketRequest) Paket {
	return Paket{
		Produk:        paket.Produk,
		BeratPaket:    paket.BeratPaket,
		NamaPengirim:  paket.NamaPengirim,
		NamaPenerima:  paket.NamaPenerima,
		NomorPenerima: paket.NomorPenerima,
		AlamatTujuan:  paket.AlamatTujuan,
	}
}

func (s *service) Create(context *gin.Context) {

	paket := readBody(context)

	paketModel := convertDTOToModel(paket)

	//nomorResi := generateNomorResi()

	//paketOnDb, err := s.repository.FindByNoResi(nomorResi)
	//if paketOnDb.ID == userModel.NamaKantor {
	//	errorModel.GenerateDuplicateRegisterUser(context)
	//	return
	//}

	err := s.repository.Create(paketModel)
	if err != nil {
		errorModel.GenerateFailedAddData(context, "paket")
		return
	}

	errorModel.GenerateNonErrorMessage(context, "Berhasil menambahkan data paket")
	return
}


