package analytics

import "context"

type Analyzer[I any, O any] interface {
	Analyze(ctx context.Context, input I) (O, error)
}
