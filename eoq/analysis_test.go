package eoq_test

import (
	"context"
	"math"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/eoq"
)

func TestAnalyzeCalculatesEOQ(t *testing.T) {
	output, err := eoq.Analyze(context.Background(), eoq.Input{
		Items: []eoq.Item{
			{SKU: "A", Name: "Alpha", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}

	if math.Abs(output.Results[0].OptimalQuantity-244.9489) > 0.01 {
		t.Fatalf("expected EOQ around 244.95, got %.4f", output.Results[0].OptimalQuantity)
	}
	if math.Abs(output.Results[0].OrdersPerYear-4.8990) > 0.01 {
		t.Fatalf("expected orders per year around 4.90, got %.4f", output.Results[0].OrdersPerYear)
	}
}

func TestAnalyzeReturnsZeroForInvalidEconomicInputs(t *testing.T) {
	output, err := eoq.Analyze(context.Background(), eoq.Input{
		Items: []eoq.Item{
			{SKU: "A", Name: "Alpha", AnnualDemand: 0, OrderingCost: 50, HoldingCost: 2},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].OptimalQuantity != 0 {
		t.Fatalf("expected EOQ 0, got %.2f", output.Results[0].OptimalQuantity)
	}
}
