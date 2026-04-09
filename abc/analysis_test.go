package abc

import (
	"context"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestABC(t *testing.T) {
	products := []Product{
		{SKU: "001", Name: "Product A", Quantity: 10, Price: 10},
		{SKU: "002", Name: "Product B", Quantity: 5, Price: 20},
		{SKU: "003", Name: "Product C", Quantity: 1, Price: 5},
	}

	a := New()
	a.Calculate(products)

	if len(a.Result) != len(products) {
		t.Fatalf("Expected %d results, got %d", len(products), len(a.Result))
	}
	if a.Result[0].PriceTotal < a.Result[1].PriceTotal {
		t.Errorf("Results not sorted by PriceTotal descending")
	}

	expectedTotals := []float64{100, 100, 5}
	for _, res := range a.Result {
		found := false
		for _, total := range expectedTotals {
			if res.PriceTotal == total {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected PriceTotal in results: %v", res.PriceTotal)
		}
	}

	groups := map[string]bool{}
	for _, res := range a.Result {
		groups[res.Group] = true
	}
	if len(groups) < 2 {
		t.Errorf("Expected at least 2 different groups, got: %v", groups)
	}
}

func TestAnalyzeMatchesLegacyCalculate(t *testing.T) {
	products := []Product{
		{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
		{SKU: "B", Name: "Beta", Quantity: 2, Price: 10},
		{SKU: "C", Name: "Gamma", Quantity: 1, Price: 5},
	}

	legacy := New()
	legacy.Calculate(products)

	output, err := Analyze(context.Background(), Input{Products: products})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if !reflect.DeepEqual(legacy.Result, output.Results) {
		t.Fatalf("Analyze results differ from legacy API.\nlegacy=%+v\nnew=%+v", legacy.Result, output.Results)
	}
	if output.TotalRevenue != 125 {
		t.Fatalf("expected total revenue 125, got %.2f", output.TotalRevenue)
	}
	if output.Summary.TotalItems != 3 || output.Summary.TotalRevenue != 125 {
		t.Fatalf("unexpected summary: %+v", output.Summary)
	}
	if output.Summary.ACount != 1 || output.Summary.BCount != 0 || output.Summary.CCount != 2 {
		t.Fatalf("unexpected group summary: %+v", output.Summary)
	}
}

func TestAnalyzeAcceptsItemsField(t *testing.T) {
	output, err := Analyze(context.Background(), Input{
		Items: []Item{
			{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
			{SKU: "B", Name: "Beta", Quantity: 1, Price: 5},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].SKU != "A" {
		t.Fatalf("expected SKU A first, got %s", output.Results[0].SKU)
	}
	if output.Results[0].OriginalIndex != 0 || output.Results[1].OriginalIndex != 1 {
		t.Fatalf("unexpected original indexes: %+v", output.Results)
	}
}

func TestAnalyzePrefersItemsOverLegacyProducts(t *testing.T) {
	output, err := Analyze(context.Background(), Input{
		Items: []Item{
			{SKU: "NEW", Name: "Preferred", Quantity: 2, Price: 50},
		},
		Products: []Product{
			{SKU: "OLD", Name: "Legacy", Quantity: 1, Price: 1},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}
	if output.Results[0].SKU != "NEW" {
		t.Fatalf("expected Items field to take precedence, got %s", output.Results[0].SKU)
	}
}

func TestAnalyzeDetailedKeepsOriginalIndex(t *testing.T) {
	products := []Product{
		{SKU: "A", Name: "Alpha", Quantity: 1, Price: 10},
		{SKU: "B", Name: "Beta", Quantity: 10, Price: 10},
	}

	output, err := AnalyzeDetailed(context.Background(), Input{Products: products})
	if err != nil {
		t.Fatalf("AnalyzeDetailed returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].OriginalIndex != 1 {
		t.Fatalf("expected first result to point to original index 1, got %d", output.Results[0].OriginalIndex)
	}
	if output.Results[1].OriginalIndex != 0 {
		t.Fatalf("expected second result to point to original index 0, got %d", output.Results[1].OriginalIndex)
	}
}

func TestBuildResults(t *testing.T) {
	products := []Product{
		{SKU: "001", Name: "Product A", Quantity: 2, Price: 50},
		{SKU: "002", Name: "Product B", Quantity: 1, Price: 30},
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
			OriginalIndex:    0,
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
			OriginalIndex:    1,
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

func TestAnalyzeReturnsErrorOnInvalidThresholds(t *testing.T) {
	_, err := Analyze(context.Background(), Input{
		Thresholds: Thresholds{
			AMaxShare: 90,
			BMaxShare: 80,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}

func TestAnalyzeReturnsErrorOnInvalidProductValues(t *testing.T) {
	_, err := Analyze(context.Background(), Input{
		Items: []Item{
			{SKU: "A", Name: "Alpha", Quantity: -1, Price: 10},
		},
	})
	if err == nil {
		t.Fatal("expected invalid product values error, got nil")
	}
	if !strings.Contains(err.Error(), "items[].Quantity") {
		t.Fatalf("expected quantity error, got %v", err)
	}
}

func TestAnalyzeHandlesZeroGrandTotalWithoutNaN(t *testing.T) {
	output, err := Analyze(context.Background(), Input{
		Items: []Item{
			{SKU: "A", Name: "Alpha", Quantity: 0, Price: 100},
			{SKU: "B", Name: "Beta", Quantity: 0, Price: 50},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.TotalRevenue != 0 {
		t.Fatalf("expected zero total revenue, got %.2f", output.TotalRevenue)
	}
	for _, result := range output.Results {
		if math.IsNaN(result.ShareTotal) || math.IsInf(result.ShareTotal, 0) {
			t.Fatalf("expected finite ShareTotal, got %+v", result)
		}
		if math.IsNaN(result.ShareAccumulated) || math.IsInf(result.ShareAccumulated, 0) {
			t.Fatalf("expected finite ShareAccumulated, got %+v", result)
		}
		if result.ShareTotal != 0 {
			t.Fatalf("expected zero ShareTotal, got %.4f", result.ShareTotal)
		}
		if result.ShareAccumulated != 0 {
			t.Fatalf("expected zero ShareAccumulated, got %.4f", result.ShareAccumulated)
		}
	}
}
