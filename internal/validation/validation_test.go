package validation

import (
	"math"
	"strings"
	"testing"
)

func TestRequireNonNegativeFloatRejectsNaN(t *testing.T) {
	err := RequireNonNegativeFloat("field", math.NaN())
	if err == nil {
		t.Fatal("expected error for NaN, got nil")
	}
	if !strings.Contains(err.Error(), "must be finite") {
		t.Fatalf("expected finite error, got %v", err)
	}
}

func TestRequirePositiveFloatRejectsInfinity(t *testing.T) {
	err := RequirePositiveFloat("field", math.Inf(1))
	if err == nil {
		t.Fatal("expected error for infinity, got nil")
	}
	if !strings.Contains(err.Error(), "must be finite") {
		t.Fatalf("expected finite error, got %v", err)
	}
}
