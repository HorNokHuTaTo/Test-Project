package it02model

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
