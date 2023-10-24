package models

import "time"

type User struct {
	UserID            int       `json:"user_id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
