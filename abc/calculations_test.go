package abc

import (
	"math"
	"reflect"
	"testing"
)

func TestCalculatePriceTotal(t *testing.T) {
	a := New()
	products := []Product{
		{Quantity: 2, Price: 10},
		{Quantity: 3, Price: 5},
	}
	got := a.calculatePriceTotal(products)
	want := []float64{20, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestCalculateGrandTotal(t *testing.T) {
	a := New()
	totals := []float64{20, 15, 5}
	got := a.calculateGrandTotal(totals)
	want := 40.0
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func TestCalculateCostPercentage(t *testing.T) {
	a := New()
	pairs := byValue{
		{value: 50, index: 0},
		{value: 50, index: 1},
	}
	got := a.calculateCostPercentage(pairs, 100)
	want := []float64{50, 50}
	for i := range got {
		if math.Abs(got[i]-want[i]) > 0.0001 {
			t.Errorf("Expected %.2f, got %.2f", want[i], got[i])
		}
	}
}

func TestCalculateAccumulatedShare(t *testing.T) {
	a := New()
	percentages := []float64{50, 30, 20}
	got := a.calculateAccumulatedShare(percentages)
	want := []float64{50, 80, 100}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
