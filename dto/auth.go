package dto

type LoginRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppName string `json:"app_name"`
	AuthType string `json:"auth_type"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}