package inventory_test

import (
	"context"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/eoq"
	"github.com/Sales-Analysis/abc-helper-lib/fsn"
	"github.com/Sales-Analysis/abc-helper-lib/hml"
	"github.com/Sales-Analysis/abc-helper-lib/inventory"
	"github.com/Sales-Analysis/abc-helper-lib/reorderpoint"
	"github.com/Sales-Analysis/abc-helper-lib/safetystock"
	"github.com/Sales-Analysis/abc-helper-lib/sde"
	"github.com/Sales-Analysis/abc-helper-lib/ved"
)

func TestFacadeABCDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.ABC(context.Background(), abc.Input{
		Products: []abc.Product{
			{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
			{SKU: "B", Name: "Beta", Quantity: 1, Price: 5},
		},
	})
	if err != nil {
		t.Fatalf("ABC returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].SKU != "A" {
		t.Fatalf("expected highest-value SKU A first, got %s", output.Results[0].SKU)
	}
}

func TestFacadeVEDDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.VED(context.Background(), ved.Input{
		Items: []ved.Item{
			{SKU: "V-1", Name: "Vital", CriticalityScore: 90},
			{SKU: "D-1", Name: "Desirable", CriticalityScore: 20},
		},
	})
	if err != nil {
		t.Fatalf("VED returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "V" {
		t.Fatalf("expected first item group V, got %s", output.Results[0].Group)
	}
}

func TestFacadeFSNDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.FSN(context.Background(), fsn.Input{
		Items: []fsn.Item{
			{SKU: "F-1", Name: "Fast", Movements: []float64{1, 1, 1}},
			{SKU: "N-1", Name: "Non", Movements: []float64{0, 0, 0}},
		},
	})
	if err != nil {
		t.Fatalf("FSN returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "F" {
		t.Fatalf("expected first item group F, got %s", output.Results[0].Group)
	}
}

func TestFacadeHMLDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.HML(context.Background(), hml.Input{
		Items: []hml.Item{
			{SKU: "H-1", Name: "High", UnitCost: 300},
			{SKU: "L-1", Name: "Low", UnitCost: 20},
		},
	})
	if err != nil {
		t.Fatalf("HML returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "H" {
		t.Fatalf("expected first item group H, got %s", output.Results[0].Group)
	}
}

func TestFacadeSDEDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.SDE(context.Background(), sde.Input{
		Items: []sde.Item{
			{SKU: "S-1", Name: "Scarce", LeadTimeDays: 90},
			{SKU: "E-1", Name: "Easy", LeadTimeDays: 5},
		},
	})
	if err != nil {
		t.Fatalf("SDE returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].Group != "S" {
		t.Fatalf("expected first item group S, got %s", output.Results[0].Group)
	}
}

func TestFacadeEOQDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.EOQ(context.Background(), eoq.Input{
		Items: []eoq.Item{
			{SKU: "A", Name: "Alpha", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
		},
	})
	if err != nil {
		t.Fatalf("EOQ returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}
	if output.Results[0].OptimalQuantity <= 0 {
		t.Fatalf("expected positive EOQ, got %.4f", output.Results[0].OptimalQuantity)
	}
}

func TestFacadeReorderPointDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.ReorderPoint(context.Background(), reorderpoint.Input{
		Items: []reorderpoint.Item{
			{SKU: "A", Name: "Alpha", AverageDemandPerPeriod: 20, LeadTimePeriods: 4, SafetyStock: 33},
		},
	})
	if err != nil {
		t.Fatalf("ReorderPoint returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}
	if output.Results[0].ReorderPoint != 113 {
		t.Fatalf("expected reorder point 113, got %.4f", output.Results[0].ReorderPoint)
	}
}

func TestFacadeSafetyStockDelegatesToAnalyzer(t *testing.T) {
	facade := inventory.New()

	output, err := facade.SafetyStock(context.Background(), safetystock.Input{
		Items: []safetystock.Item{
			{SKU: "A", Name: "Alpha", DemandStdDev: 10, LeadTimePeriods: 4, ServiceFactor: 1.65},
		},
	})
	if err != nil {
		t.Fatalf("SafetyStock returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}
	if output.Results[0].SafetyStock != 33 {
		t.Fatalf("expected safety stock 33, got %.4f", output.Results[0].SafetyStock)
	}
}
