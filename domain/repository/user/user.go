package user

import (
	"bbdk/domain/entity"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

// Repository defines the methods that any repository implementation must have
type Repository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
