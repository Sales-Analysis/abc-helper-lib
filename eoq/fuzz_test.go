package eoq

import (
	"context"
	"math"
	"testing"
)

func FuzzAnalyze(f *testing.F) {
	f.Add(1200.0, 50.0, 2.0)
	f.Add(0.0, 50.0, 2.0)
	f.Add(-1.0, 50.0, 2.0)

	f.Fuzz(func(t *testing.T, annualDemand float64, orderingCost float64, holdingCost float64) {
		annualDemand = normalizeFiniteFloat(annualDemand, 1_000_000)
		orderingCost = normalizeFiniteFloat(orderingCost, 1_000_000)
		holdingCost = normalizeFiniteFloat(holdingCost, 1_000_000)

		output, err := Analyze(context.Background(), Input{
			Items: []Item{
				{
					SKU:          "A",
					Name:         "Alpha",
					AnnualDemand: annualDemand,
					OrderingCost: orderingCost,
					HoldingCost:  holdingCost,
				},
			},
		})

		if !isFinite(annualDemand) || !isFinite(orderingCost) || !isFinite(holdingCost) || annualDemand < 0 || orderingCost < 0 || holdingCost < 0 {
			if err == nil {
				t.Fatalf("expected validation error for input %+v", []float64{annualDemand, orderingCost, holdingCost})
			}
			return
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result := output.Results[0]
		if !isFinite(result.OptimalQuantity) || !isFinite(result.OrdersPerYear) {
			t.Fatalf("expected finite results, got %+v", result)
		}
		if result.OptimalQuantity < 0 || result.OrdersPerYear < 0 {
			t.Fatalf("expected non-negative results, got %+v", result)
		}
	})
}

func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func normalizeFiniteFloat(value float64, limit float64) float64 {
	if !isFinite(value) {
		return value
	}
	return math.Mod(value, limit)
}
