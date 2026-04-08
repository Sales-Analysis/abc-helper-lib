# HML

Back to [Inventory analyses](../inventory.md)

## Purpose

HML classifies items by unit cost.

## When To Use

- when cost intensity matters for stocking policy;
- when you need a quick cost-based segmentation without revenue weighting.

## When Not To Use

- when total contribution matters more than unit cost;
- when unit costs are not comparable across the item set.

## Inputs

- `SKU`, `Name`
- `UnitCost`
- optional thresholds `HighMinUnitCost`, `MediumMinUnitCost`

## Output

- `UnitCost`
- `Group` in `H/M/L`
- `Summary` with total items and `H/M/L` counts

## Rules

- if thresholds are not provided, they are derived from the sorted data
- the implementation splits the sorted unit costs into three bands
- results are sorted by `UnitCost` descending

## Example

```go
out, err := analytics.New().Inventory().HML(ctx, hml.Input{
    Items: []hml.Item{
        {SKU: "A", Name: "Item A", UnitCost: 250},
    },
})
```

## Interpretation

`H` items are the most expensive per unit, `M` items sit in the middle band,
and `L` items are the lowest-cost items.

## Common Mistakes

- assuming HML is the same as ABC by value;
- forgetting that auto-derived thresholds depend on the dataset composition.

## Related Analyses

- [ABC](abc.md)
- [VED](ved.md)
