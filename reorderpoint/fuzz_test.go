package reorderpoint

import (
	"context"
	"math"
	"testing"
)

func FuzzAnalyze(f *testing.F) {
	f.Add(20.0, 4.0, 33.0)
	f.Add(0.0, 4.0, 33.0)
	f.Add(-1.0, 4.0, 33.0)

	f.Fuzz(func(t *testing.T, averageDemand float64, leadTime float64, safetyStock float64) {
		averageDemand = normalizeFiniteFloat(averageDemand, 1_000_000)
		leadTime = normalizeFiniteFloat(leadTime, 1_000_000)
		safetyStock = normalizeFiniteFloat(safetyStock, 1_000_000)

		output, err := Analyze(context.Background(), Input{
			Items: []Item{
				{
					SKU:                    "A",
					Name:                   "Alpha",
					AverageDemandPerPeriod: averageDemand,
					LeadTimePeriods:        leadTime,
					SafetyStock:            safetyStock,
				},
			},
		})

		if !isFinite(averageDemand) || !isFinite(leadTime) || !isFinite(safetyStock) || averageDemand < 0 || leadTime < 0 || safetyStock < 0 {
			if err == nil {
				t.Fatalf("expected validation error for input %+v", []float64{averageDemand, leadTime, safetyStock})
			}
			return
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result := output.Results[0]
		if !isFinite(result.LeadTimeDemand) || !isFinite(result.ReorderPoint) {
			t.Fatalf("expected finite results, got %+v", result)
		}
		if result.LeadTimeDemand < 0 || result.ReorderPoint < 0 {
			t.Fatalf("expected non-negative results, got %+v", result)
		}
		if result.ReorderPoint < result.SafetyStock {
			t.Fatalf("expected reorder point >= safety stock, got %+v", result)
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
