package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	FromMerchantId uuid.UUID       `json:"from_merchant_id"`
	ToMerchantId   uuid.UUID       `json:"to_merchant_id"`
	Amount         decimal.Decimal `json:"amount"`
	Currency       string          `json:"currency"`
	PercentFee     decimal.Decimal `json:"percent_fee"`
	IdempotencyKey string          `json:"idempotency_key"`
}
