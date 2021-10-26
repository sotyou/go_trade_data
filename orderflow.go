package data

type (
	OrderFlow struct {
		Ts     int64   `parquet:"name=ts, type=INT64"`
		Price  float64 `parquet:"name=price, type=DOUBLE"`
		Amount float64 `parquet:"name=amount, type=DOUBLE"`
		Side   int32   `parquet:"name=side, type=INT32"`
	}
	OrderFlows []OrderFlow
)
