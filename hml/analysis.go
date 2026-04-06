package hml

import (
	"context"
	"sort"
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
			UnitCost:      item.UnitCost,
			Group:         classify(item.UnitCost, thresholds),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].UnitCost == results[j].UnitCost {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].UnitCost > results[j].UnitCost
	})

	return Output{Results: results}, nil
}

func normalizedThresholds(items []Item, thresholds Thresholds) Thresholds {
	if thresholds.HighMinUnitCost > thresholds.MediumMinUnitCost {
		return thresholds
	}
	return deriveThresholds(items)
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
