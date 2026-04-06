package abcxyz

import (
	"context"

	"github.com/Sales-Analysis/abc-helper-lib/abc"
	"github.com/Sales-Analysis/abc-helper-lib/xyz"
)

type Input struct {
	Items         []Item
	ABCThresholds abc.Thresholds
	XYZThresholds xyz.Thresholds
}

type Output struct {
	Results []ItemResult
}

type Item struct {
	SKU      string
	Name     string
	Quantity int
	Price    float64
	Demands  []float64
}

type ItemResult struct {
	OriginalIndex          int
	SKU                    string
	Name                   string
	Quantity               int
	PriceUnit              float64
	PriceTotal             float64
	ShareTotal             float64
	ShareAccumulated       float64
	ABCGroup               string
	AverageDemand          float64
	StandardDeviation      float64
	CoefficientOfVariation float64
	XYZGroup               string
	CombinedGroup          string
}

func Analyze(ctx context.Context, input Input) (Output, error) {
	if ctx != nil {
		if err := ctx.Err(); err != nil {
			return Output{}, err
		}
	}

	abcProducts := make([]abc.Product, len(input.Items))
	xyzItems := make([]xyz.Item, len(input.Items))

	for i, item := range input.Items {
		abcProducts[i] = abc.Product{
			SKU:      item.SKU,
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
		}
		xyzItems[i] = xyz.Item{
			SKU:     item.SKU,
			Name:    item.Name,
			Demands: append([]float64(nil), item.Demands...),
		}
	}

	abcOutput, err := abc.AnalyzeDetailed(ctx, abc.Input{
		Products:   abcProducts,
		Thresholds: input.ABCThresholds,
	})
	if err != nil {
		return Output{}, err
	}

	xyzOutput, err := xyz.Analyze(ctx, xyz.Input{
		Items:      xyzItems,
		Thresholds: input.XYZThresholds,
	})
	if err != nil {
		return Output{}, err
	}

	xyzByIndex := make(map[int]xyz.ItemResult, len(xyzOutput.Results))
	for _, result := range xyzOutput.Results {
		xyzByIndex[result.OriginalIndex] = result
	}

	results := make([]ItemResult, len(abcOutput.Results))
	for i, result := range abcOutput.Results {
		xyzResult := xyzByIndex[result.OriginalIndex]
		source := input.Items[result.OriginalIndex]

		results[i] = ItemResult{
			OriginalIndex:          result.OriginalIndex,
			SKU:                    result.SKU,
			Name:                   result.Name,
			Quantity:               source.Quantity,
			PriceUnit:              source.Price,
			PriceTotal:             result.PriceTotal,
			ShareTotal:             result.ShareTotal,
			ShareAccumulated:       result.ShareAccumulated,
			ABCGroup:               result.Group,
			AverageDemand:          xyzResult.AverageDemand,
			StandardDeviation:      xyzResult.StandardDeviation,
			CoefficientOfVariation: xyzResult.CoefficientOfVariation,
			XYZGroup:               xyzResult.Group,
			CombinedGroup:          result.Group + xyzResult.Group,
		}
	}

	return Output{Results: results}, nil
}
