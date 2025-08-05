package abc

type ABC struct {
	SKU              []int
	Name             []string
	Quantity         []int
	PriceUnit        []float64
	PriceTotal       []float64
	ShareTotal       []float64
	ShareAccumulated []float64
	Group            []string
}

// Product is the input struct for the analysis.
type Product struct {
	SKU      string
	Name     string
	Quantity int
	Price    float64
}

// pair is an unexported struct used to link a value with its original index for sorting.
type pair struct {
	value int
	index int
}
