package gmcontribution

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

type Input struct {
	Items []Item
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU          string
	Name         string
	Revenue      float64
	COGS         float64
	VariableCost float64
	FixedCost    float64
}

type ItemResult struct {
	OriginalIndex      int
	SKU                string
	Name               string
	Revenue            float64
	COGS               float64
	VariableCost       float64
	FixedCost          float64
	GrossMargin        float64
	GrossMarginRate    float64
	ContributionMargin float64
	ContributionRate   float64
	NetContribution    float64
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	results := make([]ItemResult, len(input.Items))
	for i, item := range input.Items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}
		if err := validateItem(item); err != nil {
			return Output{}, err
		}

		grossMargin := item.Revenue - item.COGS
		contributionMargin := item.Revenue - item.VariableCost
		results[i] = ItemResult{
			OriginalIndex:      i,
			SKU:                item.SKU,
			Name:               item.Name,
			Revenue:            item.Revenue,
			COGS:               item.COGS,
			VariableCost:       item.VariableCost,
			FixedCost:          item.FixedCost,
			GrossMargin:        grossMargin,
			GrossMarginRate:    safeRate(grossMargin, item.Revenue),
			ContributionMargin: contributionMargin,
			ContributionRate:   safeRate(contributionMargin, item.Revenue),
			NetContribution:    contributionMargin - item.FixedCost,
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].ContributionMargin == results[j].ContributionMargin {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].ContributionMargin > results[j].ContributionMargin
	})

	return Output{Results: results}, nil
}

func validateItem(item Item) error {
	if err := validation.RequireNonNegativeFloat("items[].Revenue", item.Revenue); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].COGS", item.COGS); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].VariableCost", item.VariableCost); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].FixedCost", item.FixedCost); err != nil {
		return err
	}
	return nil
}

func safeRate(value float64, base float64) float64 {
	if base <= 0 {
		return 0
	}
	return (value / base) * 100
}
