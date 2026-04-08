package pareto_test

import (
	"context"
	"math"
	"path/filepath"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/internal/testutil"
	"github.com/Sales-Analysis/abc-helper-lib/pareto"
)

func TestAnalyzeMatchesGoldenFile(t *testing.T) {
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

	for i := range output.Results {
		output.Results[i].ValueShare = round(output.Results[i].ValueShare, 4)
		output.Results[i].CumulativeValueShare = round(output.Results[i].CumulativeValueShare, 4)
		output.Results[i].CumulativeItemShare = round(output.Results[i].CumulativeItemShare, 4)
	}
	output.Summary.TopItemShare = round(output.Summary.TopItemShare, 4)
	output.Summary.TopItemValueShare = round(output.Summary.TopItemValueShare, 4)
	output.Summary.TopValueShareTarget = round(output.Summary.TopValueShareTarget, 4)

	testutil.AssertGoldenJSON(t, filepath.Join("testdata", "analyze.golden.json"), output)
}

func round(value float64, places int) float64 {
	factor := math.Pow10(places)
	return math.Round(value*factor) / factor
}
