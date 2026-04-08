package rfm

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/internal/validation"
)

const missingLastOrderRecencyDays = 365000

type Input struct {
	Customers    []Customer
	AnalysisTime time.Time
}

type Output struct {
	Results []CustomerResult
}

type Customer struct {
	CustomerID    string
	Name          string
	LastOrderAt   time.Time
	Orders        int
	MonetaryValue float64
}

type CustomerResult struct {
	OriginalIndex int
	CustomerID    string
	Name          string
	RecencyDays   int
	Frequency     int
	MonetaryValue float64
	RScore        int
	FScore        int
	MScore        int
	RFMScore      string
	Segment       string
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	analysisTime := input.AnalysisTime
	if analysisTime.IsZero() {
		analysisTime = time.Now().UTC()
	}

	results := make([]CustomerResult, len(input.Customers))
	for i, customer := range input.Customers {
		if ctx != nil {
			if err := ctx.Err(); err != nil {
				return Output{}, err
			}
		}
		if err := validateCustomer(customer); err != nil {
			return Output{}, err
		}

		results[i] = CustomerResult{
			OriginalIndex: i,
			CustomerID:    customer.CustomerID,
			Name:          customer.Name,
			RecencyDays:   daysSinceLastOrder(analysisTime, customer.LastOrderAt),
			Frequency:     customer.Orders,
			MonetaryValue: customer.MonetaryValue,
		}
	}

	assignScores(results, func(item CustomerResult) float64 { return float64(item.RecencyDays) }, true, func(index int, score int) {
		results[index].RScore = score
	})
	assignScores(results, func(item CustomerResult) float64 { return float64(item.Frequency) }, false, func(index int, score int) {
		results[index].FScore = score
	})
	assignScores(results, func(item CustomerResult) float64 { return item.MonetaryValue }, false, func(index int, score int) {
		results[index].MScore = score
	})

	for i := range results {
		results[i].RFMScore = fmt.Sprintf("%d%d%d", results[i].RScore, results[i].FScore, results[i].MScore)
		results[i].Segment = segment(results[i])
	}

	sort.SliceStable(results, func(i, j int) bool {
		if results[i].RFMScore == results[j].RFMScore {
			return results[i].OriginalIndex < results[j].OriginalIndex
		}
		return results[i].RFMScore > results[j].RFMScore
	})

	return Output{Results: results}, nil
}

func validateCustomer(customer Customer) error {
	if err := validation.RequireNonNegativeInt("customers[].Orders", customer.Orders); err != nil {
		return err
	}
	return nil
}

func daysSinceLastOrder(analysisTime time.Time, lastOrderAt time.Time) int {
	if lastOrderAt.IsZero() {
		return missingLastOrderRecencyDays
	}
	if lastOrderAt.After(analysisTime) {
		return 0
	}
	return int(analysisTime.Sub(lastOrderAt).Hours() / 24)
}

func assignScores(results []CustomerResult, metric func(CustomerResult) float64, ascending bool, assign func(index int, score int)) {
	type rankedItem struct {
		Index int
		Value float64
	}

	ranked := make([]rankedItem, len(results))
	for i, result := range results {
		ranked[i] = rankedItem{
			Index: i,
			Value: metric(result),
		}
	}

	sort.SliceStable(ranked, func(i, j int) bool {
		if ranked[i].Value == ranked[j].Value {
			return ranked[i].Index < ranked[j].Index
		}
		if ascending {
			return ranked[i].Value < ranked[j].Value
		}
		return ranked[i].Value > ranked[j].Value
	})

	for rank, item := range ranked {
		assign(item.Index, scoreByRank(rank, len(ranked)))
	}
}

func scoreByRank(rank int, total int) int {
	if total <= 1 {
		return 5
	}
	score := 5 - int(math.Round((float64(rank)*4)/float64(total-1)))
	if score < 1 {
		return 1
	}
	if score > 5 {
		return 5
	}
	return score
}

func segment(result CustomerResult) string {
	switch {
	case result.RScore >= 4 && result.FScore >= 4 && result.MScore >= 4:
		return "Champions"
	case result.RScore >= 4 && result.FScore >= 3:
		return "Loyal"
	case result.RScore >= 3 && result.FScore >= 2 && result.MScore >= 2:
		return "Potential Loyalist"
	case result.RScore <= 2 && result.FScore >= 3:
		return "At Risk"
	case result.RScore <= 2 && result.FScore <= 2:
		return "Hibernating"
	default:
		return "Promising"
	}
}
