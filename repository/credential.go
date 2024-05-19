package repository

import (
	"database/sql"
	"errors"
	"v1/domain"
)

const (
	insertCredentials = "INSERT INTO credentials (username, password) VALUES (?, ?)"
	getCredentials    = "SELECT id, username, password FROM credentials WHERE username = ? and password = ?"
	getCredentialIDByUserName = "SELECT id FROM credentials WHERE username = ?"
)

type CredentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) *CredentialRepository {
	return &CredentialRepository{
		db: db,
	}
}

func (lr *CredentialRepository) InsertCredentials(credentials domain.Credentials) (*int64,error) {
	result, err := lr.db.Exec(insertCredentials, credentials.Username, credentials.Password)
	if err != nil {
		return nil, err
	}

	lasteInsertedID, _ := result.LastInsertId()

	return &lasteInsertedID, nil
}

func (lr *CredentialRepository) GetCredentials(username string, password string) (domain.Credentials, error) {
	var credentials domain.Credentials
	err := lr.db.QueryRow(getCredentials, username, password).Scan(&credentials.ID, &credentials.Username, &credentials.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Credentials{}, errors.New("user and password not found")
		}
		return domain.Credentials{}, err
	}
	
	return credentials, nil
}	

func (lr *CredentialRepository) GetCredentialIDByUsername(username string) (int64, error) {
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