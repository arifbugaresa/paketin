package user

import (
	"time"
)

type User struct {
	ID           int64
	NamaKantor   string
	NamaAdmin    string
	NomorKantor  string
	AlamatKantor string
	Password     string
	IsSuperAdmin bool `gorm:"default:false"`
	CreatedAt    time.Time
	Deleted      string `gorm:"default:false"`
	DeletedAt    time.Time
}
