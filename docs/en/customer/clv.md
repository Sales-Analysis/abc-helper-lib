# CLV

Back to [Customer analyses](../customer.md)

## Purpose

CLV estimates predicted customer lifetime value from retention, margin, and
purchase behavior.

## When To Use

- when you need a value estimate for customer economics;
- when acquisition cost should be compared against expected future value.

## When Not To Use

- when retention and discount assumptions are unknown;
- when you only need descriptive segmentation.

## Inputs

- `Revenue`
- `Orders`
- `PeriodsObserved`
- `GrossMarginRate`
- `RetentionRate`
- `DiscountRate`
- `AcquisitionCost`

## Output

- `AverageOrderValue`
- `PurchaseFrequency`
- normalized rate fields
- `PredictedCLV`

## Rules

- `AverageOrderValue = Revenue / Orders`
- `PurchaseFrequency = Orders / PeriodsObserved`
- percent-form rates in `0..100` are normalized to `0..1`
- if the denominator `1 + DiscountRate - RetentionRate` is non-positive, the implementation returns `-AcquisitionCost`

## Example

```go
out, err := analytics.New().Customer().CLV(ctx, clv.Input{
    Customers: []clv.Customer{
        {CustomerID: "C1", Name: "Customer 1", Revenue: 1200, Orders: 6, PeriodsObserved: 3, GrossMarginRate: 0.4, RetentionRate: 0.8, DiscountRate: 0.1, AcquisitionCost: 50},
    },
})
```

## Interpretation

CLV is model-based. Small changes in retention or discount assumptions can
change the result significantly.

## Common Mistakes

- mixing ratio and percent conventions without checking normalization;
- treating modeled CLV as booked profit.

## Related Analyses

- [RFM](rfm.md)
- [Churn / Retention](churn.md)
