package servicelevel

import (
	"context"
	"math"
	"testing"
)

func FuzzAnalyze(f *testing.F) {
	f.Add(100.0, 95.0, 20, 2)
	f.Add(50.0, 45.0, 0, 0)
	f.Add(100.0, 110.0, 20, 2)

	f.Fuzz(func(t *testing.T, demanded float64, fulfilled float64, totalCycles int, stockoutCycles int) {
		demanded = normalizeFiniteFloat(demanded, 1_000_000)
		fulfilled = normalizeFiniteFloat(fulfilled, 1_000_000)
		totalCycles = normalizeInt(totalCycles, 10_000)
		stockoutCycles = normalizeInt(stockoutCycles, 10_000)

		output, err := Analyze(context.Background(), Input{
			Items: []Item{
				{
					SKU:            "A",
					Name:           "Alpha",
					DemandedUnits:  demanded,
					FulfilledUnits: fulfilled,
					TotalCycles:    totalCycles,
					StockoutCycles: stockoutCycles,
				},
			},
		})

		invalid := !isFinite(demanded) || !isFinite(fulfilled) || demanded < 0 || fulfilled < 0 || totalCycles < 0 || stockoutCycles < 0
		invalid = invalid || (totalCycles > 0 && stockoutCycles > totalCycles)
		invalid = invalid || (demanded > 0 && fulfilled > demanded)

		if invalid {
			if err == nil {
				t.Fatalf("expected validation error for input demanded=%.4f fulfilled=%.4f totalCycles=%d stockoutCycles=%d", demanded, fulfilled, totalCycles, stockoutCycles)
			}
			return
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result := output.Results[0]
		if !isFinite(result.UnfulfilledUnits) || result.UnfulfilledUnits < 0 {
			t.Fatalf("expected non-negative unfulfilled units, got %+v", result)
		}
		for _, value := range []float64{result.FillRate, result.CycleServiceLevel, result.StockoutRate, result.ServiceLevel} {
			if !isFinite(value) {
				t.Fatalf("expected finite metric, got %+v", result)
			}
			if value < 0 || value > 100 {
				t.Fatalf("expected percentage metric in [0,100], got %+v", result)
			}
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

func normalizeInt(value int, limit int) int {
	if limit <= 0 {
		return value
	}
	return value % limit
}
