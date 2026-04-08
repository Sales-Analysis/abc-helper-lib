package reorderpoint

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
	SKU                    string
	Name                   string
	AverageDemandPerPeriod float64
	LeadTimePeriods        float64
	SafetyStock            float64
}

type ItemResult struct {
	OriginalIndex          int
	SKU                    string
	Name                   string
	AverageDemandPerPeriod float64
	LeadTimePeriods        float64
	SafetyStock            float64
	LeadTimeDemand         float64
	ReorderPoint           float64
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

		leadTimeDemand := calculateLeadTimeDemand(item.AverageDemandPerPeriod, item.LeadTimePeriods)
		results[i] = ItemResult{
			OriginalIndex:          i,
			SKU:                    item.SKU,
			Name:                   item.Name,
			AverageDemandPerPeriod: item.AverageDemandPerPeriod,
			LeadTimePeriods:        item.LeadTimePeriods,
			SafetyStock:            item.SafetyStock,
			LeadTimeDemand:         leadTimeDemand,
			ReorderPoint:           leadTimeDemand + item.SafetyStock,
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].ReorderPoint == results[j].ReorderPoint {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].ReorderPoint > results[j].ReorderPoint
	})

	return Output{Results: results}, nil
}

func validateItem(item Item) error {
	if err := validation.RequireNonNegativeFloat("items[].AverageDemandPerPeriod", item.AverageDemandPerPeriod); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].LeadTimePeriods", item.LeadTimePeriods); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].SafetyStock", item.SafetyStock); err != nil {
		return err
	}
	return nil
}

func calculateLeadTimeDemand(averageDemandPerPeriod float64, leadTimePeriods float64) float64 {
	if averageDemandPerPeriod <= 0 || leadTimePeriods <= 0 {
		return 0
	}
	return averageDemandPerPeriod * leadTimePeriods
}
