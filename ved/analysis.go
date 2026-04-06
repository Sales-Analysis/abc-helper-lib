package ved

import (
	"context"
	"sort"
)

const (
	defaultVitalMinScore     = 70
	defaultEssentialMinScore = 40
)

type Thresholds struct {
	VitalMinScore     float64
	EssentialMinScore float64
}

type Input struct {
	Items      []Item
	Thresholds Thresholds
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU              string
	Name             string
	CriticalityScore float64
}

type ItemResult struct {
	OriginalIndex    int
	SKU              string
	Name             string
	CriticalityScore float64
	Group            string
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

		results[i] = ItemResult{
			OriginalIndex:    i,
			SKU:              item.SKU,
			Name:             item.Name,
			CriticalityScore: item.CriticalityScore,
			Group:            classify(item.CriticalityScore, thresholds),
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].CriticalityScore == results[j].CriticalityScore {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].CriticalityScore > results[j].CriticalityScore
	})

	return Output{Results: results}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.VitalMinScore <= 0 || t.VitalMinScore > 100 {
		t.VitalMinScore = defaultVitalMinScore
	}
	if t.EssentialMinScore <= 0 || t.EssentialMinScore > 100 {
		t.EssentialMinScore = defaultEssentialMinScore
	}
	if t.VitalMinScore <= t.EssentialMinScore {
		t.VitalMinScore = defaultVitalMinScore
		t.EssentialMinScore = defaultEssentialMinScore
	}
	return t
}

func classify(score float64, thresholds Thresholds) string {
	switch {
	case score >= thresholds.VitalMinScore:
		return "V"
	case score >= thresholds.EssentialMinScore:
		return "E"
	default:
		return "D"
	}
}
