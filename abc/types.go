package abc

type ABC struct {
	Result []ProductResult
}

type Thresholds struct {
	AMaxShare float64
	BMaxShare float64
}

type Input struct {
	Products   []Product
	Thresholds Thresholds
}

type Output struct {
	Results      []ProductResult
	TotalRevenue float64
}

type DetailedOutput struct {
	Results      []IndexedProductResult
	TotalRevenue float64
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

type IndexedProductResult struct {
	OriginalIndex int
	ProductResult
}

// The pair is an unexported struct used to link a value with its original index for sorting.
type pair struct {
	value float64
	index int
}
