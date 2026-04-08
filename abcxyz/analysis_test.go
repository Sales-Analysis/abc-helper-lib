package abcxyz_test

import (
	"context"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/abcxyz"
)

func TestAnalyzeCombinesABCAndXYZGroups(t *testing.T) {
	output, err := abcxyz.Analyze(context.Background(), abcxyz.Input{
		ABCThresholds: abc.Thresholds{
			AMaxShare: 80,
			BMaxShare: 98,
		},
		Items: []abcxyz.Item{
			{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10, Demands: []float64{100, 100, 100}},
			{SKU: "B", Name: "Beta", Quantity: 3, Price: 10, Demands: []float64{100, 120, 80}},
			{SKU: "C", Name: "Gamma", Quantity: 1, Price: 5, Demands: []float64{0, 200, 0}},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}

	if output.Results[0].CombinedGroup != "AX" {
		t.Fatalf("expected first combined group AX, got %s", output.Results[0].CombinedGroup)
	}
	if output.Results[1].CombinedGroup != "BY" {
		t.Fatalf("expected second combined group BY, got %s", output.Results[1].CombinedGroup)
	}
	if output.Results[2].CombinedGroup != "CZ" {
		t.Fatalf("expected third combined group CZ, got %s", output.Results[2].CombinedGroup)
	}
	if output.Summary.TotalItems != 3 {
		t.Fatalf("expected summary total items 3, got %d", output.Summary.TotalItems)
	}
	if output.Summary.ABCGroupCounts["A"] != 1 || output.Summary.ABCGroupCounts["B"] != 1 || output.Summary.ABCGroupCounts["C"] != 1 {
		t.Fatalf("unexpected ABC summary: %+v", output.Summary.ABCGroupCounts)
	}
	if output.Summary.XYZGroupCounts["X"] != 1 || output.Summary.XYZGroupCounts["Y"] != 1 || output.Summary.XYZGroupCounts["Z"] != 1 {
		t.Fatalf("unexpected XYZ summary: %+v", output.Summary.XYZGroupCounts)
	}
	if output.Summary.CombinedGroupCounts["AX"] != 1 || output.Summary.CombinedGroupCounts["BY"] != 1 || output.Summary.CombinedGroupCounts["CZ"] != 1 {
		t.Fatalf("unexpected combined summary: %+v", output.Summary.CombinedGroupCounts)
	}
}

func TestAnalyzeReturnsErrorOnInvalidNestedThresholds(t *testing.T) {
	_, err := abcxyz.Analyze(context.Background(), abcxyz.Input{
		ABCThresholds: abc.Thresholds{
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
