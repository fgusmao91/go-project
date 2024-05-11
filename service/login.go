package service

import (
	"errors"
	"time"
	"v1/domain"

	"github.com/dgrijalva/jwt-go"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (ls *LoginService) AuthenticateUser(credentials domain.Credentials) (string, error) {
	if credentials.Username != "fgusmao" || credentials.Password != "admin" {
		return "",errors.New(domain.ErrInvalidCredentials)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}