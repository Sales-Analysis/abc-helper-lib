# EOQ

Назад к [товарным анализам](../inventory.md)

## Назначение

EOQ рассчитывает экономически оптимальный размер заказа, балансируя ordering и
holding cost.

## Когда использовать

- когда нужен базовый размер пополнения;
- когда известны annual demand, ordering cost и holding cost.

## Когда не использовать

- когда cost-входы ненадежны;
- когда ограничения пополнения делают классическую EOQ-модель нереалистичной.

## Вход

- `AnnualDemand`
- `OrderingCost`
- `HoldingCost`

## Выход

- `OptimalQuantity`
- `OrdersPerYear`

## Формула

- `EOQ = sqrt((2 * AnnualDemand * OrderingCost) / HoldingCost)`
- `OrdersPerYear = AnnualDemand / EOQ`

## Пример

```go
out, err := analytics.New().Inventory().EOQ(ctx, eoq.Input{
    Items: []eoq.Item{
        {SKU: "A", Name: "Item A", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
    },
})
```

## Интерпретация

Рост спроса или ordering cost увеличивает EOQ. Рост holding cost уменьшает EOQ.

## Частые ошибки

- смешивать annual и monthly demand при annual cost assumptions;
- ожидать, что EOQ сам учитывает safety stock или supplier lead time.

## Связанные анализы

- [Safety Stock](safetystock.md)
- [Reorder Point](reorderpoint.md)
