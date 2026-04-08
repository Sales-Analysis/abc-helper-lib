# CLV

Назад к [клиентским анализам](../customer.md)

## Назначение

CLV оценивает прогнозный customer lifetime value на основе retention, margin и
purchase behavior.

## Когда использовать

- когда нужна value-оценка клиентской экономики;
- когда acquisition cost нужно сравнивать с ожидаемой будущей ценностью.

## Когда не использовать

- когда неизвестны retention- и discount-assumptions;
- когда нужна только описательная сегментация.

## Вход

- `Revenue`
- `Orders`
- `PeriodsObserved`
- `GrossMarginRate`
- `RetentionRate`
- `DiscountRate`
- `AcquisitionCost`

## Выход

- `AverageOrderValue`
- `PurchaseFrequency`
- нормализованные rate-поля
- `PredictedCLV`

## Правила

- `AverageOrderValue = Revenue / Orders`
- `PurchaseFrequency = Orders / PeriodsObserved`
- rates в percent-форме `0..100` нормализуются в `0..1`
- если знаменатель `1 + DiscountRate - RetentionRate` неположительный,
  реализация возвращает `-AcquisitionCost`

## Пример

```go
out, err := analytics.New().Customer().CLV(ctx, clv.Input{
    Customers: []clv.Customer{
        {CustomerID: "C1", Name: "Customer 1", Revenue: 1200, Orders: 6, PeriodsObserved: 3, GrossMarginRate: 0.4, RetentionRate: 0.8, DiscountRate: 0.1, AcquisitionCost: 50},
    },
})
```

## Интерпретация

CLV — model-based оценка. Небольшие изменения в retention или discount
assumptions могут заметно изменить результат.

## Частые ошибки

- смешивать ratio и percent conventions без проверки нормализации;
- трактовать modeled CLV как уже полученную прибыль.

## Связанные анализы

- [RFM](rfm.md)
- [Churn / Retention](churn.md)
