package models

import "time"

type User struct {
	UserID            int       `db:"user_id" json:"user_id"`
	Username          string    `db:"username" json:"username"`
	HashedPassword    string    `db:"hashed_password" json:"hashed_password"`
	Email             string    `db:"email" json:"email"`
	PasswordChangedAt time.Time `db:"password_changed_at" json:"password_changed_at"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}
