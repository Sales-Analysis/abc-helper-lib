package abc

import (
	"sort"
)

// Constructor for the ABC struct. Creates and returns a new empty ABC object.
func New() *ABC {
	return &ABC{}
}

// Main method that runs the entire ABC analysis workflow
func (a *ABC) Calculate(products []Product) {
	priceTotal := a.calculatePriceTotal(products)
	grandTotal := a.calculateGrandTotal(priceTotal)
	pairs := a.rankProductsByValue(priceTotal)
	costPercentage := a.calculateCostPercentage(pairs, grandTotal)
	accumulatedShare := a.calculateAccumulatedShare(costPercentage)
	groups := a.assignGroup(accumulatedShare)

	a.Result = a.buildResults(products, priceTotal, pairs, costPercentage, accumulatedShare, groups)
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

// Generates a list of product indexes and sorts them by their revenue in descending order.
func (a *ABC) rankProductsByValue(priceTotal []float64) byValue {
	indexes := createIndexesSlice(len(priceTotal))
	pairs := sortIndexByValue(indexes, priceTotal)
	return pairs
}

// Creates a slice of sequential indexes [0, 1, 2, ...].
func createIndexesSlice(lenSlice int) []int {
	indexes := []int{}
	for i := 0; i < lenSlice; i++ {
		indexes = append(indexes, i)
	}
	return indexes
}

// Creates an array of (value, index) pairs and sorts them in descending order of value.
func sortIndexByValue(indexes []int, values []float64) byValue {
	pairs := make(byValue, len(values))

	for i := 0; i < len(values); i++ {
		pairs[i] = pair{values[i], indexes[i]}
	}
	sort.Sort(pairs)
	return pairs
}
