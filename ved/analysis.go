package ved

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
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
	Summary Summary
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

type Summary struct {
	TotalItems int
	VCount     int
	ECount     int
	DCount     int
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

		group := classify(item.CriticalityScore, thresholds)
		results[i] = ItemResult{
			OriginalIndex:    i,
			SKU:              item.SKU,
			Name:             item.Name,
			CriticalityScore: item.CriticalityScore,
			Group:            group,
		}
		switch group {
		case "V":
			summary.VCount++
		case "E":
			summary.ECount++
		case "D":
			summary.DCount++
		}
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].CriticalityScore == results[j].CriticalityScore {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].CriticalityScore > results[j].CriticalityScore
	})

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func (t Thresholds) normalized() (Thresholds, error) {
	if t.VitalMinScore == 0 {
		t.VitalMinScore = defaultVitalMinScore
	}
	if t.EssentialMinScore == 0 {
		t.EssentialMinScore = defaultEssentialMinScore
	}
	if err := validation.RequirePercent("thresholds.VitalMinScore", t.VitalMinScore); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePercent("thresholds.EssentialMinScore", t.EssentialMinScore); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterFloat("thresholds.VitalMinScore", t.VitalMinScore, "thresholds.EssentialMinScore", t.EssentialMinScore); err != nil {
		return Thresholds{}, err
	}
	return t, nil
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
