package abc

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/Sales-Analysis/abc-helper-lib/internal/testutil"
)

func TestAnalyzeMatchesGoldenFile(t *testing.T) {
	output, err := Analyze(context.Background(), Input{
		Items: []Item{
			{SKU: "A", Name: "Alpha", Quantity: 10, Price: 10},
			{SKU: "B", Name: "Beta", Quantity: 5, Price: 5},
		},
	})
	if err != nil {
		t.Fatalf("Analyze returned error: %v", err)
	}

	testutil.AssertGoldenJSON(t, filepath.Join("testdata", "analyze.golden.json"), output)
}
