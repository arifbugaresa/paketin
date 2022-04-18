package paket

import "gorm.io/gorm"

type Repository interface {
	Create(paket Paket) error
	FindByNoResi(noResi string) (Paket, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(paket Paket) error {
	err := r.db.Create(&paket).Error

	return err
}

func (r *repository) FindByNoResi(noResi string) (Paket, error) {
	var paket Paket
	err := r.db.Where("nomor_resi = ?", noResi).First(&paket).Error

	return paket, err
}

