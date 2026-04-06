package inventory

import (
	"context"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/abcxyz"
	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

// Facade groups inventory-focused analyses behind a small orchestration layer.
type Facade struct{}

func New() *Facade {
	return &Facade{}
}

func (f *Facade) ABC(ctx context.Context, input abc.Input) (abc.Output, error) {
	return abc.Analyze(ctx, input)
}

func (f *Facade) XYZ(ctx context.Context, input xyz.Input) (xyz.Output, error) {
	return xyz.Analyze(ctx, input)
}

func (f *Facade) ABCXYZ(ctx context.Context, input abcxyz.Input) (abcxyz.Output, error) {
	return abcxyz.Analyze(ctx, input)
}
