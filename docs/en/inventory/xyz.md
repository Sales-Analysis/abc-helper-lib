# XYZ

Back to [Inventory analyses](../inventory.md)

## Purpose

XYZ classifies items by demand variability using the coefficient of variation.

## When To Use

- when replenishment depends on demand stability;
- when you want to separate predictable demand from volatile demand;
- as the variability axis for [ABC-XYZ](abcxyz.md).

## When Not To Use

- when value contribution is more important than demand behavior;
- when the input does not contain a period-by-period demand history.

## Inputs

- `SKU`, `Name`
- `Demands []float64`
- optional thresholds `XMaxCV`, `YMaxCV`

## Output

- `OriginalIndex`
- `Periods`
- `AverageDemand`
- `StandardDeviation`
- `CoefficientOfVariation`
- `Group` in `X/Y/Z`
- `Summary` with total items and `X/Y/Z` counts

## Rules

- mean and standard deviation are calculated from the demand series
- `CoefficientOfVariation = StandardDeviation / AverageDemand * 100`
- default thresholds are `X <= 10`, `Y <= 25`, otherwise `Z`
- empty demand series or zero mean demand are classified as `Z`

## Example

```go
out, err := analytics.New().Inventory().XYZ(ctx, xyz.Input{
    Items: []xyz.Item{
        {SKU: "A", Name: "Item A", Demands: []float64{10, 12, 9}},
    },
})
```

## Interpretation

`X` items have the most stable demand. `Y` items are moderately variable. `Z`
items are highly variable or effectively inactive.

## Common Mistakes

- passing sales totals instead of period demand values;
- comparing CV thresholds without confirming the unit is percent;
- reading `Z` as “bad” without checking whether the item simply has zero demand.

## Related Analyses

- [ABC](abc.md)
- [ABC-XYZ](abcxyz.md)
- [FSN](fsn.md)
