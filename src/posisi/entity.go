package posisi

import "time"

type Posisi struct {
	ID                int64 `gorm:"column:id;primaryKey;autoIncrement"`
	NomorResi         string
	LokasiSekarang    string
	LokasiSelanjutnya string
	CreatedAt         time.Time
}
