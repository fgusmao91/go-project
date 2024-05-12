package service

import (
	"crypto/sha256"
	"fmt"
	"time"
	"v1/domain"
	"v1/dto"
	"v1/repository"

	"github.com/dgrijalva/jwt-go"
)

const ErrInvalidCredentials = "invalid credentials"

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (ls *AuthService) AuthenticateUser(credentials dto.Credentials) (string, error) {
	hashPassword := hashPassword(credentials.Password)
	_, err := ls.authRepository.GetCredentials(credentials.Username, hashPassword)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ls *AuthService) RegisterUser(credentials dto.Credentials) error {
	domainCredentials := domain.Credentials{
		Username: credentials.Username,
		Password: hashPassword(credentials.Password),
	}
	err := ls.authRepository.InsertCredentials(domainCredentials)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}