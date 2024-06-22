package gorm

import (
	"bbdk/domain/entity"
	transactionRepo "bbdk/domain/repository/transaction"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *entity.Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return transactionRepo.ErrAlreadyExist
			}
		}
		return err
	}
	return nil
}

func (r *TransactionRepository) GetTransactionByID(id uint) (*entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, transactionRepo.ErrNotFound
		}
		return nil, err
	}
	return &transaction, nil
}

func (r *TransactionRepository) UpdateTransaction(transaction *entity.Transaction) error {
	tx := r.db.Where("id", transaction.ID).UpdateColumns(transaction)
	if err := tx.Error; err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return transactionRepo.ErrAlreadyExist
			}
		}
		return err
	}
	if tx.RowsAffected < 1 {
		return transactionRepo.ErrNotFound
	}
	return nil
}

func (r *TransactionRepository) DeleteTransaction(id uint) error {
	tx := r.db.Delete(&entity.Transaction{}, id)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected < 1 {
		return transactionRepo.ErrNotFound
	}
	return nil
}

func (r *TransactionRepository) GetAll(offset, limit int) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	if err := r.db.Offset(offset).Limit(limit).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) FindByField(field string, value interface{}) (*entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.Where(field+" = ?", value).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, transactionRepo.ErrNotFound
		}
		return nil, err
	}
	return &transaction, nil
}
