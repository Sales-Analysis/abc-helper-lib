package customer

import (
	"context"

	"github.com/Sales-Analysis/abc-helper-lib/churn"
	"github.com/Sales-Analysis/abc-helper-lib/clv"
	"github.com/Sales-Analysis/abc-helper-lib/rfm"
)

// Facade groups customer-focused analytics behind a small orchestration layer.
type Facade struct{}

func New() *Facade {
	return &Facade{}
}

func (f *Facade) RFM(ctx context.Context, input rfm.Input) (rfm.Output, error) {
	return rfm.Analyze(ctx, input)
}

func (f *Facade) CLV(ctx context.Context, input clv.Input) (clv.Output, error) {
	return clv.Analyze(ctx, input)
}

func (f *Facade) Churn(ctx context.Context, input churn.Input) (churn.Output, error) {
	return churn.Analyze(ctx, input)
}
