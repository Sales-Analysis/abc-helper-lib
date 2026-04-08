# RFM

Назад к [клиентским анализам](../customer.md)

## Назначение

RFM сегментирует клиентов по recency, frequency и monetary value.

## Когда использовать

- когда нужна поведенческая сегментация клиентов;
- когда marketing или CRM-actions зависят от quality bands клиента.

## Когда не использовать

- когда выборка слишком мала для rank-based segmentation;
- когда нужна денежная оценка, а не сегментация.

## Вход

- `LastOrderAt`
- `Orders`
- `MonetaryValue`
- опциональный `AnalysisTime`

## Выход

- `RecencyDays`
- `RScore`, `FScore`, `MScore`
- `RFMScore`
- `Segment`

## Правила

- каждая ось ранжируется по шкале `1..5`
- меньшая recency лучше, большая frequency и monetary value лучше
- отсутствие даты последнего заказа трактуется как очень старая recency
- результаты сортируются по `RFMScore` по убыванию

## Пример

```go
out, err := analytics.New().Customer().RFM(ctx, rfm.Input{
    AnalysisTime: time.Now(),
    Customers: []rfm.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -10), Orders: 5, MonetaryValue: 1200},
    },
})
```

## Интерпретация

RFM — относительная модель сегментации. Scores зависят от текущего датасета,
а не от универсальных fixed thresholds.

## Частые ошибки

- сравнивать scores между сильно разными датасетами без контекста;
- ожидать threshold-based модель вместо rank-based scoring.

## Связанные анализы

- [CLV](clv.md)
- [Churn / Retention](churn.md)
