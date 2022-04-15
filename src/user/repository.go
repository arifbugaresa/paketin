package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) error
	FindByNamaKantor(namaKantor string) (User, error)
	FindAll() ([]User, error)
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

func (r *repository) FindAll() ([]User, error) {
	var allUserDB []User
	err := r.db.Where("deleted = ?", "false").Find(&allUserDB).Error

	return allUserDB, err
}