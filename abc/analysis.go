package abc

import (
	"context"
	"sort"
)

const (
	defaultAMaxShare = 80
	defaultBMaxShare = 95
)

// Deprecated: use Analyze or AnalyzeDetailed instead of the stateful ABC type.
type ABC struct {
	Result []ProductResult
}

type Thresholds struct {
	AMaxShare float64
	BMaxShare float64
}

type Input struct {
	Products   []Product
	Thresholds Thresholds
}

type Output struct {
	Results      []ProductResult
	TotalRevenue float64
}

type DetailedOutput struct {
	Results      []IndexedProductResult
	TotalRevenue float64
}

// Product is the input struct for the analysis.
type Product struct {
	SKU      string
	Name     string
	Quantity int
	Price    float64
}

type ProductResult struct {
	SKU              string
	Name             string
	Quantity         int
	PriceUnit        float64
	PriceTotal       float64
	ShareTotal       float64
	ShareAccumulated float64
	Group            string
}

type IndexedProductResult struct {
	OriginalIndex int
	ProductResult
}

type pair struct {
	value float64
	index int
}

type byValue []pair

func (a byValue) Len() int           { return len(a) }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValue) Less(i, j int) bool { return a[i].value > a[j].value }

// Deprecated: use Analyze or AnalyzeDetailed instead of New and Calculate.
func New() *ABC {
	return &ABC{}
}

// Deprecated: use Analyze or AnalyzeDetailed instead of the stateful Calculate method.
func (a *ABC) Calculate(products []Product) {
	output, err := Analyze(context.Background(), Input{Products: products})
	if err != nil {
		a.Result = nil
		return
	}
	a.Result = output.Results
}

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

func (a *ABC) calculatePriceTotal(products []Product) []float64 {
	total := make([]float64, 0, len(products))
	for _, value := range products {
		total = append(total, float64(value.Quantity)*value.Price)
	}
	return total
}

func (a *ABC) calculateGrandTotal(totals []float64) float64 {
	var grandTotal float64
	for _, total := range totals {
		grandTotal += total
	}
	return grandTotal
}

func (a *ABC) calculateCostPercentage(pairs byValue, grandTotal float64) []float64 {
	costPercentage := make([]float64, 0, len(pairs))
	for _, value := range pairs {
		v := (value.value / grandTotal) * 100
		costPercentage = append(costPercentage, v)
	}
	return costPercentage
}

func (a *ABC) calculateAccumulatedShare(costPercentage []float64) []float64 {
	accumulatedShare := make([]float64, 0, len(costPercentage))
	as := 0.0
	for _, value := range costPercentage {
		as += value
		accumulatedShare = append(accumulatedShare, as)
	}
	return accumulatedShare
}

func (a *ABC) assignGroup(accumulatedShare []float64) []string {
	return assignGroupWithThresholds(accumulatedShare, Thresholds{})
}

func assignGroupWithThresholds(accumulatedShare []float64, thresholds Thresholds) []string {
	thresholds = thresholds.normalized()
	groups := make([]string, 0, len(accumulatedShare))
	for _, value := range accumulatedShare {
		if value <= thresholds.AMaxShare {
			groups = append(groups, "A")
		} else if value <= thresholds.BMaxShare {
			groups = append(groups, "B")
		} else {
			groups = append(groups, "C")
		}
	}
	return groups
}

// buildResults compiles the final list of ProductResult from calculation data.
func (a *ABC) buildResults(
	products []Product,
	priceTotal []float64,
	pairs byValue,
	costPercentage []float64,
	accumulatedShare []float64,
	groups []string,
) []ProductResult {
	results := make([]ProductResult, len(products))
	for i, pair := range pairs {
		idx := pair.index

		results[i] = ProductResult{
			SKU:              products[idx].SKU,
			Name:             products[idx].Name,
			Quantity:         products[idx].Quantity,
			PriceUnit:        products[idx].Price,
			PriceTotal:       priceTotal[idx],
			ShareTotal:       costPercentage[i],
			ShareAccumulated: accumulatedShare[i],
			Group:            groups[i],
		}
	}
	return results
}

func (a *ABC) rankProductsByValue(priceTotal []float64) byValue {
	indexes := createIndexesSlice(len(priceTotal))
	return sortIndexByValue(indexes, priceTotal)
}

func createIndexesSlice(lenSlice int) []int {
	indexes := make([]int, 0, lenSlice)
	for i := 0; i < lenSlice; i++ {
		indexes = append(indexes, i)
	}
	return indexes
}

func sortIndexByValue(indexes []int, values []float64) byValue {
	pairs := make(byValue, len(values))
	for i := 0; i < len(values); i++ {
		pairs[i] = pair{values[i], indexes[i]}
	}
	sort.Sort(pairs)
	return pairs
}
