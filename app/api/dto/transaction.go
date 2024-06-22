package dto

import "bbdk/domain/entity"

type TransactionRequest struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"user_id" binding:"fkGorm=users"`
	Amount float64 `json:"amount"binding:"required"`
	Type   string  `json:"type" binding:"required"` // e.g., game_referral, p2e, seazen_zero
}

func (req *TransactionRequest) ToEntity() *entity.Transaction {
	return &entity.Transaction{
		ID:     req.ID,
		UserID: req.UserID,
		Amount: req.Amount,
		Type:   req.Type,
	}
}

type TransactionResponse struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"` // e.g., game_referral, p2e, seazen_zero
}

func (resp *TransactionResponse) FromEntity(transaction *entity.Transaction) {
	resp.ID = transaction.ID
	resp.UserID = transaction.UserID
	resp.Amount = transaction.Amount
	resp.Type = transaction.Type
}
