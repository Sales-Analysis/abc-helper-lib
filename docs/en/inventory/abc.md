# ABC

Back to [Inventory analyses](../inventory.md)

## Purpose

ABC classifies items by revenue contribution after sorting them by total sales
value.

## When To Use

- when you need to identify high-impact SKUs by revenue;
- when assortment prioritization depends on contribution rather than movement;
- as the value axis for [ABC-XYZ](abcxyz.md).

## When Not To Use

- when the main concern is demand stability rather than value concentration;
- when unit cost or criticality should drive the grouping instead.

## Inputs

- `SKU`, `Name`
- `Quantity`
- `Price`
- preferred collection field: `Items`
- legacy collection alias: `Products`
- optional thresholds `AMaxShare`, `BMaxShare`

## Output

- `OriginalIndex`
- `PriceTotal`
- `ShareTotal`
- `ShareAccumulated`
- `Group` in `A/B/C`
- `TotalRevenue`
- `Summary` with total items and `A/B/C` counts

## Rules

- `PriceTotal = Quantity * Price`
- items are sorted by `PriceTotal` descending
- `ShareTotal = PriceTotal / TotalRevenue * 100`
- `ShareAccumulated` is calculated cumulatively in sorted order
- default thresholds are `A <= 80%`, `B <= 95%`, otherwise `C`

## Example

```go
out, err := analytics.New().Inventory().ABC(ctx, abc.Input{
    Items: []abc.Item{
        {SKU: "A", Name: "Item A", Quantity: 10, Price: 100},
        {SKU: "B", Name: "Item B", Quantity: 5, Price: 50},
    },
})
```

## Interpretation

`A` items are the main revenue drivers. `B` items are important but secondary.
`C` items are low-contribution items under the configured thresholds.

## Common Mistakes

- mixing unit price and total revenue in the input;
- using legacy `Products` in new integrations instead of the preferred `Items`;
- expecting the output order to match the input order;
- using ABC to infer demand stability.

## Related Analyses

- [XYZ](xyz.md)
- [ABC-XYZ](abcxyz.md)
- [Pareto 80/20](pareto.md)
