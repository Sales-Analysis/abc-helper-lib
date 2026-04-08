# ABC-XYZ

Back to [Inventory analyses](../inventory.md)

## Purpose

ABC-XYZ combines revenue importance and demand stability into one compound
classification such as `AX` or `CZ`.

## When To Use

- when assortment decisions depend on both contribution and predictability;
- when you need a compact grid for replenishment or planning policy.

## When Not To Use

- when only one axis matters and the second axis would add noise;
- when value and demand histories do not refer to the same item set.

## Inputs

- all item fields needed for `ABC`
- all demand fields needed for `XYZ`
- optional ABC thresholds
- optional XYZ thresholds

## Output

- all main `ABC` outputs
- all main `XYZ` outputs
- `CombinedGroup` such as `AX`, `BY`, or `CZ`
- `Summary` with aggregate counts for `ABC`, `XYZ`, and combined groups

## Rules

- the package runs `ABC` and `XYZ` separately
- results are merged by `OriginalIndex`
- `CombinedGroup = ABCGroup + XYZGroup`

## Example

```go
out, err := analytics.New().Inventory().ABCXYZ(ctx, abcxyz.Input{
    Items: []abcxyz.Item{
        {SKU: "A", Name: "Item A", Quantity: 10, Price: 100, Demands: []float64{10, 12, 9}},
    },
})
```

## Interpretation

`AX` usually means strategically important and stable. `CZ` usually means low
contribution and volatile demand.

## Common Mistakes

- assuming the package recalculates demand from quantity and price;
- passing mismatched datasets for value and demand;
- ignoring the fact that the output follows the `ABC` sorted order.

## Related Analyses

- [ABC](abc.md)
- [XYZ](xyz.md)
