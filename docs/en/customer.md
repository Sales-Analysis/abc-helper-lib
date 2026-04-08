# Customer Analyses

This page summarizes the customer-focused analyses currently implemented in
`abc-helper-lib`.

## Summary Table

| Analysis | Package | Facade Method | Primary Inputs | Main Output |
| :--- | :--- | :--- | :--- | :--- |
| `RFM` | `rfm` | `Customer().RFM` | last order, order count, monetary value | RFM score and segment |
| `CLV` | `clv` | `Customer().CLV` | revenue, orders, retention, discount | predicted customer lifetime value |
| `Churn/Retention` | `churn` | `Customer().Churn` | last order date | customer status and summary rates |

## RFM

- Purpose: segment customers by recency, frequency, and monetary value.
- Input:
  `LastOrderAt`, `Orders`, `MonetaryValue`, and optional `AnalysisTime`.
- Scoring:
  each dimension is ranked on a `1..5` scale using relative position in the
  dataset.
- Segment labels currently include:
  `Champions`, `Loyal`, `Potential Loyalist`, `At Risk`, `Hibernating`,
  `Promising`.
- Special case:
  a missing last order date is treated as extremely old recency.

## CLV

- Purpose: estimate predicted customer lifetime value.
- Input:
  `Revenue`, `Orders`, `PeriodsObserved`, `GrossMarginRate`,
  `RetentionRate`, `DiscountRate`, `AcquisitionCost`.
- Derived metrics:
  average order value and purchase frequency.
- Formula:
  `PredictedCLV = (AOV * PurchaseFrequency * GrossMarginRate * (RetentionRate / (1 + DiscountRate - RetentionRate))) - AcquisitionCost`.
- Ratio normalization:
  if `GrossMarginRate`, `RetentionRate`, or `DiscountRate` are passed in the
  `0..100` range, they are converted to `0..1`.
- If the denominator is non-positive or essential business inputs are missing,
  the implementation returns `-AcquisitionCost`.

## Churn / Retention

- Purpose: classify customers by inactivity window and summarize retention.
- Input:
  `LastOrderAt`, optional `AnalysisTime`, optional thresholds.
- Default thresholds:
  `AtRisk = 30 days`, `Churned = 90 days`.
- Output per customer:
  `DaysSinceLastOrder`, `Status`, `IsRetained`, `IsChurned`.
- Summary output:
  total customers, active count, at-risk count, churned count, retention rate,
  and churn rate.
