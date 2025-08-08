package abc

// Assigns each product to one of three ABC categories
func (a *ABC) assignGroup(accumulatedShare []float64) []string {
	groups := []string{}
	for _, value := range accumulatedShare {
		if value <= 80 {
			groups = append(groups, "A")
		} else if (value > 80) && (value <= 95) {
			groups = append(groups, "B")
		} else {
			groups = append(groups, "C")
		}
	}
	return groups
}
