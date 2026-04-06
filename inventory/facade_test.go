package inventory_test

import (
	"context"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/inventory"
)

func TestFacadeABCDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.ABC(context.Background(), abc.Input{
		Products: []abc.Product{
			{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
			{SKU: "B", Name: "Beta", Quantity: 1, Price: 5},
		},
	})
	if err != nil {
		t.Fatalf("ABC returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].SKU != "A" {
		t.Fatalf("expected highest-value SKU A first, got %s", output.Results[0].SKU)
	}
}
