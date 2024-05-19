package domain

type Credentials struct {
	ID int64
	Username string `json:"username"`
	Password string `json:"password"`
}