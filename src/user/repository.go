package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) error
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
