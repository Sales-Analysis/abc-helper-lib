# Pareto 80/20

Back to [Inventory analyses](../inventory.md)

## Purpose

Pareto checks whether a small share of items accounts for a large share of
value.

## When To Use

- when you want a quick concentration view over item value;
- when you need to confirm or challenge an `80/20` assumption.

## When Not To Use

- when item value is not the correct concentration metric;
- when you need classification thresholds like `ABC` rather than a summary rule.

## Inputs

- `Value`
- optional thresholds `TopValueShare`, `TopItemShare`

## Output

- per-item `ValueShare`
- `CumulativeValueShare`
- `CumulativeItemShare`
- top-set flags
- summary with `ParetoPrincipleMet`

## Rules

- default thresholds are `80%` value and `20%` items
- top item count is rounded up
- summary checks whether the configured top item share reaches the configured top value share

## Example

```go
out, err := analytics.New().Inventory().Pareto(ctx, pareto.Input{
    Items: []pareto.Item{
        {SKU: "A", Name: "Item A", Value: 1000},
    },
})
```

## Interpretation

This analysis is about concentration, not direct operational grouping.

## Common Mistakes

- interpreting the result as a replacement for `ABC`;
- forgetting that thresholds are configurable and not fixed to `80/20`.

## Related Analyses

- [ABC](abc.md)
- [GM/Contribution](gmcontribution.md)
