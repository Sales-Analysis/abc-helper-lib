package churn_test

import (
	"context"
	"math"
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
