package data

type Position struct {
	OrderID string
	Price float64
	Amount float64
	Volume int32
	LeveRate int8
	Direction OrderDirection
}

type Positions []Position
