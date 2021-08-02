package repository

import (
	"github.com/verryp/go-startup/app/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Store(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
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

func (r *userRepo) FindByEmail(email string) (user model.User, err error) {
	user = model.User{}
	err = r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return
	}

	return
}

func (r *userRepo) FindByID(ID int) (user model.User, err error) {
	user = model.User{}
	err = r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return
	}

	return
}

func (r *userRepo) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
