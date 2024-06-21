package user

import (
	"bbdk/domain/entity"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

type Repository interface {
	CreateUser(user *entity.User) error
	GetUserByID(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) error
	GetAll(offset, limit int) ([]*entity.User, error)
}
