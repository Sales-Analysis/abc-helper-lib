package safetystock

import (
	"context"
	"math"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

const defaultServiceFactor = 1.65

type Input struct {
	Items []Item
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU             string
	Name            string
	DemandStdDev    float64
	LeadTimePeriods float64
	ServiceFactor   float64
}

type ItemResult struct {
	OriginalIndex   int
	SKU             string
	Name            string
	DemandStdDev    float64
	LeadTimePeriods float64
	ServiceFactor   float64
	SafetyStock     float64
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

		serviceFactor := normalizedServiceFactor(item.ServiceFactor)
		results[i] = ItemResult{
			OriginalIndex:   i,
			SKU:             item.SKU,
			Name:            item.Name,
			DemandStdDev:    item.DemandStdDev,
			LeadTimePeriods: item.LeadTimePeriods,
			ServiceFactor:   serviceFactor,
			SafetyStock:     calculateSafetyStock(item.DemandStdDev, item.LeadTimePeriods, serviceFactor),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].SafetyStock == results[j].SafetyStock {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].SafetyStock > results[j].SafetyStock
	})

	return Output{Results: results}, nil
}

func normalizedServiceFactor(serviceFactor float64) float64 {
	if serviceFactor == 0 {
		return defaultServiceFactor
	}
	return serviceFactor
}

func calculateSafetyStock(demandStdDev float64, leadTimePeriods float64, serviceFactor float64) float64 {
	if demandStdDev <= 0 || leadTimePeriods <= 0 || serviceFactor <= 0 {
		return 0
	}
	return serviceFactor * demandStdDev * math.Sqrt(leadTimePeriods)
}

func validateItem(item Item) error {
	if err := validation.RequireNonNegativeFloat("items[].DemandStdDev", item.DemandStdDev); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].LeadTimePeriods", item.LeadTimePeriods); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].ServiceFactor", item.ServiceFactor); err != nil {
		return err
	}
	return nil
}
