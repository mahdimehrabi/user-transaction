package service

import (
	"bbdk/domain/dto"
	"bbdk/domain/entity"
	transactionRepo "bbdk/domain/repository/transaction"
	logger "bbdk/infrastructure/log"
	"errors"
)

type TransactionService interface {
	CreateTransaction(transaction *entity.Transaction) error
	GetTransactionByID(id uint) (*entity.Transaction, error)
	UpdateTransaction(transaction *entity.Transaction) error
	DeleteTransaction(id uint) error
	GetAllTransactions(page, pageSize int) ([]*entity.Transaction, error)
	FindTransactionByField(field string, value interface{}) (*entity.Transaction, error)
	GetTransactionReportByUserID(userID uint) ([]*dto.Report, error)
}

type transactionService struct {
	transactionRepo transactionRepo.Repository
	logger          logger.Logger
}

// NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo transactionRepo.Repository, logger logger.Logger) *transactionService {
	return &transactionService{transactionRepo: transactionRepo, logger: logger}
}

func (s *transactionService) CreateTransaction(transaction *entity.Transaction) error {
	if err := s.transactionRepo.CreateTransaction(transaction); err != nil {
		if errors.Is(err, transactionRepo.ErrAlreadyExist) {
			return transactionRepo.ErrAlreadyExist
		}
		s.logger.Errorf("failed to create transaction: %s", err.Error())
		return err
	}
	return nil
}

func (s *transactionService) GetTransactionByID(id uint) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.GetTransactionByID(id)
	if err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			return nil, err
		}
		s.logger.Errorf("failed to get transaction: %s", err.Error())
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) UpdateTransaction(transaction *entity.Transaction) error {
	if err := s.transactionRepo.UpdateTransaction(transaction); err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			return err
		} else if errors.Is(err, transactionRepo.ErrAlreadyExist) {
			return transactionRepo.ErrAlreadyExist
		}
		s.logger.Errorf("failed to update transaction: %s", err.Error())
		return err
	}
	return nil
}

func (s *transactionService) DeleteTransaction(id uint) error {
	if err := s.transactionRepo.DeleteTransaction(id); err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			return err
		}
		s.logger.Errorf("failed to delete transaction: %s", err.Error())
		return err
	}
	return nil
}

func (s *transactionService) GetAllTransactions(page, pageSize int) ([]*entity.Transaction, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, ErrInvalidPage
	}
	offset := (page - 1) * pageSize
	limit := pageSize

	transactions, err := s.transactionRepo.GetAll(offset, limit)
	if err != nil {
		s.logger.Errorf("failed to get all transactions: %s", err.Error())
		return nil, err
	}
	return transactions, nil
}

func (s *transactionService) FindTransactionByField(field string, value interface{}) (*entity.Transaction, error) {
	transaction, err := s.transactionRepo.FindByField(field, value)
	if err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			return nil, err
		}
		s.logger.Errorf("failed to find transaction by field: %s", err.Error())
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) GetTransactionReportByUserID(userID uint) ([]*dto.Report, error) {
	report, err := s.transactionRepo.FindSumByTypeUserID(userID)
	if err != nil {
		s.logger.Errorf("failed to fetch transaction report for user %d: %s", userID, err.Error())
		return nil, err
	}
	return report, nil
}
