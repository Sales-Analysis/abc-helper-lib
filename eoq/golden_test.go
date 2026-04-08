package eoq_test

import (
	"context"
	"math"
	"path/filepath"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/eoq"
	"github.com/Sales-Analysis/abc-helper-lib/internal/testutil"
)

func TestAnalyzeMatchesGoldenFile(t *testing.T) {
	output, err := eoq.Analyze(context.Background(), eoq.Input{
		Items: []eoq.Item{
			{SKU: "A", Name: "Alpha", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	for i := range output.Results {
		output.Results[i].AnnualDemand = round(output.Results[i].AnnualDemand, 4)
		output.Results[i].OrderingCost = round(output.Results[i].OrderingCost, 4)
		output.Results[i].HoldingCost = round(output.Results[i].HoldingCost, 4)
		output.Results[i].OptimalQuantity = round(output.Results[i].OptimalQuantity, 4)
		output.Results[i].OrdersPerYear = round(output.Results[i].OrdersPerYear, 4)
	}

	testutil.AssertGoldenJSON(t, filepath.Join("testdata", "analyze.golden.json"), output)
}

func round(value float64, places int) float64 {
	factor := math.Pow10(places)
	return math.Round(value*factor) / factor
}
