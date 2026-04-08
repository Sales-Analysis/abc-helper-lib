package hml_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/hml"
)

func TestAnalyzeDerivesThresholdsFromData(t *testing.T) {
	output, err := hml.Analyze(context.Background(), hml.Input{
		Items: []hml.Item{
			{SKU: "H-1", Name: "High", UnitCost: 500},
			{SKU: "M-1", Name: "Medium", UnitCost: 150},
			{SKU: "L-1", Name: "Low", UnitCost: 30},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "H" {
		t.Fatalf("expected first item group H, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "M" {
		t.Fatalf("expected second item group M, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "L" {
		t.Fatalf("expected third item group L, got %s", output.Results[2].Group)
	}
	if output.Summary.TotalItems != 3 || output.Summary.HCount != 1 || output.Summary.MCount != 1 || output.Summary.LCount != 1 {
		t.Fatalf("unexpected summary: %+v", output.Summary)
	}
}

func TestAnalyzeSupportsCustomThresholds(t *testing.T) {
	output, err := hml.Analyze(context.Background(), hml.Input{
		Thresholds: hml.Thresholds{
			HighMinUnitCost:   200,
			MediumMinUnitCost: 100,
		},
		Items: []hml.Item{
			{SKU: "A", Name: "A", UnitCost: 300},
			{SKU: "B", Name: "B", UnitCost: 150},
			{SKU: "C", Name: "C", UnitCost: 50},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].Group != "H" {
		t.Fatalf("expected first item group H, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "M" {
		t.Fatalf("expected second item group M, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "L" {
		t.Fatalf("expected third item group L, got %s", output.Results[2].Group)
	}
}

func TestAnalyzeReturnsErrorOnIncompleteThresholdOverride(t *testing.T) {
	_, err := hml.Analyze(context.Background(), hml.Input{
		Thresholds: hml.Thresholds{
			HighMinUnitCost: 200,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}
