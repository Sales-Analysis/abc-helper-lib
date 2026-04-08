# Reorder Point

Назад к [товарным анализам](../inventory.md)

## Назначение

Reorder Point рассчитывает уровень запаса, при котором нужно запускать
пополнение.

## Когда использовать

- когда уже известны средний спрос и lead time;
- когда safety stock входит в операционную политику.

## Когда не использовать

- когда спрос должен рассчитываться из сырой истории прямо внутри модели;
- когда safety stock еще не оценен.

## Вход

- `AverageDemandPerPeriod`
- `LeadTimePeriods`
- `SafetyStock`

## Выход

- `LeadTimeDemand`
- `ReorderPoint`

## Формула

- `LeadTimeDemand = AverageDemandPerPeriod * LeadTimePeriods`
- `ReorderPoint = LeadTimeDemand + SafetyStock`

## Пример

```go
out, err := analytics.New().Inventory().ReorderPoint(ctx, reorderpoint.Input{
    Items: []reorderpoint.Item{
        {SKU: "A", Name: "Item A", AverageDemandPerPeriod: 20, LeadTimePeriods: 4, SafetyStock: 33},
    },
})
```

## Интерпретация

Это операционный trigger level, а не размер заказа.

## Частые ошибки

- путать reorder point и EOQ;
- смешивать единицы lead time и единицы периода спроса.

## Связанные анализы

- [Safety Stock](safetystock.md)
- [EOQ](eoq.md)
