package data

import "testing"

func TestOrderStatus_IsTraded(t *testing.T) {
	tests := []struct {
		name string
		o    OrderStatus
		want bool
	}{
		{"unknown", OrderStatusUnknown, false},
		{"Failure", OrderStatusFailure, false},
		{"Ready", OrderStatusReady, false},
		{"Cancelled", OrderStatusCancelled, false},
		{"Submitted", OrderStatusSubmitted, false},
		{"PartlyCanceled", OrderStatusPartlyCanceled, false},
		{"Executed", OrderStatusExecuted, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsTraded(); got != tt.want {
				t.Errorf("IsTraded() = %v, want %v", got, tt.want)
			}
		})
	}
}
