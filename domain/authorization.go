package domain

type Authorizations struct {
	ID int64
	CredentialID int64 `json:"credential_id"`
	AppName string `json:"app_name"`
	AuthType string `json:"auth_type"`
}