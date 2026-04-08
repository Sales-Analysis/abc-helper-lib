package abc

import (
	"context"
	"strconv"
	"testing"
)

func BenchmarkAnalyze(b *testing.B) {
	ctx := context.Background()
	input := Input{Items: benchmarkItems(1000)}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := Analyze(ctx, input); err != nil {
			b.Fatalf("Analyze returned error: %v", err)
		}
	}
}

func benchmarkItems(count int) []Item {
	items := make([]Item, count)
	for i := range items {
		items[i] = Item{
			SKU:      "SKU-" + strconv.Itoa(i),
			Name:     "Item " + strconv.Itoa(i),
			Quantity: (i % 25) + 1,
			Price:    float64((i%100)+1) * 1.75,
		}
	}
	return items
}
