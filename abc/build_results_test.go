package abc

import (
	"reflect"
	"testing"
)

func TestBuildResults(t *testing.T) {
	products := []Product{
		{SKU: "001", Name: "Product A", Quantity: 2, Price: 50}, // total 100
		{SKU: "002", Name: "Product B", Quantity: 1, Price: 30}, // total 30
	}

	priceTotal := []float64{100, 30}
	pairs := byValue{
		{value: 100, index: 0},
		{value: 30, index: 1},
	}
	costPercentage := []float64{76.92307692, 23.07692308}
	accumulatedShare := []float64{76.92307692, 100}
	groups := []string{"A", "B"}

	a := New()
	results := a.buildResults(products, priceTotal, pairs, costPercentage, accumulatedShare, groups)

	expected := []ProductResult{
		{
			SKU:              "001",
			Name:             "Product A",
			Quantity:         2,
			PriceUnit:        50,
			PriceTotal:       100,
			ShareTotal:       costPercentage[0],
			ShareAccumulated: accumulatedShare[0],
			Group:            "A",
		},
		{
			SKU:              "002",
			Name:             "Product B",
			Quantity:         1,
			PriceUnit:        30,
			PriceTotal:       30,
			ShareTotal:       costPercentage[1],
			ShareAccumulated: accumulatedShare[1],
			Group:            "B",
		},
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %+v, got %+v", expected, results)
	}
}
