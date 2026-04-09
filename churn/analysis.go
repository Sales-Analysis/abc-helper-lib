package churn

import (
	"context"
	"sort"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
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
	thresholds, err := input.Thresholds.normalized()
	if err != nil {
		return Output{}, err
	}

	results := make([]CustomerResult, len(input.Customers))
	summary := Summary{TotalCustomers: len(input.Customers)}
	for i, customer := range input.Customers {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}
		if err := validateCustomer(customer, analysisTime); err != nil {
			return Output{}, err
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

func validateCustomer(customer Customer, analysisTime time.Time) error {
	if !customer.LastOrderAt.IsZero() && customer.LastOrderAt.After(analysisTime) {
		return validation.Invalidf("customers[].LastOrderAt must not be in the future")
	}
	return nil
}

func (t Thresholds) normalized() (Thresholds, error) {
	if t.AtRiskDays == 0 {
		t.AtRiskDays = defaultAtRiskDays
	}
	if t.ChurnDays == 0 {
		t.ChurnDays = defaultChurnDays
	}
	if err := validation.RequirePositiveInt("thresholds.AtRiskDays", t.AtRiskDays); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePositiveInt("thresholds.ChurnDays", t.ChurnDays); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterInt("thresholds.ChurnDays", t.ChurnDays, "thresholds.AtRiskDays", t.AtRiskDays); err != nil {
		return Thresholds{}, err
	}
	return t, nil
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
