package servicelevel

import (
	"context"
	"sort"
)

type Input struct {
	Items []Item
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU            string
	Name           string
	DemandedUnits  float64
	FulfilledUnits float64
	TotalCycles    int
	StockoutCycles int
}

type ItemResult struct {
	OriginalIndex     int
	SKU               string
	Name              string
	DemandedUnits     float64
	FulfilledUnits    float64
	UnfulfilledUnits  float64
	FillRate          float64
	TotalCycles       int
	StockoutCycles    int
	CycleServiceLevel float64
	StockoutRate      float64
	ServiceLevel      float64
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

		fillRate := safeUnitRate(item.FulfilledUnits, item.DemandedUnits)
		cycleServiceLevel := calculateCycleServiceLevel(item.TotalCycles, item.StockoutCycles)
		results[i] = ItemResult{
			OriginalIndex:     i,
			SKU:               item.SKU,
			Name:              item.Name,
			DemandedUnits:     item.DemandedUnits,
			FulfilledUnits:    item.FulfilledUnits,
			UnfulfilledUnits:  maxFloat(item.DemandedUnits-item.FulfilledUnits, 0),
			FillRate:          fillRate,
			TotalCycles:       item.TotalCycles,
			StockoutCycles:    item.StockoutCycles,
			CycleServiceLevel: cycleServiceLevel,
			StockoutRate:      calculateStockoutRate(item.TotalCycles, item.StockoutCycles),
			ServiceLevel:      selectServiceLevel(fillRate, cycleServiceLevel, item.TotalCycles),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].ServiceLevel == results[j].ServiceLevel {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].ServiceLevel > results[j].ServiceLevel
	})

	return Output{Results: results}, nil
}

func safeUnitRate(value float64, total float64) float64 {
	if total <= 0 {
		return 0
	}
	return (value / total) * 100
}

func calculateCycleServiceLevel(totalCycles int, stockoutCycles int) float64 {
	if totalCycles <= 0 {
		return 0
	}
	return (float64(totalCycles-stockoutCycles) / float64(totalCycles)) * 100
}

func calculateStockoutRate(totalCycles int, stockoutCycles int) float64 {
	if totalCycles <= 0 {
		return 0
	}
	return (float64(stockoutCycles) / float64(totalCycles)) * 100
}

func selectServiceLevel(fillRate float64, cycleServiceLevel float64, totalCycles int) float64 {
	if totalCycles > 0 {
		return cycleServiceLevel
	}
	return fillRate
}

func maxFloat(left float64, right float64) float64 {
	if left > right {
		return left
	}
	return right
}
