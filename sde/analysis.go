package sde

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

type Thresholds struct {
	ScarceMinLeadTime    int
	DifficultMinLeadTime int
}

type Input struct {
	Items      []Item
	Thresholds Thresholds
}

type Output struct {
	Results []ItemResult
	Summary Summary
}

type Item struct {
	SKU          string
	Name         string
	LeadTimeDays int
}

type ItemResult struct {
	OriginalIndex int
	SKU           string
	Name          string
	LeadTimeDays  int
	Group         string
}

type Summary struct {
	TotalItems int
	SCount     int
	DCount     int
	ECount     int
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	for _, item := range input.Items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}
		if err := validateItem(item); err != nil {
			return Output{}, err
		}
	}

	thresholds, err := normalizedThresholds(input.Items, input.Thresholds)
	if err != nil {
		return Output{}, err
	}
	results := make([]ItemResult, len(input.Items))
	summary := Summary{TotalItems: len(input.Items)}
	for i, item := range input.Items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}

		group := classify(item.LeadTimeDays, thresholds)
		results[i] = ItemResult{
			OriginalIndex: i,
			SKU:           item.SKU,
			Name:          item.Name,
			LeadTimeDays:  item.LeadTimeDays,
			Group:         group,
		}
		switch group {
		case "S":
			summary.SCount++
		case "D":
			summary.DCount++
		case "E":
			summary.ECount++
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].LeadTimeDays == results[j].LeadTimeDays {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].LeadTimeDays > results[j].LeadTimeDays
	})

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func validateItem(item Item) error {
	if err := validation.RequireNonNegativeInt("items[].LeadTimeDays", item.LeadTimeDays); err != nil {
		return err
	}
	return nil
}

func normalizedThresholds(items []Item, thresholds Thresholds) (Thresholds, error) {
	if thresholds.ScarceMinLeadTime == 0 && thresholds.DifficultMinLeadTime == 0 {
		return deriveThresholds(items), nil
	}
	if err := validation.RequirePositiveInt("thresholds.ScarceMinLeadTime", thresholds.ScarceMinLeadTime); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePositiveInt("thresholds.DifficultMinLeadTime", thresholds.DifficultMinLeadTime); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterInt("thresholds.ScarceMinLeadTime", thresholds.ScarceMinLeadTime, "thresholds.DifficultMinLeadTime", thresholds.DifficultMinLeadTime); err != nil {
		return Thresholds{}, err
	}
	return thresholds, nil
}

func deriveThresholds(items []Item) Thresholds {
	if len(items) == 0 {
		return Thresholds{}
	}

	values := make([]int, len(items))
	for i, item := range items {
		values[i] = item.LeadTimeDays
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	scarcePos := ceilDiv(len(values), 3)
	difficultPos := ceilDiv(2*len(values), 3)

	return Thresholds{
		ScarceMinLeadTime:    values[scarcePos-1],
		DifficultMinLeadTime: values[difficultPos-1],
	}
}

func classify(leadTimeDays int, thresholds Thresholds) string {
	switch {
	case leadTimeDays >= thresholds.ScarceMinLeadTime:
		return "S"
	case leadTimeDays >= thresholds.DifficultMinLeadTime:
		return "D"
	default:
		return "E"
	}
}

func ceilDiv(numerator int, denominator int) int {
	return (numerator + denominator - 1) / denominator
}
