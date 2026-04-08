# FSN

Back to [Inventory analyses](../inventory.md)

## Purpose

FSN classifies items by movement frequency across periods.

## When To Use

- when you need to separate fast-moving, slow-moving, and non-moving items;
- when movement history is more relevant than revenue share.

## When Not To Use

- when only demand variability matters;
- when movement history is too short or not period-based.

## Inputs

- `SKU`, `Name`
- `Movements []float64`
- optional thresholds `FastMinActivityRate`, `SlowMinActivityRate`

## Output

- `Periods`
- `ActivePeriods`
- `ActivityRate`
- `TotalMovement`
- `AverageMovement`
- `LastMovementPeriod`
- `Group` in `F/S/N`

## Rules

- active periods are periods with movement greater than zero
- `ActivityRate = ActivePeriods / Periods * 100`
- default thresholds are `F >= 70%`, `S >= 1%`, otherwise `N`
- items with zero active periods are always `N`

## Example

```go
out, err := analytics.New().Inventory().FSN(ctx, fsn.Input{
    Items: []fsn.Item{
        {SKU: "A", Name: "Item A", Movements: []float64{3, 0, 2, 1}},
    },
})
```

## Interpretation

`F` items move frequently, `S` items move occasionally, and `N` items are
inactive under the chosen horizon.

## Common Mistakes

- confusing activity rate with volume;
- treating a high single movement as enough to make an item `F`.

## Related Analyses

- [XYZ](xyz.md)
- [Pareto 80/20](pareto.md)
