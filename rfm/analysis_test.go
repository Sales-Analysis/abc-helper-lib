package rfm_test

import (
	"context"
	"testing"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/rfm"
)

func TestAnalyzeAssignsRFMScoreAndSegment(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := rfm.Analyze(context.Background(), rfm.Input{
		AnalysisTime: now,
		Customers: []rfm.Customer{
			{CustomerID: "A", Name: "Alpha", LastOrderAt: now.AddDate(0, 0, -5), Orders: 10, MonetaryValue: 1000},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -10), Orders: 8, MonetaryValue: 800},
			{CustomerID: "C", Name: "Gamma", LastOrderAt: now.AddDate(0, 0, -20), Orders: 5, MonetaryValue: 500},
			{CustomerID: "D", Name: "Delta", LastOrderAt: now.AddDate(0, 0, -40), Orders: 3, MonetaryValue: 300},
			{CustomerID: "E", Name: "Epsilon", LastOrderAt: now.AddDate(0, 0, -120), Orders: 1, MonetaryValue: 50},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if len(output.Results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(output.Results))
	}
	if output.Results[0].CustomerID != "A" {
		t.Fatalf("expected first customer A, got %s", output.Results[0].CustomerID)
	}
	if output.Results[0].RFMScore != "555" {
		t.Fatalf("expected first customer score 555, got %s", output.Results[0].RFMScore)
	}
	if output.Results[0].Segment != "Champions" {
		t.Fatalf("expected first customer segment Champions, got %s", output.Results[0].Segment)
	}
	if output.Results[4].Segment != "Hibernating" {
		t.Fatalf("expected last customer segment Hibernating, got %s", output.Results[4].Segment)
	}
}

func TestAnalyzeTreatsMissingLastOrderAsVeryStale(t *testing.T) {
	now := time.Date(2026, time.April, 6, 0, 0, 0, 0, time.UTC)

	output, err := rfm.Analyze(context.Background(), rfm.Input{
		AnalysisTime: now,
		Customers: []rfm.Customer{
			{CustomerID: "A", Name: "Alpha", Orders: 5, MonetaryValue: 500},
			{CustomerID: "B", Name: "Beta", LastOrderAt: now.AddDate(0, 0, -5), Orders: 4, MonetaryValue: 400},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	if output.Results[1].CustomerID != "A" {
		t.Fatalf("expected stale customer A to rank lower, got %s", output.Results[1].CustomerID)
	}
}
