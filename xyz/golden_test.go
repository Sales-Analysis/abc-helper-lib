package xyz_test

import (
	"context"
	"math"
	"path/filepath"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/internal/testutil"
	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

func TestAnalyzeMatchesGoldenFile(t *testing.T) {
	output, err := xyz.Analyze(context.Background(), xyz.Input{
		Items: []xyz.Item{
			{SKU: "X", Name: "Stable", Demands: []float64{10, 10, 10}},
			{SKU: "Y", Name: "Moderate", Demands: []float64{10, 12, 8}},
			{SKU: "Z", Name: "Erratic", Demands: []float64{0, 20, 0}},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	for i := range output.Results {
		output.Results[i].AverageDemand = round(output.Results[i].AverageDemand, 4)
		output.Results[i].StandardDeviation = round(output.Results[i].StandardDeviation, 4)
		output.Results[i].CoefficientOfVariation = round(output.Results[i].CoefficientOfVariation, 4)
	}

	testutil.AssertGoldenJSON(t, filepath.Join("testdata", "analyze.golden.json"), output)
}

func round(value float64, places int) float64 {
	factor := math.Pow10(places)
	return math.Round(value*factor) / factor
}
