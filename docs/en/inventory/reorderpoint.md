# Reorder Point

Back to [Inventory analyses](../inventory.md)

## Purpose

Reorder Point calculates the inventory level at which replenishment should
start.

## When To Use

- when you already know average demand and lead time;
- when safety stock is part of the operating policy.

## When Not To Use

- when demand should be modeled from raw history inside the calculation;
- when safety stock has not been estimated yet.

## Inputs

- `AverageDemandPerPeriod`
- `LeadTimePeriods`
- `SafetyStock`

## Output

- `LeadTimeDemand`
- `ReorderPoint`

## Formula

- `LeadTimeDemand = AverageDemandPerPeriod * LeadTimePeriods`
- `ReorderPoint = LeadTimeDemand + SafetyStock`

## Example

```go
out, err := analytics.New().Inventory().ReorderPoint(ctx, reorderpoint.Input{
    Items: []reorderpoint.Item{
        {SKU: "A", Name: "Item A", AverageDemandPerPeriod: 20, LeadTimePeriods: 4, SafetyStock: 33},
    },
})
```

## Interpretation

The result is an operational trigger level, not an order quantity.

## Common Mistakes

- confusing reorder point with EOQ;
- mixing lead time units with demand period units.

## Related Analyses

- [Safety Stock](safetystock.md)
- [EOQ](eoq.md)
