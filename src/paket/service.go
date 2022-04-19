package paket

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
	Find(context *gin.Context)
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

func (s *service) FindAll(context *gin.Context) {

	listPaketDB, err := s.repository.FindAll()
	if err != nil {
		errorModel.GenerateInternalServerError(context, "gagal mengambil semua data paket")
		return
	}

	var listPaketResponse []PaketResponse
	for _, p := range listPaketDB {
		paketResponse := convertModelToDTOOut(p)
		listPaketResponse = append(listPaketResponse, paketResponse)
	}

	errorModel.GenerateNonErrorMessageWithData(context, "Berhasil mengambil seluruh data user.", listPaketResponse)
	return
}

func (s *service) Create(context *gin.Context) {

	paket := readBody(context)

	paketModel := convertDTOToModel(paket)

	paketModel.NomorResi = s.generateNomorResi()
	paketModel.CreatedAt = time.Now()

	err := s.repository.Create(paketModel)
	if err != nil {
		errorModel.GenerateFailedAddData(context, "paket")
		return
	}

	errorModel.GenerateNonErrorMessage(context, "Berhasil menambahkan data paket")
	return
}

func (s *service) Find(context *gin.Context) {

	idString := context.Param("id")
	id, err := strconv.Atoi(idString)

	// Validation Get User On DB
	paketDB, err := s.repository.FindByID(id)
	if err != nil {
		errorModel.GenerateRequestIDError(context, "id paket")
		return
	} else if paketDB.ID < 1 {
		errorModel.GenerateInternalServerError(context, "id paket tidak ditemukan")
		return
	}

	paket := convertModelToViewDTOOut(paketDB)

	errorModel.GenerateNonErrorMessageWithData(context, "Berhasil mengambil data paket", paket)
	return
}

func convertModelToViewDTOOut(paket Paket) PaketDetaiResponse {
	return PaketDetaiResponse{
		ID:            paket.ID,
		NomorResi:     paket.NomorResi,
		Produk:        paket.Produk,
		BeratPaket:    paket.BeratPaket,
		NamaPengirim:  paket.NamaPengirim,
		NamaPenerima:  paket.NomorPenerima,
		NomorPenerima: paket.NomorPenerima,
		AlamatTujuan:  paket.AlamatTujuan,
		Status:        paket.Status,
	}
}

func (s *service) generateNomorResi() string {

	var nomorResi string

	isUnix := false
	for isUnix == false {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomInteger := r1.Intn(10000000)

		nomorResi = "PKTN" + strconv.Itoa(randomInteger)

		paket, _ := s.repository.FindByNoResi(nomorResi)
		if paket.ID == 0 {
			isUnix = true
		}
	}

	return nomorResi
}

func convertModelToDTOOut(paket Paket) PaketResponse {
	return PaketResponse{
		ID:        paket.ID,
		NomorResi: paket.NomorResi,
		Produk:    paket.Produk,
	}
}
