package paket

type Paket struct {
	ID            int64 `gorm:"column:id;primaryKey;autoIncrement"`
	NomorResi     string
	Produk        string
	BeratPaket    int64
	NamaPengirim  string
	NamaPenerima  string
	NomorPenerima string
	AlamatTujuan  string
	Status        string `gorm:"default:dikemas"`
}

//func (p Paket) ValidateStatus(status string) {
//
//	if (status != "dikemas") || (status != "dikirim") || (status != "selesai"){
//		errorModel.GenerateRequestStatusError(con)
//	}
//}
