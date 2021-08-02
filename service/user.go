package service

import (
	"github.com/verryp/go-startup/app/input"
	"github.com/verryp/go-startup/app/model"
	"github.com/verryp/go-startup/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input input.UserRegisterInput) (model.User, error)
	IsEmailAvailable(input input.EmailRequest) (bool, error)
}

type userService struct {
	repository repository.UserRepo
}

func NewUserService(r repository.UserRepo) *userService {
	return &userService{r}
}

func (s *userService) RegisterUser(input input.UserRegisterInput) (model.User, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passHash)
	user.Role = "user"

	newUser, err := s.repository.Store(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) IsEmailAvailable(req input.EmailRequest) (bool, error) {
	email := req.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}
