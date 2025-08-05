package abc

import "fmt"

func New() *ABC {
	return &ABC{}
}

func (a *ABC) Calculate(products []Product) int {
	a.PriceTotal = a.calculatePriceTotal(products)
	a.rankProductsByValue(products)
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

func (a *ABC) rankProductsByValue(products []Product) int {
	// Привести к типу ABC, не изменять 'products []Product'
	fmt.Println(a.PriceTotal)
	indexes := []int{}
	for i := range a.PriceTotal {
		indexes = append(indexes, i)
	}
	fmt.Println(indexes)
	return 0
}
