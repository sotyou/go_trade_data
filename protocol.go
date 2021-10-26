package data

type OrderResProtocol interface {
    ToOrder() Order
    GetOrderID() string
}

type TpSlOrderResProtocol interface {
    GetTpOrderID() string
    GetSlOrderID() string
}

type CancelTpSlResProtocol interface {
    GetSuccessOrderID() string
}

type WssOrderResProtocol interface {
    ToOrder() Order
    ToTrade() Trade
}

type PositionResProtocol interface {
    ToPositions() Positions
}
