package eoq

import (
	"context"
	"math"
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
	AnnualDemand float64
	OrderingCost float64
	HoldingCost  float64
}

type ItemResult struct {
	OriginalIndex   int
	SKU             string
	Name            string
	AnnualDemand    float64
	OrderingCost    float64
	HoldingCost     float64
	OptimalQuantity float64
	OrdersPerYear   float64
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

		optimalQuantity := calculateEOQ(item.AnnualDemand, item.OrderingCost, item.HoldingCost)
		results[i] = ItemResult{
			OriginalIndex:   i,
			SKU:             item.SKU,
			Name:            item.Name,
			AnnualDemand:    item.AnnualDemand,
			OrderingCost:    item.OrderingCost,
			HoldingCost:     item.HoldingCost,
			OptimalQuantity: optimalQuantity,
			OrdersPerYear:   calculateOrdersPerYear(item.AnnualDemand, optimalQuantity),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].OptimalQuantity == results[j].OptimalQuantity {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].OptimalQuantity > results[j].OptimalQuantity
	})

	return Output{Results: results}, nil
}

func validateItem(item Item) error {
	if err := validation.RequireNonNegativeFloat("items[].AnnualDemand", item.AnnualDemand); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].OrderingCost", item.OrderingCost); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].HoldingCost", item.HoldingCost); err != nil {
		return err
	}
	return nil
}

func calculateEOQ(annualDemand float64, orderingCost float64, holdingCost float64) float64 {
	if annualDemand <= 0 || orderingCost <= 0 || holdingCost <= 0 {
		return 0
	}
	return math.Sqrt((2 * annualDemand * orderingCost) / holdingCost)
}

func calculateOrdersPerYear(annualDemand float64, optimalQuantity float64) float64 {
	if annualDemand <= 0 || optimalQuantity <= 0 {
		return 0
	}
	return annualDemand / optimalQuantity
}
