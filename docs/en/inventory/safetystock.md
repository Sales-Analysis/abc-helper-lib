# Safety Stock

Back to [Inventory analyses](../inventory.md)

## Purpose

Safety Stock estimates buffer inventory against demand variability during lead
time.

## When To Use

- when demand uncertainty needs a stock buffer;
- when reorder policy is sensitive to service targets.

## When Not To Use

- when the input does not contain a meaningful demand standard deviation;
- when service factor assumptions are unclear.

## Inputs

- `DemandStdDev`
- `LeadTimePeriods`
- optional `ServiceFactor`

## Output

- normalized `ServiceFactor`
- `SafetyStock`

## Formula

- `SafetyStock = ServiceFactor * DemandStdDev * sqrt(LeadTimePeriods)`
- default `ServiceFactor = 1.65`

## Example

```go
out, err := analytics.New().Inventory().SafetyStock(ctx, safetystock.Input{
    Items: []safetystock.Item{
        {SKU: "A", Name: "Item A", DemandStdDev: 10, LeadTimePeriods: 4, ServiceFactor: 1.65},
    },
})
```

## Interpretation

Higher variability, higher lead time, or higher service factor increases safety
stock.

## Common Mistakes

- passing a service level percentage instead of a z-factor;
- using different time units for standard deviation and lead time.

## Related Analyses

- [Reorder Point](reorderpoint.md)
- [Service Level](servicelevel.md)
