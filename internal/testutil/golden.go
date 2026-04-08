package testutil

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func AssertGoldenJSON(t testing.TB, path string, got any) {
	t.Helper()

	expected, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden file %s: %v", path, err)
	}

	actual, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		t.Fatalf("marshal golden payload: %v", err)
	}
	actual = append(actual, '\n')

	if bytes.Equal(bytes.TrimSpace(actual), bytes.TrimSpace(expected)) {
		return
	}

	t.Fatalf("golden mismatch for %s\nexpected:\n%s\nactual:\n%s", path, expected, actual)
}
