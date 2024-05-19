package domain

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorizations struct {
	CredentialID int64 `json:"credential_id"`
	AppName string `json:"app_name"`
	AuthType string `json:"auth_type"`
}