# VED

Back to [Inventory analyses](../inventory.md)

## Purpose

VED classifies items by operational criticality.

## When To Use

- when stockout impact matters more than value or demand volatility;
- when procurement or medical inventory policy depends on criticality.

## When Not To Use

- when there is no meaningful criticality score;
- when the score is not comparable across items.

## Inputs

- `SKU`, `Name`
- `CriticalityScore`
- optional thresholds `VitalMinScore`, `EssentialMinScore`

## Output

- `CriticalityScore`
- `Group` in `V/E/D`

## Rules

- default thresholds are `V >= 70`, `E >= 40`, otherwise `D`
- results are sorted by `CriticalityScore` descending

## Example

```go
out, err := analytics.New().Inventory().VED(ctx, ved.Input{
    Items: []ved.Item{
        {SKU: "A", Name: "Item A", CriticalityScore: 85},
    },
})
```

## Interpretation

`V` items are the most critical, `E` items are important but not vital, and
`D` items are the least critical under the configured scale.

## Common Mistakes

- using an arbitrary score without agreed business meaning;
- assuming the package derives criticality from other item fields.

## Related Analyses

- [ABC](abc.md)
- [HML](hml.md)
