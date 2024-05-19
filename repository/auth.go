package repository

import (
	"database/sql"
	"errors"
	"v1/domain"
)

const (
	insertCredentials = "INSERT INTO credentials (username, password) VALUES (?, ?)"
	InsertAuthorizations = "INSERT INTO authorizations (credential_id, app_name, auth_type) VALUES (?, ?, ?)"
	getCredentials    = "SELECT username, password FROM credentials WHERE username = ? and password = ?"
	getCredentialIDByUserName = "SELECT id FROM credentials WHERE username = ?"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (lr *AuthRepository) InsertCredentials(credentials domain.Credentials) (*int64,error) {
	result, err := lr.db.Exec(insertCredentials, credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}

	lasteInsertedID, _ := result.LastInsertId()

	return &lasteInsertedID, nil
}

func (lr *AuthRepository) InsertAuthorizations(authorizations domain.Authorizations) (*int64,error) {
	result, err := lr.db.Exec(InsertAuthorizations, authorizations.CredentialID, authorizations.AppName, authorizations.AuthType)
	if err != nil {
		return nil, err
	}

	lasteInsertedID, _ := result.LastInsertId()

	return &lasteInsertedID, nil
}

func (lr *AuthRepository) GetCredentials(username string, password string) (domain.Credentials, error) {
	var credentials domain.Credentials
	err := lr.db.QueryRow(getCredentials, username, password).Scan(&credentials.Username, &credentials.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Credentials{}, errors.New("user not found")
		}
		return domain.Credentials{}, err
	}
	
	return credentials, nil
}

func (lr *AuthRepository) GetCredentialIDByUsername(username string) (int64, error) {
	var credentialID int64
	err := lr.db.QueryRow(getCredentialIDByUserName, username).Scan(&credentialID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}
	
	return credentialID, nil
}