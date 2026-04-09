package clv

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

type Input struct {
	Customers []Customer
}

type Output struct {
	Results []CustomerResult
}

type Customer struct {
	CustomerID      string
	Name            string
	Revenue         float64
	Orders          int
	PeriodsObserved float64
	GrossMarginRate float64
	RetentionRate   float64
	DiscountRate    float64
	AcquisitionCost float64
}

type CustomerResult struct {
	OriginalIndex     int
	CustomerID        string
	Name              string
	Revenue           float64
	Orders            int
	PeriodsObserved   float64
	AverageOrderValue float64
	PurchaseFrequency float64
	GrossMarginRate   float64
	RetentionRate     float64
	DiscountRate      float64
	AcquisitionCost   float64
	PredictedCLV      float64
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	results := make([]CustomerResult, len(input.Customers))
	for i, customer := range input.Customers {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}
		if err := validateCustomer(customer); err != nil {
			return Output{}, err
		}

		grossMarginRate := normalizeRatio(customer.GrossMarginRate)
		retentionRate := normalizeRatio(customer.RetentionRate)
		discountRate := normalizeRatio(customer.DiscountRate)
		aov := averageOrderValue(customer.Revenue, customer.Orders)
		purchaseFrequency := averagePurchaseFrequency(customer.Orders, customer.PeriodsObserved)

		results[i] = CustomerResult{
			OriginalIndex:     i,
			CustomerID:        customer.CustomerID,
			Name:              customer.Name,
			Revenue:           customer.Revenue,
			Orders:            customer.Orders,
			PeriodsObserved:   customer.PeriodsObserved,
			AverageOrderValue: aov,
			PurchaseFrequency: purchaseFrequency,
			GrossMarginRate:   grossMarginRate,
			RetentionRate:     retentionRate,
			DiscountRate:      discountRate,
			AcquisitionCost:   customer.AcquisitionCost,
			PredictedCLV:      calculateCLV(aov, purchaseFrequency, grossMarginRate, retentionRate, discountRate, customer.AcquisitionCost),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].PredictedCLV == results[j].PredictedCLV {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].PredictedCLV > results[j].PredictedCLV
	})

	return Output{Results: results}, nil
}

func validateCustomer(customer Customer) error {
	if err := validation.RequireNonNegativeFloat("customers[].Revenue", customer.Revenue); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeInt("customers[].Orders", customer.Orders); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("customers[].PeriodsObserved", customer.PeriodsObserved); err != nil {
		return err
	}
	if err := validation.RequireRatio("customers[].GrossMarginRate", customer.GrossMarginRate); err != nil {
		return err
	}
	if err := validation.RequireRatio("customers[].RetentionRate", customer.RetentionRate); err != nil {
		return err
	}
	if err := validation.RequireRatio("customers[].DiscountRate", customer.DiscountRate); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("customers[].AcquisitionCost", customer.AcquisitionCost); err != nil {
		return err
	}
	return nil
}

func normalizeRatio(value float64) float64 {
	if value <= 0 {
		return 0
	}
	if value > 1 && value <= 100 {
		return value / 100
	}
	return value
}

func averageOrderValue(revenue float64, orders int) float64 {
	if revenue <= 0 || orders <= 0 {
		return 0
	}
	return revenue / float64(orders)
}

func averagePurchaseFrequency(orders int, periodsObserved float64) float64 {
	if orders <= 0 || periodsObserved <= 0 {
		return 0
	}
	return float64(orders) / periodsObserved
}

func calculateCLV(aov float64, purchaseFrequency float64, grossMarginRate float64, retentionRate float64, discountRate float64, acquisitionCost float64) float64 {
	if aov <= 0 || purchaseFrequency <= 0 || grossMarginRate <= 0 || retentionRate <= 0 {
		return 0 - acquisitionCost
	}

	denominator := 1 + discountRate - retentionRate
	if denominator <= 0 {
		return 0 - acquisitionCost
	}

	marginAdjustedPeriodValue := aov * purchaseFrequency * grossMarginRate
	return (marginAdjustedPeriodValue * (retentionRate / denominator)) - acquisitionCost
}
