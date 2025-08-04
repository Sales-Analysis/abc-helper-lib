package abc

import (
	"fmt"
)

type ABC struct {
	SKU []int
	Name []string
	Quantity []int
	PriceUnit []float64
	PriceAll []float64
	ShareTotal []float64
	ShareAccumulated []float64
	Group []string
}


func New() *ABC {
	return &ABC{}
}


func (a *ABC) Calculate(products []Product) int{
	fmt.Println(products)
	return 0
}
