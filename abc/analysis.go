package abc

import (
	"fmt"
	"sort"
)

func New() *ABC {
	return &ABC{}
}

func (a *ABC) Calculate(products []Product) int {
	priceTotal := a.calculatePriceTotal(products)
	grandTotal := a.calculateGrandTotal(priceTotal)
	pairs := a.rankProductsByValue(priceTotal)
	costPercentage := calculateCostPercentage(pairs, grandTotal)
	fmt.Println(costPercentage)
	return 0
}

// Determining each product's contribution to total revenue
func (a *ABC) calculatePriceTotal(products []Product) []float64 {
	total := []float64{}
	for _, value := range products {
		total = append(total, float64(value.Quantity)*value.Price)
	}
	return total
}

// CalculateGrandTotal sums a slice of totals and returns the grand total.
func (a *ABC) calculateGrandTotal(totals []float64) float64 {
	var grandTotal float64
	for _, total := range totals {
		grandTotal += total
	}
	return grandTotal
}

func (a byValue) Len() int           { return len(a) }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValue) Less(i, j int) bool { return a[i].value > a[j].value }

// Ranking products by importance
func (a *ABC) rankProductsByValue(priceTotal []float64) byValue {
	indexes := createIndexesSlice(len(priceTotal))
	pairs := sortIndexByValue(indexes, priceTotal)
	return pairs
}

func createIndexesSlice(lenSlice int) []int {
	indexes := []int{}
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

func calculateCostPercentage(pairs byValue, grandTotal float64) []float64 {
	costPercentage := []float64{}
	for _, value := range pairs {
		v := (value.value / grandTotal) * 100
		costPercentage = append(costPercentage, v)
	}
	return costPercentage
}
