package abc

// Assigns each product to one of three ABC categories
func (a *ABC) assignGroup(accumulatedShare []float64) []string {
	return assignGroupWithThresholds(accumulatedShare, Thresholds{})
}

func assignGroupWithThresholds(accumulatedShare []float64, thresholds Thresholds) []string {
	thresholds = thresholds.normalized()
	groups := []string{}
	for _, value := range accumulatedShare {
		if value <= thresholds.AMaxShare {
			groups = append(groups, "A")
		} else if value <= thresholds.BMaxShare {
			groups = append(groups, "B")
		} else {
			groups = append(groups, "C")
		}
	}
	return groups
}
