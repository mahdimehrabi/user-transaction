package transaction

import (
	"bbdk/domain/dto"
	"bbdk/domain/entity"
	"errors"
)

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrNotFound     = errors.New("not found")
)

type Repository interface {
	CreateTransaction(Transaction *entity.Transaction) error
	GetTransactionByID(id uint) (*entity.Transaction, error)
	UpdateTransaction(Transaction *entity.Transaction) error
	DeleteTransaction(id uint) error
	GetAll(offset, limit int) ([]*entity.Transaction, error)
	FindByField(field string, value interface{}) (*entity.Transaction, error)
	FindSumByTypeUserID(userID uint) ([]*dto.Report, error)
}
