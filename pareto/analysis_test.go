package pareto_test

import (
	"context"
	"math"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/pareto"
)

func TestAnalyzeBuildsParetoSummary(t *testing.T) {
	output, err := pareto.Analyze(context.Background(), pareto.Input{
		Items: []pareto.Item{
			{SKU: "A", Name: "Alpha", Value: 80},
			{SKU: "B", Name: "Beta", Value: 15},
			{SKU: "C", Name: "Gamma", Value: 5},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if !output.Results[0].InTopValueSet {
		t.Fatal("expected first item to be in top value set")
	}
	if !output.Results[0].InTopItemSet {
		t.Fatal("expected first item to be in top item set")
	}
	if output.Results[1].InTopItemSet {
		t.Fatal("expected second item to be outside top 20 percent item set")
	}
	if math.Abs(output.Summary.TopItemValueShare-80) > 0.001 {
		t.Fatalf("expected top item value share 80, got %.4f", output.Summary.TopItemValueShare)
	}
	if !output.Summary.ParetoPrincipleMet {
		t.Fatal("expected Pareto principle to be met")
	}
}

func TestAnalyzeSupportsCustomThresholds(t *testing.T) {
	output, err := pareto.Analyze(context.Background(), pareto.Input{
		Thresholds: pareto.Thresholds{
			TopValueShare: 70,
			TopItemShare:  40,
		},
		Items: []pareto.Item{
			{SKU: "A", Name: "Alpha", Value: 60},
			{SKU: "B", Name: "Beta", Value: 25},
			{SKU: "C", Name: "Gamma", Value: 15},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Summary.TopItemCount != 2 {
		t.Fatalf("expected top item count 2, got %d", output.Summary.TopItemCount)
	}
	if !output.Summary.ParetoPrincipleMet {
		t.Fatal("expected custom Pareto rule to be met")
	}
}
