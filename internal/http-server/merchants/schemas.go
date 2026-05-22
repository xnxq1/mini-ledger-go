package merchants

import "github.com/shopspring/decimal"

type CreateMerchantRequest struct {
	Name       string          `json:"name"`
	PercentFee decimal.Decimal `json:"percent_fee"`
}
