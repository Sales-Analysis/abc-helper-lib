# Churn / Retention

Назад к [клиентским анализам](../customer.md)

## Назначение

Churn классифицирует клиентов по окну неактивности и считает summary по
retention и churn rates.

## Когда использовать

- когда нужно операционно мониторить статус клиента;
- когда retention-reporting опирается на простую inactivity-модель.

## Когда не использовать

- когда нужен predictive probability model;
- когда бизнес-определение churn не основано на временном окне.

## Вход

- `LastOrderAt`
- опциональный `AnalysisTime`
- опциональные пороги `AtRiskDays`, `ChurnDays`

## Выход

- `DaysSinceLastOrder`
- `Status`
- `IsRetained`
- `IsChurned`
- summary counts и rates

## Правила

- дефолтные пороги: `30` дней до `AtRisk` и `90` дней до `Churned`
- отсутствие даты последнего заказа маппится в возраст churn-порога
- summary отдает `RetentionRate` и `ChurnRate`

## Пример

```go
out, err := analytics.New().Customer().Churn(ctx, churn.Input{
    AnalysisTime: time.Now(),
    Customers: []churn.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -45)},
    },
})
```

## Интерпретация

Это rules-based status model. Он полезен для reporting и operations, но не
заменяет полноценную predictive churn-систему.

## Частые ошибки

- считать `AtRisk` уже ушедшим клиентом;
- использовать разные inactivity windows в разных командах.

## Связанные анализы

- [RFM](rfm.md)
- [CLV](clv.md)
