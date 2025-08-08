package abc

// Calculates the revenue for each product: quantity * price.
func (a *ABC) calculatePriceTotal(products []Product) []float64 {
	total := []float64{}
	for _, value := range products {
		total = append(total, float64(value.Quantity)*value.Price)
	}
	return total
}

// Sums all revenue values to get the total revenue across all products.
func (a *ABC) calculateGrandTotal(totals []float64) float64 {
	var grandTotal float64
	for _, total := range totals {
		grandTotal += total
	}
	return grandTotal
}

// Calculates each product’s share of the total revenue (in percentage).
func (a *ABC) calculateCostPercentage(pairs byValue, grandTotal float64) []float64 {
	costPercentage := []float64{}
	for _, value := range pairs {
		v := (value.value / grandTotal) * 100
		costPercentage = append(costPercentage, v)
	}
	return costPercentage
}

// Calculates cumulative revenue share — how much total percentage is reached as products are considered in order of importance.
func (a *ABC) calculateAccumulatedShare(costPercentage []float64) []float64 {
	accumulatedShare := []float64{}
	as := 0.0
	for _, value := range costPercentage {
		as += value
		accumulatedShare = append(accumulatedShare, as)
	}
	return accumulatedShare
}
