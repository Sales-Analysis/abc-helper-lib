# Customer Analyses

This page is the index for customer-focused analyses implemented in
`abc-helper-lib`.

## Table Of Contents

- [RFM](customer/rfm.md)
- [CLV](customer/clv.md)
- [Churn / Retention](customer/churn.md)
- [Decision guide](decision-guide.md)
- [API conventions](api.md)

## Summary Table

| Analysis | Package | Facade Method | Primary Inputs | Main Output |
| :--- | :--- | :--- | :--- | :--- |
| [`RFM`](customer/rfm.md) | `rfm` | `Customer().RFM` | last order, order count, monetary value | RFM score and segment |
| [`CLV`](customer/clv.md) | `clv` | `Customer().CLV` | revenue, orders, retention, discount | predicted customer lifetime value |
| [`Churn/Retention`](customer/churn.md) | `churn` | `Customer().Churn` | last order date | customer status and summary rates |

## Related Technical Notes

- [Decision guide](decision-guide.md)
- [API conventions](api.md)
- `RFM` is a segmentation model
- `CLV` is a value estimation model
- `Churn` is a retention status model
