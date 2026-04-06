package abc

import "context"

const (
	defaultAMaxShare = 80
	defaultBMaxShare = 95
)

// Analyze runs ABC analysis with a stateless API suitable for orchestration layers.
func Analyze(ctx context.Context, input Input) (Output, error) {
	detailed, err := AnalyzeDetailed(ctx, input)
	if err != nil {
		return Output{}, err
	}

	results := make([]ProductResult, len(detailed.Results))
	for i, item := range detailed.Results {
		results[i] = item.ProductResult
	}

	return Output{
		Results:      results,
		TotalRevenue: detailed.TotalRevenue,
	}, nil
}

// AnalyzeDetailed runs ABC analysis and keeps original indexes for composite analyses.
func AnalyzeDetailed(ctx context.Context, input Input) (DetailedOutput, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return DetailedOutput{}, err
		}
	}

	products := append([]Product(nil), input.Products...)
	analysis := New()

	priceTotal := analysis.calculatePriceTotal(products)
	grandTotal := analysis.calculateGrandTotal(priceTotal)
	pairs := analysis.rankProductsByValue(priceTotal)
	costPercentage := analysis.calculateCostPercentage(pairs, grandTotal)
	accumulatedShare := analysis.calculateAccumulatedShare(costPercentage)
	groups := assignGroupWithThresholds(accumulatedShare, input.Thresholds)
	results := analysis.buildResults(products, priceTotal, pairs, costPercentage, accumulatedShare, groups)

	indexedResults := make([]IndexedProductResult, len(results))
	for i, pair := range pairs {
		indexedResults[i] = IndexedProductResult{
			OriginalIndex: pair.index,
			ProductResult: results[i],
		}
	}

	return DetailedOutput{
		Results:      indexedResults,
		TotalRevenue: grandTotal,
	}, nil
}

func (t Thresholds) normalized() Thresholds {
	if t.AMaxShare <= 0 || t.AMaxShare >= 100 {
		t.AMaxShare = defaultAMaxShare
	}
	if t.BMaxShare <= 0 || t.BMaxShare > 100 {
		t.BMaxShare = defaultBMaxShare
	}
	if t.BMaxShare <= t.AMaxShare {
		t.AMaxShare = defaultAMaxShare
		t.BMaxShare = defaultBMaxShare
	}
	return t
}
