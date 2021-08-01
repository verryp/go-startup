package repository

import (
	"github.com/verryp/go-startup/app/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Store(user model.User) (model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) Store(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}