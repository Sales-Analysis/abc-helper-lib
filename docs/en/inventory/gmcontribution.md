# GM/Contribution

Back to [Inventory analyses](../inventory.md)

## Purpose

GM/Contribution calculates gross margin and contribution metrics from revenue
and cost components.

## When To Use

- when assortment decisions need margin quality, not just revenue;
- when variable and fixed cost structure matters.

## When Not To Use

- when only top-line sales value is available;
- when cost allocation is too noisy to support per-item interpretation.

## Inputs

- `Revenue`
- `COGS`
- `VariableCost`
- `FixedCost`

## Output

- `GrossMargin`
- `GrossMarginRate`
- `ContributionMargin`
- `ContributionRate`
- `NetContribution`

## Rules

- `GrossMargin = Revenue - COGS`
- `ContributionMargin = Revenue - VariableCost`
- `NetContribution = ContributionMargin - FixedCost`
- rate metrics are expressed in percent of revenue

## Example

```go
out, err := analytics.New().Inventory().GMContribution(ctx, gmcontribution.Input{
    Items: []gmcontribution.Item{
        {SKU: "A", Name: "Item A", Revenue: 1000, COGS: 600, VariableCost: 700, FixedCost: 100},
    },
})
```

## Interpretation

This analysis separates gross margin from contribution logic, which is useful
when variable cost and fixed cost tell different business stories.

## Common Mistakes

- confusing `COGS` and `VariableCost`;
- comparing items only by revenue after computing richer margin metrics.

## Related Analyses

- [Pareto 80/20](pareto.md)
- [ABC](abc.md)
