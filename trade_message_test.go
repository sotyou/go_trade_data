package data

import "testing"

func TestOrderCommand_RequestQuantity(t *testing.T) {
	type fields struct {
		Qty   float64
		Price float64
	}
	type args struct {
		price float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"market buy", fields{Qty: 1, Price: 0}, args{100}, 1},
		{"market sell", fields{Qty: -1, Price: 0}, args{100}, -0.01},
		{"limit buy", fields{Qty: 1, Price: 100}, args{100}, 1},
		{"limit sell", fields{Qty: -1, Price: 100}, args{100}, -1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderCommand{
				Qty:   tt.fields.Qty,
				Price: tt.fields.Price,
			}
			if got := o.RequestQuantity(tt.args.price); got != tt.want {
				t.Errorf("RequestQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderCommand_Cash(t *testing.T) {
	type fields struct {
		Qty   float64
		Price float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"market buy", fields{Qty: 1, Price: 0}, 1},
		{"market sell", fields{Qty: -1, Price: 0}, -1},
		{"limit buy", fields{Qty: 1, Price: 100}, 100},
		{"limit sell", fields{Qty: -1, Price: 100}, -100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderCommand{
				Qty:   tt.fields.Qty,
				Price: tt.fields.Price,
			}
			if got := o.Cash(); got != tt.want {
				t.Errorf("Cash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderCommand_Volume(t *testing.T) {
	type fields struct {
		Qty   float64
		Price float64
	}
	type args struct {
		price float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"market buy", fields{Qty: 1, Price: 0}, args{100}, 0.01},
		{"market sell", fields{Qty: -1, Price: 0}, args{100}, -0.01},
		{"limit buy", fields{Qty: 1, Price: 100}, args{100}, 1},
		{"limit sell", fields{Qty: -1, Price: 100}, args{100}, -1},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OrderCommand{
				Qty:   tt.fields.Qty,
				Price: tt.fields.Price,
			}
			if got := o.Volume(tt.args.price); got != tt.want {
				t.Errorf("Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}