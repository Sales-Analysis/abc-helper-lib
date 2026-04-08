# Service Level

Back to [Inventory analyses](../inventory.md)

## Purpose

Service Level estimates fulfillment performance from demand and stockout data.

## When To Use

- when you want a service KPI for each item;
- when both unit fulfillment and cycle behavior matter.

## When Not To Use

- when only one metric exists and the output might be overinterpreted;
- when lead-time variability is needed directly in the calculation.

## Inputs

- `DemandedUnits`
- `FulfilledUnits`
- `TotalCycles`
- `StockoutCycles`

## Output

- `FillRate`
- `CycleServiceLevel`
- `StockoutRate`
- `UnfulfilledUnits`
- exported `ServiceLevel`

## Rules

- `FillRate = FulfilledUnits / DemandedUnits * 100`
- `CycleServiceLevel = (TotalCycles - StockoutCycles) / TotalCycles * 100`
- `StockoutRate = StockoutCycles / TotalCycles * 100`
- if `TotalCycles > 0`, `ServiceLevel` uses `CycleServiceLevel`
- otherwise `ServiceLevel` falls back to `FillRate`

## Example

```go
out, err := analytics.New().Inventory().ServiceLevel(ctx, servicelevel.Input{
    Items: []servicelevel.Item{
        {SKU: "A", Name: "Item A", DemandedUnits: 100, FulfilledUnits: 95, TotalCycles: 20, StockoutCycles: 2},
    },
})
```

## Interpretation

The package keeps both fill-rate and cycle-based views so you can see whether
the same item looks strong in one metric but weak in another.

## Common Mistakes

- assuming the exported `ServiceLevel` is always the fill rate;
- passing cycle counts without a consistent cycle definition.

## Related Analyses

- [Safety Stock](safetystock.md)
- [Reorder Point](reorderpoint.md)
