package service

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *authService) ValidateToken(token string) (*jwt.Token, error) {
	valToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return valToken, nil
}
