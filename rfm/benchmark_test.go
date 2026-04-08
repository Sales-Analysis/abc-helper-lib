package rfm_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/Sales-Analysis/abc-helper-lib/rfm"
)

func BenchmarkAnalyze(b *testing.B) {
	ctx := context.Background()
	analysisTime := time.Date(2026, time.April, 8, 0, 0, 0, 0, time.UTC)
	input := rfm.Input{
		AnalysisTime: analysisTime,
		Customers:    benchmarkCustomers(analysisTime, 1000),
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := rfm.Analyze(ctx, input); err != nil {
			b.Fatalf("Analyze returned error: %v", err)
		}
	}
}

func benchmarkCustomers(analysisTime time.Time, count int) []rfm.Customer {
	customers := make([]rfm.Customer, count)
	for i := range customers {
		customers[i] = rfm.Customer{
			CustomerID:    "C-" + strconv.Itoa(i),
			Name:          "Customer " + strconv.Itoa(i),
			LastOrderAt:   analysisTime.AddDate(0, 0, -((i % 180) + 1)),
			Orders:        (i % 20) + 1,
			MonetaryValue: float64((i%50)+1) * 35.5,
		}
	}
	return customers
}
