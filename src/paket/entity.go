package paket

import "time"

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
	DeliveredAt   time.Time
	CreatedAt     time.Time
}
