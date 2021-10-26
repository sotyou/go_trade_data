package data

type (
	OrderBook struct {
		Ts     int64   `parquet:"name=ts, type=INT64"`
		Level  int32   `parquet:"name=level, type=INT32"`
		Price  float64 `parquet:"name=price, type=DOUBLE"`
		Amount float64 `parquet:"name=amount, type=DOUBLE"`
		Side   int32   `parquet:"name=side, type=INT32"`
	}
	OrderBooks []OrderBook
)
