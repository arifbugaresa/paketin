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
	CreatedAt    time.Time
	Deleted      string
}
