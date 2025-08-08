package abc_test

import (
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
)

func TestABC(t *testing.T) {

	products := []abc.Product{
		{SKU: "001", Name: "Product A", Quantity: 10, Price: 10}, // 100 total
		{SKU: "002", Name: "Product B", Quantity: 5, Price: 20},  // 100 total
		{SKU: "003", Name: "Product C", Quantity: 1, Price: 5},   // 5 total
	}

	a := abc.New()
	a.Calculate(products)

	// Then: Check number of results
	if len(a.Result) != len(products) {
		t.Fatalf("Expected %d results, got %d", len(products), len(a.Result))
	}

	// Check that order is by PriceTotal descending
	if a.Result[0].PriceTotal < a.Result[1].PriceTotal {
		t.Errorf("Results not sorted by PriceTotal descending")
	}

	// Check that total calculation is correct
	expectedTotals := []float64{100, 100, 5}
	for _, res := range a.Result {
		found := false
		for _, total := range expectedTotals {
			if res.PriceTotal == total {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected PriceTotal in results: %v", res.PriceTotal)
		}
	}

	// Check that groups are assigned (should have A, B, C)
	groups := map[string]bool{}
	for _, res := range a.Result {
		groups[res.Group] = true
	}

	if len(groups) < 2 {
		t.Errorf("Expected at least 2 different groups, got: %v", groups)
	}

}
