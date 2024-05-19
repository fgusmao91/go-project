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

func (ls *AuthService) AuthenticateUser(credentials dto.Login) (string, error) {
	hashPassword := hashPassword(credentials.Password)
	domainCredential, err := ls.authRepository.GetCredentials(credentials.Username, hashPassword)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = credentials.Username
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	domainAuthorizations, err := ls.authRepository.GetAuthorizations(domainCredential.ID)
	if err != nil {
		return "", err
	}
	
	var authorizations []dto.Authorization
	for _, authorization := range domainAuthorizations {
		auth := dto.Authorization{
			AppName: authorization.AppName,
			AuthType: authorization.AuthType,
		}
		authorizations = append(authorizations, auth)
	}

	claims["authorizations"] = authorizations

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ls *AuthService) RegisterUser(credentials dto.Login) error {
	domainCredentials := domain.Credentials{
		Username: credentials.Username,
		Password: hashPassword(credentials.Password),
	}
	_, err := ls.authRepository.InsertCredentials(domainCredentials)
	if err != nil {
		return err
	}

	return nil
}

func (ls *AuthService) AddAuthorization(userName string, authorizations dto.Authorization) error {
	credentialID, err := ls.authRepository.GetCredentialIDByUsername(userName)
	if err != nil {
		return err
	}
	
	domainAuthorizations := domain.Authorizations{
		CredentialID: credentialID,
		AppName: authorizations.AppName,
		AuthType: authorizations.AuthType,
	}
	_, err = ls.authRepository.InsertAuthorizations(domainAuthorizations)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}