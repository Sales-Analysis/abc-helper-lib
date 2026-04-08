# EOQ

Back to [Inventory analyses](../inventory.md)

## Purpose

EOQ calculates the economic order quantity that balances ordering and holding
costs.

## When To Use

- when you need a baseline replenishment quantity;
- when annual demand, ordering cost, and holding cost are known.

## When Not To Use

- when input costs are not trustworthy;
- when replenishment constraints make the classic EOQ model unrealistic.

## Inputs

- `AnnualDemand`
- `OrderingCost`
- `HoldingCost`

## Output

- `OptimalQuantity`
- `OrdersPerYear`

## Formula

- `EOQ = sqrt((2 * AnnualDemand * OrderingCost) / HoldingCost)`
- `OrdersPerYear = AnnualDemand / EOQ`

## Example

```go
out, err := analytics.New().Inventory().EOQ(ctx, eoq.Input{
    Items: []eoq.Item{
        {SKU: "A", Name: "Item A", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
    },
})
```

## Interpretation

Higher demand or ordering cost raises EOQ. Higher holding cost lowers EOQ.

## Common Mistakes

- mixing annual and monthly demand with annual cost assumptions;
- expecting EOQ to account for safety stock or supplier lead time directly.

## Related Analyses

- [Safety Stock](safetystock.md)
- [Reorder Point](reorderpoint.md)
