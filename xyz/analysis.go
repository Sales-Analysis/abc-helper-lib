package xyz

import (
	"context"
	"math"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

const (
	defaultXMaxCV = 10
	defaultYMaxCV = 25
)

type Thresholds struct {
	XMaxCV float64
	YMaxCV float64
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
	SKU     string
	Name    string
	Demands []float64
}

type ItemResult struct {
	OriginalIndex          int
	SKU                    string
	Name                   string
	Periods                int
	AverageDemand          float64
	StandardDeviation      float64
	CoefficientOfVariation float64
	Group                  string
}

type Summary struct {
	TotalItems int
	XCount     int
	YCount     int
	ZCount     int
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	thresholds, err := input.Thresholds.normalized()
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
		if err := validateItem(item); err != nil {
			return Output{}, err
		}

		mean, stddev := demandStats(item.Demands)
		cv := coefficientOfVariation(mean, stddev)

		group := classify(cv, mean, len(item.Demands), thresholds)
		results[i] = ItemResult{
			OriginalIndex:          i,
			SKU:                    item.SKU,
			Name:                   item.Name,
			Periods:                len(item.Demands),
			AverageDemand:          mean,
			StandardDeviation:      stddev,
			CoefficientOfVariation: cv,
			Group:                  group,
		}
		switch group {
		case "X":
			summary.XCount++
		case "Y":
			summary.YCount++
		case "Z":
			summary.ZCount++
		}
	}

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func validateItem(item Item) error {
	for _, demand := range item.Demands {
		if err := validation.RequireNonNegativeFloat("items[].Demands[]", demand); err != nil {
			return err
		}
	}
	return nil
}

func (t Thresholds) normalized() (Thresholds, error) {
	if t.XMaxCV == 0 {
		t.XMaxCV = defaultXMaxCV
	}
	if t.YMaxCV == 0 {
		t.YMaxCV = defaultYMaxCV
	}
	if err := validation.RequirePositiveFloat("thresholds.XMaxCV", t.XMaxCV); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePositiveFloat("thresholds.YMaxCV", t.YMaxCV); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterFloat("thresholds.YMaxCV", t.YMaxCV, "thresholds.XMaxCV", t.XMaxCV); err != nil {
		return Thresholds{}, err
	}
	return t, nil
}

func demandStats(demands []float64) (float64, float64) {
	if len(demands) == 0 {
		return 0, 0
	}

	var sum float64
	for _, value := range demands {
		sum += value
	}
	mean := sum / float64(len(demands))

	var variance float64
	for _, value := range demands {
		diff := value - mean
		variance += diff * diff
	}
	variance /= float64(len(demands))

	return mean, math.Sqrt(variance)
}

func coefficientOfVariation(mean float64, stddev float64) float64 {
	if mean == 0 {
		return 0
	}
	return (stddev / mean) * 100
}

func classify(cv float64, mean float64, periods int, thresholds Thresholds) string {
	if periods == 0 || mean == 0 {
		return "Z"
	}
	if cv <= thresholds.XMaxCV {
		return "X"
	}
	if cv <= thresholds.YMaxCV {
		return "Y"
	}
	return "Z"
}
