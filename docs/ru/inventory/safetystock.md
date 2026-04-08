# Safety Stock

Назад к [товарным анализам](../inventory.md)

## Назначение

Safety Stock оценивает буферный запас против вариативности спроса на горизонте
lead time.

## Когда использовать

- когда неопределенность спроса требует буферного запаса;
- когда политика пополнения зависит от service target.

## Когда не использовать

- когда во входе нет осмысленного стандартного отклонения спроса;
- когда assumptions по service factor не согласованы.

## Вход

- `DemandStdDev`
- `LeadTimePeriods`
- опциональный `ServiceFactor`

## Выход

- нормализованный `ServiceFactor`
- `SafetyStock`

## Формула

- `SafetyStock = ServiceFactor * DemandStdDev * sqrt(LeadTimePeriods)`
- дефолтный `ServiceFactor = 1.65`

## Пример

```go
out, err := analytics.New().Inventory().SafetyStock(ctx, safetystock.Input{
    Items: []safetystock.Item{
        {SKU: "A", Name: "Item A", DemandStdDev: 10, LeadTimePeriods: 4, ServiceFactor: 1.65},
    },
})
```

## Интерпретация

Чем выше вариативность, lead time или service factor, тем выше safety stock.

## Частые ошибки

- передавать процент service level вместо z-factor;
- использовать разные единицы времени для stddev спроса и lead time.

## Связанные анализы

- [Reorder Point](reorderpoint.md)
- [Service Level](servicelevel.md)
