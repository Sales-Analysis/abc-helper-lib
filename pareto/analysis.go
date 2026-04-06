package pareto

import (
	"context"
	"math"
	"sort"
)

const (
	defaultTopValueShare = 80
	defaultTopItemShare  = 20
)

type Thresholds struct {
	TopValueShare float64
	TopItemShare  float64
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
	SKU   string
	Name  string
	Value float64
}

type ItemResult struct {
	OriginalIndex        int
	SKU                  string
	Name                 string
	Value                float64
	ValueShare           float64
	CumulativeValueShare float64
	CumulativeItemShare  float64
	InTopValueSet        bool
	InTopItemSet         bool
}

type Summary struct {
	TotalItems          int
	TotalValue          float64
	TopItemCount        int
	TopItemShare        float64
	TopItemValueShare   float64
	TopValueShareTarget float64
	ParetoPrincipleMet  bool
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	thresholds := input.Thresholds.normalized()
	type indexedItem struct {
		Item
		OriginalIndex int
	}

	items := make([]indexedItem, len(input.Items))
	var totalValue float64
	for i, item := range input.Items {
		items[i] = indexedItem{Item: item, OriginalIndex: i}
		totalValue += item.Value
	}

	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Value == items[j].Value {
			return items[i].OriginalIndex < items[j].OriginalIndex
		}
		return items[i].Value > items[j].Value
	})

	results := make([]ItemResult, len(items))
	var cumulativeValue float64
	topItemCount := topItemsCount(len(items), thresholds.TopItemShare)
	var topItemValue float64

	for i, item := range items {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}

		cumulativeValue += item.Value
		valueShare := safeShare(item.Value, totalValue)
		cumulativeValueShare := safeShare(cumulativeValue, totalValue)
		cumulativeItemShare := safeShare(float64(i+1), float64(len(items)))
		inTopValueSet := itemIncludedInTopValueSet(cumulativeValueShare, i, results, thresholds.TopValueShare)
		inTopItemSet := i < topItemCount
		if inTopItemSet {
			topItemValue += item.Value
		}

		results[i] = ItemResult{
			OriginalIndex:        item.OriginalIndex,
			SKU:                  item.SKU,
			Name:                 item.Name,
			Value:                item.Value,
			ValueShare:           valueShare,
			CumulativeValueShare: cumulativeValueShare,
			CumulativeItemShare:  cumulativeItemShare,
			InTopValueSet:        inTopValueSet,
			InTopItemSet:         inTopItemSet,
		}
	}

	summary := Summary{
		TotalItems:          len(items),
		TotalValue:          totalValue,
		TopItemCount:        topItemCount,
		TopItemShare:        thresholds.TopItemShare,
		TopItemValueShare:   safeShare(topItemValue, totalValue),
		TopValueShareTarget: thresholds.TopValueShare,
		ParetoPrincipleMet:  safeShare(topItemValue, totalValue) >= thresholds.TopValueShare,
	}

	return Output{
		Results: results,
		Summary: summary,
	}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.TopValueShare <= 0 || t.TopValueShare > 100 {
		t.TopValueShare = defaultTopValueShare
	}
	if t.TopItemShare <= 0 || t.TopItemShare > 100 {
		t.TopItemShare = defaultTopItemShare
	}
	return t
}

func safeShare(value float64, total float64) float64 {
	if total <= 0 {
		return 0
	}
	return (value / total) * 100
}

func topItemsCount(totalItems int, topItemShare float64) int {
	if totalItems <= 0 {
		return 0
	}
	count := int(math.Ceil((topItemShare / 100) * float64(totalItems)))
	if count < 1 {
		return 1
	}
	if count > totalItems {
		return totalItems
	}
	return count
}

func itemIncludedInTopValueSet(cumulativeValueShare float64, index int, results []ItemResult, targetShare float64) bool {
	if cumulativeValueShare <= targetShare {
		return true
	}
	if index == 0 {
		return true
	}
	return results[index-1].CumulativeValueShare < targetShare
}
