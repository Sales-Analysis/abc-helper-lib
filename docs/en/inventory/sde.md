# SDE

Back to [Inventory analyses](../inventory.md)

## Purpose

SDE classifies items by procurement difficulty using lead time.

## When To Use

- when supply difficulty is a stronger signal than demand or value;
- when you need a simple lead-time segmentation.

## When Not To Use

- when lead time is unstable but not captured in one comparable metric;
- when procurement difficulty depends on other factors not represented here.

## Inputs

- `SKU`, `Name`
- `LeadTimeDays`
- optional thresholds `ScarceMinLeadTime`, `DifficultMinLeadTime`

## Output

- `LeadTimeDays`
- `Group` in `S/D/E`

## Rules

- if thresholds are not provided, they are derived from the sorted lead time distribution
- results are sorted by `LeadTimeDays` descending
- the highest lead times map to `S`, then `D`, then `E`

## Example

```go
out, err := analytics.New().Inventory().SDE(ctx, sde.Input{
    Items: []sde.Item{
        {SKU: "A", Name: "Item A", LeadTimeDays: 45},
    },
})
```

## Interpretation

`S` items are the hardest to source, `D` items are moderately difficult, and
`E` items are the easiest under the dataset thresholds.

## Common Mistakes

- treating SDE as a service-level metric;
- comparing groups across very different datasets without fixing thresholds.

## Related Analyses

- [Safety Stock](safetystock.md)
- [Reorder Point](reorderpoint.md)
