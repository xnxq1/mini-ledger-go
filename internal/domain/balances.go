package domain

import (
	"github.com/shopspring/decimal"
)

type Balance struct {
	BaseEntity
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
}
