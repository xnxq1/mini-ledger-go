package domain

import "github.com/shopspring/decimal"

type Merchant struct {
	Name       string          `json:"name"`
	PercentFee decimal.Decimal `json:"percent_fee"`
}
