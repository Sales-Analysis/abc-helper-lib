package reorderpoint_test

import (
	"context"
	"math"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/reorderpoint"
)

func TestAnalyzeCalculatesReorderPoint(t *testing.T) {
	output, err := reorderpoint.Analyze(context.Background(), reorderpoint.Input{
		Items: []reorderpoint.Item{
			{SKU: "A", Name: "Alpha", AverageDemandPerPeriod: 20, LeadTimePeriods: 4, SafetyStock: 33},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}

	if math.Abs(output.Results[0].LeadTimeDemand-80) > 0.001 {
		t.Fatalf("expected lead time demand 80, got %.4f", output.Results[0].LeadTimeDemand)
	}
	if math.Abs(output.Results[0].ReorderPoint-113) > 0.001 {
		t.Fatalf("expected reorder point 113, got %.4f", output.Results[0].ReorderPoint)
	}
}

func TestAnalyzeReturnsZeroLeadTimeDemandForInvalidInputs(t *testing.T) {
	output, err := reorderpoint.Analyze(context.Background(), reorderpoint.Input{
		Items: []reorderpoint.Item{
			{SKU: "A", Name: "Alpha", AverageDemandPerPeriod: 0, LeadTimePeriods: 4, SafetyStock: 5},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].LeadTimeDemand != 0 {
		t.Fatalf("expected lead time demand 0, got %.4f", output.Results[0].LeadTimeDemand)
	}
	if output.Results[0].ReorderPoint != 5 {
		t.Fatalf("expected reorder point 5, got %.4f", output.Results[0].ReorderPoint)
	}
}
