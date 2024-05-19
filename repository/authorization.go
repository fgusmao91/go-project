package repository

import (
	"database/sql"
	"v1/domain"
)

const (
	insertAuthorizations = "INSERT INTO authorizations (credential_id, appname, authtype) VALUES (?, ?, ?)"
	getAuthorizations = "SELECT id, credential_id, appname, authtype FROM authorizations WHERE credential_id = ?"
)

type AuthorizationRepository struct {
	db *sql.DB
}

func NewAuthorizationRepository(db *sql.DB) *AuthorizationRepository {
	return &AuthorizationRepository{
		db: db,
	}
}

func (lr *AuthorizationRepository) InsertAuthorizations(authorizations domain.Authorizations) (*int64,error) {
	result, err := lr.db.Exec(insertAuthorizations, authorizations.CredentialID, authorizations.AppName, authorizations.AuthType)
	if err != nil {
		return nil, err
	}

	lasteInsertedID, _ := result.LastInsertId()

	return &lasteInsertedID, nil
}

func (lr *AuthorizationRepository) GetAuthorizations(credentialID int64) ([]domain.Authorizations, error) {
	rows, err := lr.db.Query(getAuthorizations, credentialID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authorizations []domain.Authorizations
	for rows.Next() {
		var authorization domain.Authorizations
		err := rows.Scan(&authorization.ID,&authorization.CredentialID, &authorization.AppName, &authorization.AuthType)
		if err != nil {
			return nil, err
		}
		authorizations = append(authorizations, authorization)
	}
	
	return authorizations, nil
}