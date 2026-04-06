package gmcontribution_test

import (
	"context"
	"math"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/gmcontribution"
)

func TestAnalyzeCalculatesMargins(t *testing.T) {
	output, err := gmcontribution.Analyze(context.Background(), gmcontribution.Input{
		Items: []gmcontribution.Item{
			{SKU: "A", Name: "Alpha", Revenue: 100, COGS: 60, VariableCost: 70, FixedCost: 10},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}

	result := output.Results[0]
	if math.Abs(result.GrossMargin-40) > 0.001 {
		t.Fatalf("expected gross margin 40, got %.4f", result.GrossMargin)
	}
	if math.Abs(result.GrossMarginRate-40) > 0.001 {
		t.Fatalf("expected gross margin rate 40, got %.4f", result.GrossMarginRate)
	}
	if math.Abs(result.ContributionMargin-30) > 0.001 {
		t.Fatalf("expected contribution margin 30, got %.4f", result.ContributionMargin)
	}
	if math.Abs(result.NetContribution-20) > 0.001 {
		t.Fatalf("expected net contribution 20, got %.4f", result.NetContribution)
	}
}

func TestAnalyzeSortsByContributionMargin(t *testing.T) {
	output, err := gmcontribution.Analyze(context.Background(), gmcontribution.Input{
		Items: []gmcontribution.Item{
			{SKU: "B", Name: "Beta", Revenue: 100, VariableCost: 90},
			{SKU: "A", Name: "Alpha", Revenue: 100, VariableCost: 40},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].SKU != "A" {
		t.Fatalf("expected first SKU A, got %s", output.Results[0].SKU)
	}
}
