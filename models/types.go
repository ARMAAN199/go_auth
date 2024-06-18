package models

import "time"

type Cars struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  uint   `json:"year"`
}

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	IsAdmin      bool      `json:"is_admin"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
