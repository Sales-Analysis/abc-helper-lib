package reorderpoint

import (
	"math"
	"testing"
	"testing/quick"
)

func TestCalculateLeadTimeDemandMatchesMultiplication(t *testing.T) {
	err := quick.Check(func(averageDemand float64, leadTime float64) bool {
		averageDemand = normalizeNonNegative(averageDemand)
		leadTime = normalizeNonNegative(leadTime)

		leadTimeDemand := calculateLeadTimeDemand(averageDemand, leadTime)
		if averageDemand == 0 || leadTime == 0 {
			return leadTimeDemand == 0
		}
		return almostEqual(leadTimeDemand, averageDemand*leadTime, 1e-6)
	}, nil)
	if err != nil {
		t.Fatalf("lead-time demand property failed: %v", err)
	}
}

func normalizeNonNegative(value float64) float64 {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0
	}
	return math.Abs(math.Mod(value, 1_000_000))
}

func almostEqual(left float64, right float64, relativeTolerance float64) bool {
	diff := math.Abs(left - right)
	scale := math.Max(1, math.Max(math.Abs(left), math.Abs(right)))
	return diff <= relativeTolerance*scale
}
