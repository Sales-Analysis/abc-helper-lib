package customer_test

import (
	"context"
	"testing"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/churn"
	"github.com/Sales-Analysis/abc-helper-lib/clv"
	"github.com/Sales-Analysis/abc-helper-lib/customer"
	"github.com/Sales-Analysis/abc-helper-lib/rfm"
)

func TestFacadeRFMDelegatesToAnalyzer(t *testing.T) {
	facade := customer.New()
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := facade.RFM(context.Background(), rfm.Input{
		AnalysisTime: now,
		Customers: []rfm.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, -5), Orders: 10, MonetaryValue: 1000},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -50), Orders: 1, MonetaryValue: 100},
		},
	})
	if err != nil {
		t.Fatalf("RFM returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Results[0].CustomerID != "A" {
		t.Fatalf("expected first customer A, got %s", output.Results[0].CustomerID)
	}
}

func TestFacadeCLVDelegatesToAnalyzer(t *testing.T) {
	facade := customer.New()

	output, err := facade.CLV(context.Background(), clv.Input{
		Customers: []clv.Customer{
			{
				CustomerID:      "A",
				Name:            "Alpha",
				Revenue:         1000,
				Orders:          10,
				PeriodsObserved: 5,
				GrossMarginRate: 0.4,
				RetentionRate:   0.8,
				DiscountRate:    0.1,
				AcquisitionCost: 50,
			},
		},
	})
	if err != nil {
		t.Fatalf("CLV returned error: %v", err)
	}

	if len(output.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(output.Results))
	}
	if output.Results[0].PredictedCLV <= 0 {
		t.Fatalf("expected positive CLV, got %.4f", output.Results[0].PredictedCLV)
	}
}

func TestFacadeChurnDelegatesToAnalyzer(t *testing.T) {
	facade := customer.New()
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := facade.Churn(context.Background(), churn.Input{
		AnalysisTime: now,
		Customers: []churn.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, -10)},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -120)},
		},
	})
	if err != nil {
		t.Fatalf("Churn returned error: %v", err)
	}

	if len(output.Results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(output.Results))
	}
	if output.Summary.ChurnedCount != 1 {
		t.Fatalf("expected churned count 1, got %d", output.Summary.ChurnedCount)
	}
}
