package abc_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
)

func TestAnalyzeMatchesLegacyCalculate(t *testing.T) {
	products := []abc.Product{
		{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
		{SKU: "B", Name: "Beta", Quantity: 2, Price: 10},
		{SKU: "C", Name: "Gamma", Quantity: 1, Price: 5},
	}

	legacy := abc.New()
	legacy.Calculate(products)

	output, err := abc.Analyze(context.Background(), abc.Input{Products: products})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if !reflect.DeepEqual(legacy.Result, output.Results) {
		t.Fatalf("Analyze results differ from legacy API.\nlegacy=%+v\nnew=%+v", legacy.Result, output.Results)
	}

	if output.TotalRevenue != 125 {
		t.Fatalf("expected total revenue 125, got %.2f", output.TotalRevenue)
	}
}

func TestAnalyzeDetailedKeepsOriginalIndex(t *testing.T) {
	products := []abc.Product{
		{SKU: "A", Name: "Alpha", Quantity: 1, Price: 10},
		{SKU: "B", Name: "Beta", Quantity: 10, Price: 10},
	}

	output, err := abc.AnalyzeDetailed(context.Background(), abc.Input{Products: products})
	if err != nil {
		t.Fatalf("AnalyzeDetailed returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].OriginalIndex != 1 {
		t.Fatalf("expected first result to point to original index 1, got %d", output.Results[0].OriginalIndex)
	}
	if output.Results[1].OriginalIndex != 0 {
		t.Fatalf("expected second result to point to original index 0, got %d", output.Results[1].OriginalIndex)
	}
}
