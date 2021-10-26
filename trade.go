package data

import "time"

type Trade struct {
	Pair                ContractPair
	TradeId             string
	OrderId             string
	ClientOrderId       string
	TradeTime           int64
	TradeDirection      OrderDirection
	TradeOffset         OrderOffset
	Status              OrderStatus
	TradePrice          float64
	TradeBaseQuantity   float64
	TradeTargetQuantity float64
	TradeVolume         float64
	Fee                 float64
	FeeCurrency         string
	AvgPrice            float64
	Profit              float64
	Rate                float64
	UpdatedAt           time.Time
	CreatedAt           time.Time
}
