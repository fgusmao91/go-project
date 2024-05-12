package repository

import (
	"database/sql"
	"errors"
	"v1/domain"
)

const (
	insertCredentials = "INSERT INTO credentials (username, password) VALUES (?, ?)"
	getCredentials    = "SELECT username, password FROM credentials WHERE username = ? and password = ?"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (lr *AuthRepository) InsertCredentials(credentials domain.Credentials) error {
	_, err := lr.db.Exec(insertCredentials, credentials.Username, credentials.Password)
	if err != nil {
		return err
	}

	return nil
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