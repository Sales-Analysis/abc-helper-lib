package clv_test

import (
	"context"
	"math"
	"strings"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/clv"
)

func TestAnalyzeCalculatesPredictiveCLV(t *testing.T) {
	output, err := clv.Analyze(context.Background(), clv.Input{
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
		t.Fatalf("Analyze returned error: %v", err)
	}

	result := output.Results[0]
	if math.Abs(result.AverageOrderValue-100) > 0.001 {
		t.Fatalf("expected AOV 100, got %.4f", result.AverageOrderValue)
	}
	if math.Abs(result.PurchaseFrequency-2) > 0.001 {
		t.Fatalf("expected purchase frequency 2, got %.4f", result.PurchaseFrequency)
	}
	if math.Abs(result.PredictedCLV-163.3333) > 0.01 {
		t.Fatalf("expected predicted CLV around 163.33, got %.4f", result.PredictedCLV)
	}
}

func TestAnalyzeNormalizesPercentInputs(t *testing.T) {
	output, err := clv.Analyze(context.Background(), clv.Input{
		Customers: []clv.Customer{
			{
				CustomerID:      "A",
				Name:            "Alpha",
				Revenue:         1000,
				Orders:          10,
				PeriodsObserved: 5,
				GrossMarginRate: 40,
				RetentionRate:   80,
				DiscountRate:    10,
			},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if math.Abs(output.Results[0].GrossMarginRate-0.4) > 0.001 {
		t.Fatalf("expected normalized gross margin rate 0.4, got %.4f", output.Results[0].GrossMarginRate)
	}
	if math.Abs(output.Results[0].RetentionRate-0.8) > 0.001 {
		t.Fatalf("expected normalized retention rate 0.8, got %.4f", output.Results[0].RetentionRate)
	}
}

func TestAnalyzeReturnsErrorOnInvalidRates(t *testing.T) {
	_, err := clv.Analyze(context.Background(), clv.Input{
		Customers: []clv.Customer{
			{
				CustomerID:      "A",
				Name:            "Alpha",
				Revenue:         1000,
				Orders:          10,
				PeriodsObserved: 5,
				GrossMarginRate: 120,
			},
		},
	})
	if err == nil {
		t.Fatal("expected invalid input error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}

func TestAnalyzeReturnsErrorOnInvalidMoneyFields(t *testing.T) {
	tests := []struct {
		name     string
		customer clv.Customer
		field    string
	}{
		{
			name: "revenue",
			customer: clv.Customer{
				CustomerID:      "A",
				Name:            "Alpha",
				Revenue:         math.NaN(),
				Orders:          10,
				PeriodsObserved: 5,
				GrossMarginRate: 0.4,
				RetentionRate:   0.8,
				DiscountRate:    0.1,
			},
			field: "customers[].Revenue",
		},
		{
			name: "acquisition cost",
			customer: clv.Customer{
				CustomerID:      "A",
				Name:            "Alpha",
				Revenue:         1000,
				Orders:          10,
				PeriodsObserved: 5,
				GrossMarginRate: 0.4,
				RetentionRate:   0.8,
				DiscountRate:    0.1,
				AcquisitionCost: -1,
			},
			field: "customers[].AcquisitionCost",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := clv.Analyze(context.Background(), clv.Input{
				Customers: []clv.Customer{tt.customer},
			})
			if err == nil {
				t.Fatal("expected invalid input error, got nil")
			}
			if !strings.Contains(err.Error(), tt.field) {
				t.Fatalf("expected %s error, got %v", tt.field, err)
			}
		})
	}
}
