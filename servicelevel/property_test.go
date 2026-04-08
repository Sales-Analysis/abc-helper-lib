package servicelevel

import (
	"testing"
	"testing/quick"
)

func TestCycleServiceLevelAndStockoutRateSumToHundred(t *testing.T) {
	err := quick.Check(func(totalCycles uint16, stockoutCycles uint16) bool {
		total := int(totalCycles%1000) + 1
		stockouts := int(stockoutCycles % uint16(total+1))

		cycleServiceLevel := calculateCycleServiceLevel(total, stockouts)
		stockoutRate := calculateStockoutRate(total, stockouts)
		return cycleServiceLevel >= 0 && cycleServiceLevel <= 100 &&
			stockoutRate >= 0 && stockoutRate <= 100 &&
			almostEqual(cycleServiceLevel+stockoutRate, 100, 1e-9)
	}, nil)
	if err != nil {
		t.Fatalf("cycle-level complement property failed: %v", err)
	}
}

func TestSafeUnitRateStaysWithinBounds(t *testing.T) {
	err := quick.Check(func(total uint32, fulfilled uint32) bool {
		normalizedTotal := total % 1_000_000
		demanded := float64(normalizedTotal)
		if normalizedTotal == 0 {
			return safeUnitRate(float64(fulfilled), demanded) == 0
		}

		served := float64(fulfilled % (normalizedTotal + 1))
		fillRate := safeUnitRate(served, demanded)
		return fillRate >= 0 && fillRate <= 100
	}, nil)
	if err != nil {
		t.Fatalf("fill-rate bounds property failed: %v", err)
	}
}

func almostEqual(left float64, right float64, tolerance float64) bool {
	diff := left - right
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}
