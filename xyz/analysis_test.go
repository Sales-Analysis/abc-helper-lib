package xyz_test

import (
	"context"
	"math"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

func TestAnalyzeClassifiesItemsByDemandVariability(t *testing.T) {
	output, err := xyz.Analyze(context.Background(), xyz.Input{
		Items: []xyz.Item{
			{SKU: "X", Name: "Stable", Demands: []float64{100, 100, 100}},
			{SKU: "Y", Name: "Moderate", Demands: []float64{100, 120, 80}},
			{SKU: "Z", Name: "Erratic", Demands: []float64{0, 200, 0}},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}

	if output.Results[0].Group != "X" {
		t.Fatalf("expected first item to be X, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "Y" {
		t.Fatalf("expected second item to be Y, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "Z" {
		t.Fatalf("expected third item to be Z, got %s", output.Results[2].Group)
	}
	if output.Summary.TotalItems != 3 || output.Summary.XCount != 1 || output.Summary.YCount != 1 || output.Summary.ZCount != 1 {
		t.Fatalf("unexpected summary: %+v", output.Summary)
	}

	if math.Abs(output.Results[1].CoefficientOfVariation-16.3299) > 0.01 {
		t.Fatalf("expected CV around 16.33, got %.4f", output.Results[1].CoefficientOfVariation)
	}
}

func TestAnalyzeReturnsErrorOnInvalidThresholds(t *testing.T) {
	_, err := xyz.Analyze(context.Background(), xyz.Input{
		Thresholds: xyz.Thresholds{
			XMaxCV: 30,
			YMaxCV: 20,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}

func TestAnalyzeReturnsErrorOnInvalidDemandSeriesValue(t *testing.T) {
	_, err := xyz.Analyze(context.Background(), xyz.Input{
		Items: []xyz.Item{
			{SKU: "A", Name: "A", Demands: []float64{100, math.NaN()}},
		},
	})
	if err == nil {
		t.Fatal("expected invalid demand series error, got nil")
	}
	if !strings.Contains(err.Error(), "items[].Demands[]") {
		t.Fatalf("expected demand series error, got %v", err)
	}
}
