package abc

import (
	"fmt"
)

type ABC struct {
	ID []int
	Name []string
	count []int
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
