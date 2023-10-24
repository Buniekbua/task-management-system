package db

import (
	"database/sql"
	"time"

	"github.com/buniekbua/task-managment-system/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (username, hashed_password, email, password_changed_at, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING user_id
	`
	err := r.db.QueryRow(query, user.Username, user.HashedPassword, user.Email, time.Now(), time.Now()).Scan(&user.UserID)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to create user",
		}
	}

	return nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id=$1"
	err := r.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, &DBError{
			Err: err,
			Ctx: "Failed to get user by ID",
		}
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(id int, user *models.User) error {
	query := `UPDATE users SET username = $1, hashed_password = $2, email = $3, password_changed_at = $4 WHERE user_id = $5`

	_, err := r.db.Exec(query, user.Username, user.HashedPassword, user.Email, time.Now(), id)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to update user",
		}
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return &DBError{
			Err: err,
			Ctx: "Failed to delete user",
		}
	}

	return nil
}
