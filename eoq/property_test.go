package eoq

import (
	"math"
	"testing"
	"testing/quick"
)

func TestCalculateEOQSatisfiesClosedForm(t *testing.T) {
	err := quick.Check(func(annualDemand float64, orderingCost float64, holdingCost float64) bool {
		annualDemand = normalizePositive(annualDemand)
		orderingCost = normalizePositive(orderingCost)
		holdingCost = normalizePositive(holdingCost)

		eoq := calculateEOQ(annualDemand, orderingCost, holdingCost)
		return almostEqual(eoq*eoq*holdingCost, 2*annualDemand*orderingCost, 1e-6)
	}, nil)
	if err != nil {
		t.Fatalf("closed-form property failed: %v", err)
	}
}

func TestCalculateOrdersPerYearInvertsQuantity(t *testing.T) {
	err := quick.Check(func(annualDemand float64, orderingCost float64, holdingCost float64) bool {
		annualDemand = normalizePositive(annualDemand)
		orderingCost = normalizePositive(orderingCost)
		holdingCost = normalizePositive(holdingCost)

		eoq := calculateEOQ(annualDemand, orderingCost, holdingCost)
		ordersPerYear := calculateOrdersPerYear(annualDemand, eoq)
		return almostEqual(ordersPerYear*eoq, annualDemand, 1e-6)
	}, nil)
	if err != nil {
		t.Fatalf("inverse property failed: %v", err)
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
