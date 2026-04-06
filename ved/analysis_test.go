package ved_test

import (
	"context"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/ved"
)

func TestAnalyzeClassifiesItemsByCriticality(t *testing.T) {
	output, err := ved.Analyze(context.Background(), ved.Input{
		Items: []ved.Item{
			{SKU: "V-1", Name: "Vital", CriticalityScore: 95},
			{SKU: "E-1", Name: "Essential", CriticalityScore: 55},
			{SKU: "D-1", Name: "Desirable", CriticalityScore: 15},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "V" {
		t.Fatalf("expected first item group V, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "E" {
		t.Fatalf("expected second item group E, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "D" {
		t.Fatalf("expected third item group D, got %s", output.Results[2].Group)
	}
}

func TestAnalyzeSupportsCustomThresholds(t *testing.T) {
	output, err := ved.Analyze(context.Background(), ved.Input{
		Thresholds: ved.Thresholds{
			VitalMinScore:     80,
			EssentialMinScore: 60,
		},
		Items: []ved.Item{
			{SKU: "A", Name: "A", CriticalityScore: 75},
			{SKU: "B", Name: "B", CriticalityScore: 62},
			{SKU: "C", Name: "C", CriticalityScore: 59},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].Group != "E" {
		t.Fatalf("expected first item group E, got %s", output.Results[0].Group)
	}
	if output.Results[1].Group != "E" {
		t.Fatalf("expected second item group E, got %s", output.Results[1].Group)
	}
	if output.Results[2].Group != "D" {
		t.Fatalf("expected third item group D, got %s", output.Results[2].Group)
	}
}
