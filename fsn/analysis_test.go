package fsn_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/fsn"
)

func TestAnalyzeClassifiesItemsByMovementFrequency(t *testing.T) {
	output, err := fsn.Analyze(context.Background(), fsn.Input{
		Items: []fsn.Item{
			{SKU: "F-1", Name: "Fast", Movements: []float64{10, 10, 10, 10}},
			{SKU: "S-1", Name: "Slow", Movements: []float64{0, 5, 0, 0}},
			{SKU: "N-1", Name: "Non", Movements: []float64{0, 0, 0, 0}},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "F" {
		t.Fatalf("expected first item group F, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "S" {
		t.Fatalf("expected second item group S, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "N" {
		t.Fatalf("expected third item group N, got %s", output.Results[2].Group)
	}
	if output.Results[1].LastMovementPeriod != 2 {
		t.Fatalf("expected second item last movement period 2, got %d", output.Results[1].LastMovementPeriod)
	}
	if output.Summary.TotalItems != 3 || output.Summary.FCount != 1 || output.Summary.SCount != 1 || output.Summary.NCount != 1 {
		t.Fatalf("unexpected summary: %+v", output.Summary)
	}
}

func TestAnalyzeSupportsCustomActivityThresholds(t *testing.T) {
	output, err := fsn.Analyze(context.Background(), fsn.Input{
		Thresholds: fsn.Thresholds{
			FastMinActivityRate: 50,
			SlowMinActivityRate: 25,
		},
		Items: []fsn.Item{
			{SKU: "A", Name: "A", Movements: []float64{1, 0, 1, 0}},
			{SKU: "B", Name: "B", Movements: []float64{1, 0, 0, 0}},
			{SKU: "C", Name: "C", Movements: []float64{0, 0, 0, 0}},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].Group != "F" {
		t.Fatalf("expected first item group F, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "S" {
		t.Fatalf("expected second item group S, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "N" {
		t.Fatalf("expected third item group N, got %s", output.Results[2].Group)
	}
}

func TestAnalyzeReturnsErrorOnInvalidThresholds(t *testing.T) {
	_, err := fsn.Analyze(context.Background(), fsn.Input{
		Thresholds: fsn.Thresholds{
			FastMinActivityRate: 20,
			SlowMinActivityRate: 25,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}
