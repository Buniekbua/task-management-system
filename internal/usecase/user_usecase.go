package usecase

import (
	"github.com/buniekbua/task-managment-system/internal/db"
	"github.com/buniekbua/task-managment-system/internal/models"
)

type UserUseCase interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
}

type UserUseCaseImpl struct {
	userRepository db.UserRepository
}

func NewUserUseCase(ur db.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{userRepository: ur}
}

func (uc *UserUseCaseImpl) CreateUser(user *models.User) error {
	return uc.userRepository.CreateUser(user)
}

func (uc *UserUseCaseImpl) GetUserByID(id int) (*models.User, error) {
	return uc.userRepository.GetUserByID(id)
}

func (uc *UserUseCaseImpl) UpdateUser(id int, user *models.User) error {
	return uc.userRepository.UpdateUser(id, user)
}

func (uc *UserUseCaseImpl) DeleteUser(id int) error {
	return uc.userRepository.DeleteUser(id)
}
