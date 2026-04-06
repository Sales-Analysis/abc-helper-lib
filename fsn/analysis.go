package fsn

import (
	"context"
	"sort"
)

const (
	defaultFastMinActivityRate = 70
	defaultSlowMinActivityRate = 1
)

type Thresholds struct {
	FastMinActivityRate float64
	SlowMinActivityRate float64
}

type Input struct {
	Items      []Item
	Thresholds Thresholds
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU       string
	Name      string
	Movements []float64
}

type ItemResult struct {
	OriginalIndex      int
	SKU                string
	Name               string
	Periods            int
	ActivePeriods      int
	ActivityRate       float64
	TotalMovement      float64
	AverageMovement    float64
	LastMovementPeriod int
	Group              string
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

		total, activePeriods, lastMovementPeriod := movementStats(item.Movements)
		periods := len(item.Movements)
		activityRate := calculateActivityRate(activePeriods, periods)
		averageMovement := calculateAverageMovement(total, periods)

		results[i] = ItemResult{
			OriginalIndex:      i,
			SKU:                item.SKU,
			Name:               item.Name,
			Periods:            periods,
			ActivePeriods:      activePeriods,
			ActivityRate:       activityRate,
			TotalMovement:      total,
			AverageMovement:    averageMovement,
			LastMovementPeriod: lastMovementPeriod,
			Group:              classify(activityRate, activePeriods, thresholds),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].ActivityRate == results[j].ActivityRate {
			if results[i].AverageMovement == results[j].AverageMovement {
				return results[i].OriginalIndex < results[j].OriginalIndex
			}
			return results[i].AverageMovement > results[j].AverageMovement
		}
		return results[i].ActivityRate > results[j].ActivityRate
	})

	return Output{Results: results}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.FastMinActivityRate <= 0 || t.FastMinActivityRate > 100 {
		t.FastMinActivityRate = defaultFastMinActivityRate
	}
	if t.SlowMinActivityRate < 0 || t.SlowMinActivityRate > 100 {
		t.SlowMinActivityRate = defaultSlowMinActivityRate
	}
	if t.FastMinActivityRate <= t.SlowMinActivityRate {
		t.FastMinActivityRate = defaultFastMinActivityRate
		t.SlowMinActivityRate = defaultSlowMinActivityRate
	}
	return t
}

func movementStats(movements []float64) (float64, int, int) {
	var total float64
	var activePeriods int
	var lastMovementPeriod int

	for i, movement := range movements {
		total += movement
		if movement > 0 {
			activePeriods++
			lastMovementPeriod = i + 1
		}
	}

	return total, activePeriods, lastMovementPeriod
}

func calculateActivityRate(activePeriods int, totalPeriods int) float64 {
	if totalPeriods == 0 {
		return 0
	}
	return (float64(activePeriods) / float64(totalPeriods)) * 100
}

func calculateAverageMovement(total float64, totalPeriods int) float64 {
	if totalPeriods == 0 {
		return 0
	}
	return total / float64(totalPeriods)
}

func classify(activityRate float64, activePeriods int, thresholds Thresholds) string {
	if activePeriods == 0 {
		return "N"
	}
	if activityRate >= thresholds.FastMinActivityRate {
		return "F"
	}
	if activityRate >= thresholds.SlowMinActivityRate {
		return "S"
	}
	return "N"
}
