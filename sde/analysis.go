package sde

import (
	"context"
	"sort"
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

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	thresholds := normalizedThresholds(input.Items, input.Thresholds)
	results := make([]ItemResult, len(input.Items))
	for i, item := range input.Items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}

		results[i] = ItemResult{
			OriginalIndex: i,
			SKU:           item.SKU,
			Name:          item.Name,
			LeadTimeDays:  item.LeadTimeDays,
			Group:         classify(item.LeadTimeDays, thresholds),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].LeadTimeDays == results[j].LeadTimeDays {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].LeadTimeDays > results[j].LeadTimeDays
	})

	return Output{Results: results}, nil
}

func normalizedThresholds(items []Item, thresholds Thresholds) Thresholds {
	if thresholds.ScarceMinLeadTime > thresholds.DifficultMinLeadTime {
		return thresholds
	}
	return deriveThresholds(items)
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
