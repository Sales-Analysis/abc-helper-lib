package xyz

import (
	"context"
	"math"
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

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	thresholds := input.Thresholds.normalized()
	results := make([]ItemResult, len(input.Items))

	for i, item := range input.Items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}

		mean, stddev := demandStats(item.Demands)
		cv := coefficientOfVariation(mean, stddev)

		results[i] = ItemResult{
			OriginalIndex:          i,
			SKU:                    item.SKU,
			Name:                   item.Name,
			Periods:                len(item.Demands),
			AverageDemand:          mean,
			StandardDeviation:      stddev,
			CoefficientOfVariation: cv,
			Group:                  classify(cv, mean, len(item.Demands), thresholds),
		}
	}

	return Output{Results: results}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.XMaxCV <= 0 {
		t.XMaxCV = defaultXMaxCV
	}
	if t.YMaxCV <= 0 {
		t.YMaxCV = defaultYMaxCV
	}
	if t.YMaxCV <= t.XMaxCV {
		t.XMaxCV = defaultXMaxCV
		t.YMaxCV = defaultYMaxCV
	}
	return t
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
