package abc

type ABC struct {
	SKU              []string
	Name             []string
	Quantity         []int
	PriceUnit        []float64
	PriceTotal       []float64
	ShareTotal       []float64
	ShareAccumulated []float64
	Group            []string
	Result           []ProductResult
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

// The pair is an unexported struct used to link a value with its original index for sorting.
type pair struct {
	value float64
	index int
}

// Private type for sorting
type byValue []pair
