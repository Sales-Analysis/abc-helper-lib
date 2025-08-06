package abc

import (
	"fmt"
	"sort"
)

func New() *ABC {
	return &ABC{}
}

func (a *ABC) Calculate(products []Product) int {
	a.PriceTotal = a.calculatePriceTotal(products)
	grandTotal := a.calculateGrandTotal(a.PriceTotal)
	a.rankProductsByValue(products)
	calculateCostPercentage(a.PriceTotal, grandTotal)
	fmt.Println(grandTotal)
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
func (a *ABC) rankProductsByValue(products []Product) int {
	indexes := []int{}
	for i := range a.PriceTotal {
		indexes = append(indexes, i)
	}

	pairs := make(byValue, len(a.PriceTotal))

	for i := 0; i < len(a.PriceTotal); i++ {
		pairs[i] = pair{a.PriceTotal[i], indexes[i]}
	}
	sort.Sort(pairs)

	for i, value := range pairs {
		a.PriceTotal[i] = value.value
		a.SKU = append(a.SKU, products[value.index].SKU)
		a.Name = append(a.Name, products[value.index].Name)
		a.Quantity = append(a.Quantity, products[value.index].Quantity)
		a.PriceUnit = append(a.PriceUnit, products[value.index].Price)
	}
	return 0
}

func calculateCostPercentage(priceTotal []float64, grandTotal float64) []float64 {
	costPercentage := []float64{}
	for _, value := range priceTotal {
		v := (value / grandTotal) * 100
		costPercentage = append(costPercentage, v)
	}
	return costPercentage
}
