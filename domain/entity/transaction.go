package entity

type Transaction struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"` // e.g., game_referral, p2e, seazen_zero
}
