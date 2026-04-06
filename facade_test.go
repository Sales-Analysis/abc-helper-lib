package analytics

import "testing"

func TestNewProvidesDomainFacades(t *testing.T) {
	facade := New()
	if facade == nil {
		t.Fatal("expected facade, got nil")
	}
	if facade.Inventory() == nil {
		t.Fatal("expected inventory facade, got nil")
	}
	if facade.Customer() == nil {
		t.Fatal("expected customer facade, got nil")
	}
}
