package hml

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

type Thresholds struct {
	HighMinUnitCost   float64
	MediumMinUnitCost float64
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
	SKU      string
	Name     string
	UnitCost float64
}

type ItemResult struct {
	OriginalIndex int
	SKU           string
	Name          string
	UnitCost      float64
	Group         string
}

type Summary struct {
	TotalItems int
	HCount     int
	MCount     int
	LCount     int
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
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

		group := classify(item.UnitCost, thresholds)
		results[i] = ItemResult{
			OriginalIndex: i,
			SKU:           item.SKU,
			Name:          item.Name,
			UnitCost:      item.UnitCost,
			Group:         group,
		}
		switch group {
		case "H":
			summary.HCount++
		case "M":
			summary.MCount++
		case "L":
			summary.LCount++
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].UnitCost == results[j].UnitCost {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].UnitCost > results[j].UnitCost
	})

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func normalizedThresholds(items []Item, thresholds Thresholds) (Thresholds, error) {
	if thresholds.HighMinUnitCost == 0 && thresholds.MediumMinUnitCost == 0 {
		return deriveThresholds(items), nil
	}
	if err := validation.RequirePositiveFloat("thresholds.HighMinUnitCost", thresholds.HighMinUnitCost); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePositiveFloat("thresholds.MediumMinUnitCost", thresholds.MediumMinUnitCost); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterFloat("thresholds.HighMinUnitCost", thresholds.HighMinUnitCost, "thresholds.MediumMinUnitCost", thresholds.MediumMinUnitCost); err != nil {
		return Thresholds{}, err
	}
	return thresholds, nil
}

func deriveThresholds(items []Item) Thresholds {
	if len(items) == 0 {
		return Thresholds{}
	}

	values := make([]float64, len(items))
	for i, item := range items {
		values[i] = item.UnitCost
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	highPos := ceilDiv(len(values), 3)
	mediumPos := ceilDiv(2*len(values), 3)

	return Thresholds{
		HighMinUnitCost:   values[highPos-1],
		MediumMinUnitCost: values[mediumPos-1],
	}
}

func classify(unitCost float64, thresholds Thresholds) string {
	switch {
	case unitCost >= thresholds.HighMinUnitCost:
		return "H"
	case unitCost >= thresholds.MediumMinUnitCost:
		return "M"
	default:
		return "L"
	}
}

func ceilDiv(numerator int, denominator int) int {
	return (numerator + denominator - 1) / denominator
}
