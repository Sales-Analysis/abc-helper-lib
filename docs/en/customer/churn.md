# Churn / Retention

Back to [Customer analyses](../customer.md)

## Purpose

Churn classifies customers by inactivity window and produces retention and
churn summary rates.

## When To Use

- when customer status should be monitored operationally;
- when retention reporting depends on a simple inactivity model.

## When Not To Use

- when churn needs a predictive probability model;
- when the business definition of churn is not time-window based.

## Inputs

- `LastOrderAt`
- optional `AnalysisTime`
- optional thresholds `AtRiskDays`, `ChurnDays`

## Output

- `DaysSinceLastOrder`
- `Status`
- `IsRetained`
- `IsChurned`
- summary counts and rates

## Rules

- default thresholds are `30` days for `AtRisk` and `90` days for `Churned`
- missing last order dates are mapped to churn-threshold age
- summary exposes `RetentionRate` and `ChurnRate`

## Example

```go
out, err := analytics.New().Customer().Churn(ctx, churn.Input{
    AnalysisTime: time.Now(),
    Customers: []churn.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -45)},
    },
})
```

## Interpretation

This is a rules-based status model. It is useful for reporting and operations,
not as a full predictive churn system.

## Common Mistakes

- treating `AtRisk` as already churned;
- using inconsistent inactivity windows across teams.

## Related Analyses

- [RFM](rfm.md)
- [CLV](clv.md)
