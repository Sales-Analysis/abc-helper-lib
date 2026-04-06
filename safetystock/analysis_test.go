package safetystock_test

import (
	"context"
	"math"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/safetystock"
)

func TestAnalyzeCalculatesSafetyStock(t *testing.T) {
	output, err := safetystock.Analyze(context.Background(), safetystock.Input{
		Items: []safetystock.Item{
			{SKU: "A", Name: "Alpha", DemandStdDev: 10, LeadTimePeriods: 4, ServiceFactor: 1.65},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}

	if math.Abs(output.Results[0].SafetyStock-33) > 0.001 {
		t.Fatalf("expected safety stock 33, got %.4f", output.Results[0].SafetyStock)
	}
}

func TestAnalyzeUsesDefaultServiceFactor(t *testing.T) {
	output, err := safetystock.Analyze(context.Background(), safetystock.Input{
		Items: []safetystock.Item{
			{SKU: "A", Name: "Alpha", DemandStdDev: 10, LeadTimePeriods: 4},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if math.Abs(output.Results[0].ServiceFactor-1.65) > 0.001 {
		t.Fatalf("expected default service factor 1.65, got %.4f", output.Results[0].ServiceFactor)
	}
}
