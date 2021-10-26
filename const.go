package data

type (
	OrderStatus int
)

const (
	OrderStatusUnknown        OrderStatus = iota // 未知状态
	OrderStatusFailure                           // 挂单失败
	OrderStatusReady                             // 准备中
	OrderStatusSubmitted                         // 已挂单
	OrderStatusPartlyExecuted                    // 部分成交, 其它挂单中
	OrderStatusPartlyCanceled                    // 部分成交并取消
	OrderStatusExecuted                          // 已成交
	OrderStatusCancelled                         // 已取消
)

func (o OrderStatus) IsTraded() bool {
	return o == OrderStatusExecuted
}

func (o OrderStatus) IsOrdered() bool {
	return o >= OrderStatusCancelled
}

func (o OrderStatus) String() string {
	switch o {
	case OrderStatusFailure:
		return "failure"
	case OrderStatusReady:
		return "ready"
	case OrderStatusCancelled:
		return "cancelled"
	case OrderStatusSubmitted:
		return "submitted"
	case OrderStatusPartlyExecuted:
		return "partly_executed"
	case OrderStatusPartlyCanceled:
		return "partly_canceled"
	case OrderStatusExecuted:
		return "executed"
	}
	return "unknown"
}

type ContractType int

const (
	FlatBase ContractType = iota // 法币本位
	CoinBase                     // 币本位
	MixBase                      // 现货杠杆
)

type OrderDirection int32

const (
	OrderDirectionUnknown OrderDirection = 0
	OrderDirectionBuy     OrderDirection = 1
	OrderDirectionSell    OrderDirection = -1
)

func (o OrderDirection) String() string {
	switch o {
	case OrderDirectionBuy:
		return "buy"
	case OrderDirectionSell:
		return "sell"
	}
	return "unknown"
}

func (o OrderDirection) Float() float64 {
	return float64(o)
}

type PositionSide int

const (
	PositionSideUnknown PositionSide = 0
	PositionSideOpen  PositionSide = 1
	PositionSideClose PositionSide = -1
)

func (p PositionSide) String() string {
	switch p {
	case PositionSideOpen:
		return "open"
	case PositionSideClose:
		return "close"
	}
	return "unknown"
}

func (p PositionSide) Float() float64 {
	return float64(p)
}

type OrderEvent int

const (
	OrderEventUnknown     = iota // 不支持事件
	OrderEventPlace              // 下单
	OrderEventCancel             // 撤单
	OrderEventClose              // 平仓
	OrderEventLiquidation        // 爆仓
	OrderEventDelivery           // 交割
)

// 支持这个接口, 可以用 atom.S(o) -> string
func (o OrderEvent) String() string {
	switch o {
	case OrderEventPlace:
		return "place"
	case OrderEventCancel:
		return "cancel"
	case OrderEventClose:
		return "close"
	case OrderEventLiquidation:
		return "liquidation"
	case OrderEventDelivery:
		return "delivery"
	}
	return "unknown"
}

type (
	TriggerOrderStatus int
)

const (
	TriggerOrderStatusUnknown   TriggerOrderStatus = iota // 未知状态
	TriggerOrderStatusSuccess                             // 触发成功
	TriggerOrderStatusFailure                             // 触发失败
	TriggerOrderStatusCancelled                           // 已取消
	TriggerOrderStatusSubmitted                           // 已挂单
	TriggerOrderStatusError                               // 错读的报单
)

func (tos TriggerOrderStatus) String() string {
	switch tos {
	case TriggerOrderStatusSubmitted:
		return "submitted"
	case TriggerOrderStatusCancelled:
		return "canceled"
	case TriggerOrderStatusSuccess:
		return "success"
	case TriggerOrderStatusFailure:
		return "failure"
	case TriggerOrderStatusError:
		return "error"
	}
	return "unknown"
}

type OrderPriceType int8

const (
	OrderPriceTypeMaker = iota
	OrderPriceTypeTaker
	OrderPriceTypeLimit
)

func (o OrderPriceType) String() string {
	if o == OrderPriceTypeTaker {
		return "taker"
	} else if o == OrderPriceTypeMaker {
		return "maker"
	} else if o == OrderPriceTypeLimit {
		return "limit"
	}
	return "unknown"
}

type OrderOffset int8

const (
	OrderOffsetUnknown = 0
	OrderOffsetOpen    = 1
	OrderOffsetClose   = -1
)

func (o OrderOffset) String() string {
	if o == OrderOffsetOpen {
		return "open"
	} else if o == OrderEventClose {
		return "close"
	}
	return "unknown"
}
