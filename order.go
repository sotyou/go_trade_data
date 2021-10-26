package data

type Order struct {
	Pair          ContractPair
	OrderId       string
	ClientOrderId string
	Amount        float64
	Price         float64
	Volume        int32
	Magnitude     int32
	Direction     OrderDirection
	Offset        OrderOffset
	Status        OrderStatus
	Ts            int64
}
