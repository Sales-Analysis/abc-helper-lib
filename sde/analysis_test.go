package sde_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/sde"
)

func TestAnalyzeDerivesThresholdsFromData(t *testing.T) {
	output, err := sde.Analyze(context.Background(), sde.Input{
		Items: []sde.Item{
			{SKU: "S-1", Name: "Scarce", LeadTimeDays: 120},
			{SKU: "D-1", Name: "Difficult", LeadTimeDays: 45},
			{SKU: "E-1", Name: "Easy", LeadTimeDays: 7},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "S" {
		t.Fatalf("expected first item group S, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "D" {
		t.Fatalf("expected second item group D, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "E" {
		t.Fatalf("expected third item group E, got %s", output.Results[2].Group)
	}
	if output.Summary.TotalItems != 3 || output.Summary.SCount != 1 || output.Summary.DCount != 1 || output.Summary.ECount != 1 {
		t.Fatalf("unexpected summary: %+v", output.Summary)
	}
}

func TestAnalyzeSupportsCustomThresholds(t *testing.T) {
	output, err := sde.Analyze(context.Background(), sde.Input{
		Thresholds: sde.Thresholds{
			ScarceMinLeadTime:    90,
			DifficultMinLeadTime: 30,
		},
		Items: []sde.Item{
			{SKU: "A", Name: "A", LeadTimeDays: 120},
			{SKU: "B", Name: "B", LeadTimeDays: 45},
			{SKU: "C", Name: "C", LeadTimeDays: 10},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].Group != "S" {
		t.Fatalf("expected first item group S, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "D" {
		t.Fatalf("expected second item group D, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "E" {
		t.Fatalf("expected third item group E, got %s", output.Results[2].Group)
	}
}

func TestAnalyzeReturnsErrorOnIncompleteThresholdOverride(t *testing.T) {
	_, err := sde.Analyze(context.Background(), sde.Input{
		Thresholds: sde.Thresholds{
			ScarceMinLeadTime: 90,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}
