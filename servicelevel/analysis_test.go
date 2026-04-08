package servicelevel_test

import (
	"context"
	"math"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/servicelevel"
)

func TestAnalyzeCalculatesFillRateAndCycleServiceLevel(t *testing.T) {
	output, err := servicelevel.Analyze(context.Background(), servicelevel.Input{
		Items: []servicelevel.Item{
			{SKU: "A", Name: "Alpha", DemandedUnits: 100, FulfilledUnits: 95, TotalCycles: 20, StockoutCycles: 2},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	result := output.Results[0]
	if math.Abs(result.FillRate-95) > 0.001 {
		t.Fatalf("expected fill rate 95, got %.4f", result.FillRate)
	}
	if math.Abs(result.CycleServiceLevel-90) > 0.001 {
		t.Fatalf("expected cycle service level 90, got %.4f", result.CycleServiceLevel)
	}
	if math.Abs(result.ServiceLevel-90) > 0.001 {
		t.Fatalf("expected selected service level 90, got %.4f", result.ServiceLevel)
	}
}

func TestAnalyzeFallsBackToFillRateWhenCyclesMissing(t *testing.T) {
	output, err := servicelevel.Analyze(context.Background(), servicelevel.Input{
		Items: []servicelevel.Item{
			{SKU: "A", Name: "Alpha", DemandedUnits: 50, FulfilledUnits: 45},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if math.Abs(output.Results[0].ServiceLevel-90) > 0.001 {
		t.Fatalf("expected service level 90, got %.4f", output.Results[0].ServiceLevel)
	}
}

func TestAnalyzeReturnsErrorOnImpossibleCycleInputs(t *testing.T) {
	_, err := servicelevel.Analyze(context.Background(), servicelevel.Input{
		Items: []servicelevel.Item{
			{SKU: "A", Name: "Alpha", DemandedUnits: 100, FulfilledUnits: 95, TotalCycles: 2, StockoutCycles: 3},
		},
	})
	if err == nil {
		t.Fatal("expected invalid input error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}
