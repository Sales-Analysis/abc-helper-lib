package abc

import (
	"context"
	"sort"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
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
	Items []Product
	// Deprecated: use Items for new integrations.
	Products   []Product
	Thresholds Thresholds
}

type Output struct {
	Results      []ProductResult
	TotalRevenue float64
	Summary      Summary
}

type DetailedOutput struct {
	Results      []IndexedProductResult
	TotalRevenue float64
	Summary      Summary
}

// Product is the input struct for the analysis.
type Product struct {
	SKU      string
	Name     string
	Quantity int
	Price    float64
}

// Item is the preferred inventory-facing alias for Product.
type Item = Product

type ProductResult struct {
	OriginalIndex    int
	SKU              string
	Name             string
	Quantity         int
	PriceUnit        float64
	PriceTotal       float64
	ShareTotal       float64
	ShareAccumulated float64
	Group            string
}

// ItemResult is the preferred inventory-facing alias for ProductResult.
type ItemResult = ProductResult

type IndexedProductResult = ProductResult

// IndexedItemResult is the preferred inventory-facing alias for IndexedProductResult.
type IndexedItemResult = IndexedProductResult

type Summary struct {
	TotalItems    int
	TotalRevenue  float64
	ACount        int
	BCount        int
	CCount        int
	SortedByValue bool
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
	output, err := Analyze(context.Background(), Input{Items: products})
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
	copy(results, detailed.Results)

	return Output{
		Results:      results,
		TotalRevenue: detailed.TotalRevenue,
		Summary:      detailed.Summary,
	}, nil
}

// AnalyzeDetailed runs ABC analysis and keeps original indexes for composite analyses.
func AnalyzeDetailed(ctx context.Context, input Input) (DetailedOutput, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return DetailedOutput{}, err
		}
	}

	thresholds, err := input.Thresholds.normalized()
	if err != nil {
		return DetailedOutput{}, err
	}

	products := input.normalizedItems()
	for _, product := range products {
		if err := validateProduct(product); err != nil {
			return DetailedOutput{}, err
		}
	}
	analysis := New()

	priceTotal := analysis.calculatePriceTotal(products)
	grandTotal := analysis.calculateGrandTotal(priceTotal)
	pairs := analysis.rankProductsByValue(priceTotal)
	costPercentage := analysis.calculateCostPercentage(pairs, grandTotal)
	accumulatedShare := analysis.calculateAccumulatedShare(costPercentage)
	groups := assignGroupWithThresholds(accumulatedShare, thresholds)
	results := analysis.buildResults(products, priceTotal, pairs, costPercentage, accumulatedShare, groups)
	summary := summarize(results, grandTotal)

	return DetailedOutput{
		Results:      results,
		TotalRevenue: grandTotal,
		Summary:      summary,
	}, nil
}

func validateProduct(product Product) error {
	if err := validation.RequireNonNegativeInt("items[].Quantity", product.Quantity); err != nil {
		return err
	}
	if err := validation.RequireNonNegativeFloat("items[].Price", product.Price); err != nil {
		return err
	}
	return nil
}

func (in Input) normalizedItems() []Product {
	if in.Items != nil {
		return append([]Product(nil), in.Items...)
	}
	return append([]Product(nil), in.Products...)
}

func (t Thresholds) normalized() (Thresholds, error) {
	if t.AMaxShare == 0 {
		t.AMaxShare = defaultAMaxShare
	}
	if t.BMaxShare == 0 {
		t.BMaxShare = defaultBMaxShare
	}
	if err := validation.RequirePercent("thresholds.AMaxShare", t.AMaxShare); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequirePercent("thresholds.BMaxShare", t.BMaxShare); err != nil {
		return Thresholds{}, err
	}
	if err := validation.RequireGreaterFloat("thresholds.BMaxShare", t.BMaxShare, "thresholds.AMaxShare", t.AMaxShare); err != nil {
		return Thresholds{}, err
	}
	return t, nil
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
	if grandTotal <= 0 {
		for range pairs {
			costPercentage = append(costPercentage, 0)
		}
		return costPercentage
	}
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
	thresholds, err := (Thresholds{}).normalized()
	if err != nil {
		return nil
	}
	return assignGroupWithThresholds(accumulatedShare, thresholds)
}

func assignGroupWithThresholds(accumulatedShare []float64, thresholds Thresholds) []string {
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
			OriginalIndex:    idx,
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

func summarize(results []ProductResult, totalRevenue float64) Summary {
	summary := Summary{
		TotalItems:    len(results),
		TotalRevenue:  totalRevenue,
		SortedByValue: true,
	}
	for _, result := range results {
		switch result.Group {
		case "A":
			summary.ACount++
		case "B":
			summary.BCount++
		case "C":
			summary.CCount++
		}
	}
	return summary
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
