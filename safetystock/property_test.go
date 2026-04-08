package safetystock

import (
	"math"
	"testing"
	"testing/quick"
)

func TestCalculateSafetyStockMatchesFormula(t *testing.T) {
	err := quick.Check(func(demandStdDev float64, leadTime float64, serviceFactor float64) bool {
		demandStdDev = normalizePositive(demandStdDev)
		leadTime = normalizePositive(leadTime)
		serviceFactor = normalizePositive(serviceFactor)

		safetyStock := calculateSafetyStock(demandStdDev, leadTime, serviceFactor)
		expected := serviceFactor * demandStdDev * math.Sqrt(leadTime)
		return almostEqual(safetyStock, expected, 1e-6)
	}, nil)
	if err != nil {
		t.Fatalf("safety stock formula property failed: %v", err)
	}
}

func normalizePositive(value float64) float64 {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 1
	}
	value = math.Abs(math.Mod(value, 1_000_000))
	if value == 0 {
		return 1
	}
	return value
}

func almostEqual(left float64, right float64, relativeTolerance float64) bool {
	diff := math.Abs(left - right)
	scale := math.Max(1, math.Max(math.Abs(left), math.Abs(right)))
	return diff <= relativeTolerance*scale
}
