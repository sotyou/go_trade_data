package data

import (
	"encoding/json"
	"github.com/FireCoinJp/quant-core/core/atom"
)

const (
	CommandOrder        = "order"
	CommandSleep        = "sleep"
	CommandAwake        = "awake"
	CommandCancel       = "cancel"
	CommandExit         = "exit"
	CommandCancelOnSide = "cancel_on_side"
	CommandSetting      = "setting"
)

// TradeMessage {"cmd":"order|cancel|entry|exit", }
type TradeMessage struct {
	Cmd     string          `json:"cmd"` // order, cancel, entry, exit
	Args    json.RawMessage `json:"args"`
	Comment string          `json:"comment"` // 可以为空， 参考内容
}

func (m TradeMessage) GetCommand(v interface{}) error {
	return atom.DecodeFromByte(m.Args, v)
}

type OrderCommand struct {
	Qty   float64 `json:"qty"`   // (数量)正数为做多， 负数为做空
	Price float64 `json:"price"` // 限价单的价格，如果没有的为市价单
	//Ops OrderOpts `json:"opts"`  // 特殊订单需要的参数
}

func (o OrderCommand) Byte() []byte {
	b, _ := atom.EncodeToBytes(o)
	return b
}

func (o OrderCommand) TradeSide() int64 {
	if o.Qty > 0 {
		return 1
	}
	return -1
}

func (o OrderCommand) PositionSide(holdVolume float64) int64 {
	if holdVolume*o.Qty >= 0 {
		return 1
	}
	return -1
}

func (o OrderCommand) IsMarket() bool {
	return o.Price == 0
}

func (o OrderCommand) UseCash() bool {
	return o.IsMarket() && o.Qty > 0
}

func (o OrderCommand) Volume(price float64) float64 {
	if o.IsMarket() {
		return o.Qty / price
	}
	return o.Qty
}

func (o OrderCommand) RequestQuantity(price float64) float64 {
	if o.IsMarket() && o.Qty < 0 {
		return o.Qty / price
	}
	return o.Qty
}

func (o OrderCommand) Cash() float64 {
	if o.IsMarket() {
		return o.Qty
	}
	return o.Qty * o.Price
}

type CancelOnSideCommand struct {
	Side string `json:"cancel_side"` // 平单方向, [all, buy, sell]
}

// MatchType TODO 策略单
type MatchType string

const (
	IOC   MatchType = "ioc"   // 立刻交易, 其它的退出
	FOK   MatchType = "fok"   // 全部交易或者退出
	Maker MatchType = "maker" // 只做maker
)

type OrderOpts struct {
	TimeMethod MatchType // ioc, fok, maker
	Profit     float64   // 止盈利润点
	Loss       float64   // 止损利润点
	Stop       float64   // 止损价格
	Limit      float64   // 止盈价格
}
