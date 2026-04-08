# GM/Contribution

Назад к [товарным анализам](../inventory.md)

## Назначение

GM/Contribution считает gross margin и contribution метрики из выручки и
компонентов затрат.

## Когда использовать

- когда ассортиментные решения требуют смотреть не только на revenue;
- когда важна структура variable и fixed cost.

## Когда не использовать

- когда доступен только top-line sales value;
- когда cost allocation слишком шумный для интерпретации на уровне товара.

## Вход

- `Revenue`
- `COGS`
- `VariableCost`
- `FixedCost`

## Выход

- `GrossMargin`
- `GrossMarginRate`
- `ContributionMargin`
- `ContributionRate`
- `NetContribution`

## Правила

- `GrossMargin = Revenue - COGS`
- `ContributionMargin = Revenue - VariableCost`
- `NetContribution = ContributionMargin - FixedCost`
- rate-метрики выражаются в процентах от выручки

## Пример

```go
out, err := analytics.New().Inventory().GMContribution(ctx, gmcontribution.Input{
    Items: []gmcontribution.Item{
        {SKU: "A", Name: "Item A", Revenue: 1000, COGS: 600, VariableCost: 700, FixedCost: 100},
    },
})
```

## Интерпретация

Этот анализ отделяет gross margin от contribution logic, что полезно, когда
variable и fixed cost дают разные бизнес-сигналы.

## Частые ошибки

- путать `COGS` и `VariableCost`;
- после расчета richer margin metrics все равно сравнивать товары только по revenue.

## Связанные анализы

- [Pareto 80/20](pareto.md)
- [ABC](abc.md)
