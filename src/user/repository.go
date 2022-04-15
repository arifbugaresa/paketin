package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) error
	FindByNamaKantor(namaKantor string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *repository) FindByNamaKantor(namaKantor string) (User, error) {
	var userDB User
	err := r.db.Where("nama_kantor = ?", namaKantor).First(&userDB).Error

	return userDB, err
}
