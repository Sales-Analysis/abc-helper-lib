package safetystock

import (
	"context"
	"math"
	"testing"
)

func FuzzAnalyze(f *testing.F) {
	f.Add(10.0, 4.0, 1.65)
	f.Add(10.0, 4.0, 0.0)
	f.Add(-1.0, 4.0, 1.65)

	f.Fuzz(func(t *testing.T, demandStdDev float64, leadTime float64, serviceFactor float64) {
		demandStdDev = normalizeFiniteFloat(demandStdDev, 1_000_000)
		leadTime = normalizeFiniteFloat(leadTime, 1_000_000)
		serviceFactor = normalizeFiniteFloat(serviceFactor, 1_000_000)

		output, err := Analyze(context.Background(), Input{
			Items: []Item{
				{
					SKU:             "A",
					Name:            "Alpha",
					DemandStdDev:    demandStdDev,
					LeadTimePeriods: leadTime,
					ServiceFactor:   serviceFactor,
				},
			},
		})

		if !isFinite(demandStdDev) || !isFinite(leadTime) || !isFinite(serviceFactor) || demandStdDev < 0 || leadTime < 0 || serviceFactor < 0 {
			if err == nil {
				t.Fatalf("expected validation error for input %+v", []float64{demandStdDev, leadTime, serviceFactor})
			}
			return
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result := output.Results[0]
		if !isFinite(result.ServiceFactor) || !isFinite(result.SafetyStock) {
			t.Fatalf("expected finite results, got %+v", result)
		}
		if result.SafetyStock < 0 {
			t.Fatalf("expected non-negative safety stock, got %+v", result)
		}
		if serviceFactor == 0 && result.ServiceFactor != defaultServiceFactor {
			t.Fatalf("expected default service factor %.2f, got %.4f", defaultServiceFactor, result.ServiceFactor)
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
