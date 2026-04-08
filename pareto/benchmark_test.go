package pareto_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/pareto"
)

func BenchmarkAnalyze(b *testing.B) {
	ctx := context.Background()
	input := pareto.Input{Items: benchmarkItems(1000)}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := pareto.Analyze(ctx, input); err != nil {
			b.Fatalf("Analyze returned error: %v", err)
		}
	}
}

func benchmarkItems(count int) []pareto.Item {
	items := make([]pareto.Item, count)
	for i := range items {
		items[i] = pareto.Item{
			SKU:   "SKU-" + strconv.Itoa(i),
			Name:  "Item " + strconv.Itoa(i),
			Value: float64((count - i) * ((i % 7) + 1)),
		}
	}
	return items
}
