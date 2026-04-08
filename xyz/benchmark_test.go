package xyz_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

func BenchmarkAnalyze(b *testing.B) {
	ctx := context.Background()
	input := xyz.Input{Items: benchmarkItems(1000)}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := xyz.Analyze(ctx, input); err != nil {
			b.Fatalf("Analyze returned error: %v", err)
		}
	}
}

func benchmarkItems(count int) []xyz.Item {
	items := make([]xyz.Item, count)
	for i := range items {
		items[i] = xyz.Item{
			SKU:  "SKU-" + strconv.Itoa(i),
			Name: "Item " + strconv.Itoa(i),
			Demands: []float64{
				float64((i % 15) + 5),
				float64((i % 15) + 7),
				float64((i % 15) + 3),
				float64((i % 15) + 6),
				float64((i % 15) + 4),
				float64((i % 15) + 8),
			},
		}
	}
	return items
}
