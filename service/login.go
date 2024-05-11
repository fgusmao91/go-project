package service

import (
	"errors"
	"v1/domain"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (ls *LoginService) AuthenticateUser(credentials domain.Credentials) error {
	if credentials.Username != "admin" || credentials.Password != "admin" {
		return errors.New(domain.ErrInvalidCredentials)
	}

	return nil
}