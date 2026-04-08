# Service Level

Назад к [товарным анализам](../inventory.md)

## Назначение

Service Level оценивает качество исполнения спроса по данным о fulfillment и
stockout.

## Когда использовать

- когда нужен service KPI по каждому товару;
- когда важны и unit fulfillment, и cycle behavior.

## Когда не использовать

- когда есть только одна метрика и результат легко переинтерпретировать;
- когда в модели напрямую нужна вариативность lead time.

## Вход

- `DemandedUnits`
- `FulfilledUnits`
- `TotalCycles`
- `StockoutCycles`

## Выход

- `FillRate`
- `CycleServiceLevel`
- `StockoutRate`
- `UnfulfilledUnits`
- экспортируемый `ServiceLevel`

## Правила

- `FillRate = FulfilledUnits / DemandedUnits * 100`
- `CycleServiceLevel = (TotalCycles - StockoutCycles) / TotalCycles * 100`
- `StockoutRate = StockoutCycles / TotalCycles * 100`
- если `TotalCycles > 0`, `ServiceLevel` берется из `CycleServiceLevel`
- иначе `ServiceLevel` fallback’ится на `FillRate`

## Пример

```go
out, err := analytics.New().Inventory().ServiceLevel(ctx, servicelevel.Input{
    Items: []servicelevel.Item{
        {SKU: "A", Name: "Item A", DemandedUnits: 100, FulfilledUnits: 95, TotalCycles: 20, StockoutCycles: 2},
    },
})
```

## Интерпретация

Пакет сохраняет и fill-rate, и cycle-based view, чтобы было видно, не выглядит
ли товар сильным по одной метрике и слабым по другой.

## Частые ошибки

- считать, что экспортируемый `ServiceLevel` всегда равен fill rate;
- передавать cycle counts без единого определения цикла.

## Связанные анализы

- [Safety Stock](safetystock.md)
- [Reorder Point](reorderpoint.md)
