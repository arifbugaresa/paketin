package paket

import "gorm.io/gorm"

type Repository interface {
	Create(paket Paket) error
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

