package inventory

import (
	"context"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/abcxyz"
	"github.com/Sales-Analysis/abc-helper-lib/eoq"
	"github.com/Sales-Analysis/abc-helper-lib/fsn"
	"github.com/Sales-Analysis/abc-helper-lib/hml"
	"github.com/Sales-Analysis/abc-helper-lib/reorderpoint"
	"github.com/Sales-Analysis/abc-helper-lib/safetystock"
	"github.com/Sales-Analysis/abc-helper-lib/sde"
	"github.com/Sales-Analysis/abc-helper-lib/ved"
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

func (f *Facade) VED(ctx context.Context, input ved.Input) (ved.Output, error) {
	return ved.Analyze(ctx, input)
}

func (f *Facade) FSN(ctx context.Context, input fsn.Input) (fsn.Output, error) {
	return fsn.Analyze(ctx, input)
}

func (f *Facade) HML(ctx context.Context, input hml.Input) (hml.Output, error) {
	return hml.Analyze(ctx, input)
}

func (f *Facade) SDE(ctx context.Context, input sde.Input) (sde.Output, error) {
	return sde.Analyze(ctx, input)
}

func (f *Facade) EOQ(ctx context.Context, input eoq.Input) (eoq.Output, error) {
	return eoq.Analyze(ctx, input)
}

func (f *Facade) ReorderPoint(ctx context.Context, input reorderpoint.Input) (reorderpoint.Output, error) {
	return reorderpoint.Analyze(ctx, input)
}

func (f *Facade) SafetyStock(ctx context.Context, input safetystock.Input) (safetystock.Output, error) {
	return safetystock.Analyze(ctx, input)
}
