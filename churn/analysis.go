package churn

import (
	"context"
	"sort"
	"time"
)

const (
	defaultAtRiskDays = 30
	defaultChurnDays  = 90
)

type Thresholds struct {
	AtRiskDays int
	ChurnDays  int
}

type Input struct {
	Customers    []Customer
	AnalysisTime time.Time
	Thresholds   Thresholds
}

type Output struct {
	Results []CustomerResult
	Summary Summary
}

type Customer struct {
	CustomerID  string
	Name        string
	LastOrderAt time.Time
}

type CustomerResult struct {
	OriginalIndex      int
	CustomerID         string
	Name               string
	DaysSinceLastOrder int
	Status             string
	IsRetained         bool
	IsChurned          bool
}

type Summary struct {
	TotalCustomers int
	ActiveCount    int
	AtRiskCount    int
	ChurnedCount   int
	RetentionRate  float64
	ChurnRate      float64
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	analysisTime := input.AnalysisTime
	if analysisTime.IsZero() {
		analysisTime = time.Now().UTC()
	}
	thresholds := input.Thresholds.normalized()

	results := make([]CustomerResult, len(input.Customers))
	summary := Summary{TotalCustomers: len(input.Customers)}
	for i, customer := range input.Customers {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}

		days := daysSince(analysisTime, customer.LastOrderAt)
		status := classify(days, thresholds)
		result := CustomerResult{
			OriginalIndex:      i,
			CustomerID:         customer.CustomerID,
			Name:               customer.Name,
			DaysSinceLastOrder: days,
			Status:             status,
			IsRetained:         status != "Churned",
			IsChurned:          status == "Churned",
		}
		results[i] = result

		switch status {
		case "Active":
			summary.ActiveCount++
		case "AtRisk":
			summary.AtRiskCount++
		case "Churned":
			summary.ChurnedCount++
		}
	}

	summary.RetentionRate = safePercent(float64(summary.ActiveCount+summary.AtRiskCount), float64(summary.TotalCustomers))
	summary.ChurnRate = safePercent(float64(summary.ChurnedCount), float64(summary.TotalCustomers))

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].DaysSinceLastOrder == results[j].DaysSinceLastOrder {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].DaysSinceLastOrder > results[j].DaysSinceLastOrder
	})

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.AtRiskDays <= 0 {
		t.AtRiskDays = defaultAtRiskDays
	}
	if t.ChurnDays <= 0 {
		t.ChurnDays = defaultChurnDays
	}
	if t.ChurnDays <= t.AtRiskDays {
		t.AtRiskDays = defaultAtRiskDays
		t.ChurnDays = defaultChurnDays
	}
	return t
}

func daysSince(analysisTime time.Time, lastOrderAt time.Time) int {
	if lastOrderAt.IsZero() {
		return defaultChurnDays
	}
	if lastOrderAt.After(analysisTime) {
		return 0
	}
	return int(analysisTime.Sub(lastOrderAt).Hours() / 24)
}

func classify(daysSinceLastOrder int, thresholds Thresholds) string {
	switch {
	case daysSinceLastOrder >= thresholds.ChurnDays:
		return "Churned"
	case daysSinceLastOrder >= thresholds.AtRiskDays:
		return "AtRisk"
	default:
		return "Active"
	}
}

func safePercent(value float64, total float64) float64 {
	if total <= 0 {
		return 0
	}
	return (value / total) * 100
}
