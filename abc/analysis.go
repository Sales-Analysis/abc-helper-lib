package abc

import "fmt"

type ABC struct {
	SKU              []int
	Name             []string
	Quantity         []int
	PriceUnit        []float64
	PriceAll         []float64
	ShareTotal       []float64
	ShareAccumulated []float64
	Group            []string
}

func New() *ABC {
	return &ABC{}
}

func (a *ABC) Calculate(products []Product) int {
	total := a.CalculatePriceTotal(products)
	fmt.Println(total)
	return 0
}

// Determining each product's contribution to total revenue
func (a *ABC) CalculatePriceTotal(products []Product) []float64 {
	total := []float64{}
	for _, value := range products {
		total = append(total, float64(value.Quantity)*value.Price)
	}
	return total
}
