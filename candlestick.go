package data

import "time"

type (
	// Candlestick 1根K线
	Candlestick struct {
		Amount float64   `json:"amount" csv:"-"`
		Open   float64   `json:"open" csv:"open"`
		Close  float64   `json:"close" csv:"close"`
		High   float64   `json:"high" csv:"high"`
		Time   time.Time `json:"ts" csv:"time"`
		Count  int64     `json:"count" csv:"-"`
		Low    float64   `json:"low" csv:"low"`
		Volume float64   `json:"vol" csv:"-"`
	}

	Candlesticks []Candlestick
)

func (c Candlesticks) First() Candlestick {
	return c[0]
}

func (c Candlesticks) Last() Candlestick {
	if c == nil || len(c) <= 0 {
		return Candlestick{}
	}
	return c[len(c)-1]
}


func (c Candlesticks) Tail(n int) Candlesticks {
	if n > len(c) {
		return c
	}
	return c[len(c)-n:]
}

func (c Candlesticks) Head(n int) Candlesticks {
	if n > len(c) {
		return c
	}
	return c[:n]
}