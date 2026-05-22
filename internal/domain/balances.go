package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Balance struct {
	ID       uuid.UUID       `json:"id"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
}
