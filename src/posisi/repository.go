package posisi

import "gorm.io/gorm"

type Repository interface {
	Create(posisi Posisi) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(posisi Posisi) error {
	err := r.db.Create(&posisi).Error

	return err
}