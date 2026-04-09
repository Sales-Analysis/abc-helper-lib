package churn_test

import (
	"context"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/churn"
)

func TestAnalyzeBuildsChurnSummary(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := churn.Analyze(context.Background(), churn.Input{
		AnalysisTime: now,
		Customers: []churn.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, -10)},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -45)},
			{CustomerID: "C", Name: "Gamma", LastOrderAt: now.AddDate(0, 0, -120)},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(output.Results))
	}
	if output.Results[0].Status != "Churned" {
		t.Fatalf("expected most stale customer to be Churned, got %s", output.Results[0].Status)
	}
	if output.Summary.ActiveCount != 1 {
		t.Fatalf("expected active count 1, got %d", output.Summary.ActiveCount)
	}
	if output.Summary.AtRiskCount != 1 {
		t.Fatalf("expected at risk count 1, got %d", output.Summary.AtRiskCount)
	}
	if output.Summary.ChurnedCount != 1 {
		t.Fatalf("expected churned count 1, got %d", output.Summary.ChurnedCount)
	}
	if math.Abs(output.Summary.RetentionRate-66.6667) > 0.01 {
		t.Fatalf("expected retention rate around 66.67, got %.4f", output.Summary.RetentionRate)
	}
}

func TestAnalyzeSupportsCustomThresholds(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := churn.Analyze(context.Background(), churn.Input{
		AnalysisTime: now,
		Thresholds: churn.Thresholds{
			AtRiskDays: 15,
			ChurnDays:  30,
		},
		Customers: []churn.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, -20)},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -40)},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].Status != "Churned" {
		t.Fatalf("expected first customer status Churned, got %s", output.Results[0].Status)
	}
	if output.Results[1].Status != "AtRisk" {
		t.Fatalf("expected second customer status AtRisk, got %s", output.Results[1].Status)
	}
}

func TestAnalyzeTreatsMissingLastOrderAsChurnBoundary(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := churn.Analyze(context.Background(), churn.Input{
		AnalysisTime: now,
		Customers: []churn.Customer{
			{CustomerID: "A", Name: "Alpha"},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -10)},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[0].CustomerID != "A" {
		t.Fatalf("expected missing-date customer first, got %s", output.Results[0].CustomerID)
	}
	if output.Results[0].DaysSinceLastOrder != 90 {
		t.Fatalf("expected missing-date customer days 90, got %d", output.Results[0].DaysSinceLastOrder)
	}
	if output.Results[0].Status != "Churned" {
		t.Fatalf("expected missing-date customer status Churned, got %s", output.Results[0].Status)
	}
}

func TestAnalyzeReturnsErrorOnInvalidThresholds(t *testing.T) {
	_, err := churn.Analyze(context.Background(), churn.Input{
		Thresholds: churn.Thresholds{
			AtRiskDays: 45,
			ChurnDays:  30,
		},
	})
	if err == nil {
		t.Fatal("expected invalid thresholds error, got nil")
	}
	if !strings.Contains(err.Error(), "invalid input") {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}

func TestAnalyzeReturnsErrorOnFutureLastOrder(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	_, err := churn.Analyze(context.Background(), churn.Input{
		AnalysisTime: now,
		Customers: []churn.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, 1)},
		},
	})
	if err == nil {
		t.Fatal("expected invalid input error, got nil")
	}
	if !strings.Contains(err.Error(), "customers[].LastOrderAt") {
		t.Fatalf("expected last order error, got %v", err)
	}
}
