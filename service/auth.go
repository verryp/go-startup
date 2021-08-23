package service

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
}

type authService struct {
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) GenerateToken(userID int) (token string, err error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	genToken := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	token, err = genToken.SignedString(SECRET_KEY)
	if err != nil {
		return
	}

	return
}
