package warframe_market_prime_trash_buyer

import (
	"testing"
)

func TestGetProfitableOrders(t *testing.T) {
	orders, err := GetProfitableOrders()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if orders == nil {
		t.Fatal("expected non-nil slice, got nil")
	}
}
