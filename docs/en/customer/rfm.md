# RFM

Back to [Customer analyses](../customer.md)

## Purpose

RFM segments customers by recency, frequency, and monetary value.

## When To Use

- when you need behavioral customer segmentation;
- when marketing or CRM actions depend on customer quality bands.

## When Not To Use

- when the dataset is too small for ranking-based segmentation;
- when you need monetary prediction instead of segmentation.

## Inputs

- `LastOrderAt`
- `Orders`
- `MonetaryValue`
- optional `AnalysisTime`

## Output

- `RecencyDays`
- `RScore`, `FScore`, `MScore`
- `RFMScore`
- `Segment`

## Rules

- each metric is ranked onto a `1..5` scale
- lower recency days are better, higher frequency and monetary value are better
- missing last order dates are treated as extremely old recency
- results are sorted by `RFMScore` descending

## Example

```go
out, err := analytics.New().Customer().RFM(ctx, rfm.Input{
    AnalysisTime: time.Now(),
    Customers: []rfm.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -10), Orders: 5, MonetaryValue: 1200},
    },
})
```

## Interpretation

RFM is a relative segmentation method. Scores depend on the current dataset,
not on universal thresholds.

## Common Mistakes

- comparing scores across very different datasets without context;
- expecting a fixed threshold-based model instead of rank-based scoring.

## Related Analyses

- [CLV](clv.md)
- [Churn / Retention](churn.md)
