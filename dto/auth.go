package dto

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorization struct {
	AppName string `json:"app_name"`
	AuthType string `json:"auth_type"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}