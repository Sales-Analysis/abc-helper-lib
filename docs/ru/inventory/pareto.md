# Pareto 80/20

Назад к [товарным анализам](../inventory.md)

## Назначение

Pareto проверяет, дает ли малая доля товаров большую долю value.

## Когда использовать

- когда нужен быстрый взгляд на концентрацию value;
- когда нужно подтвердить или опровергнуть гипотезу `80/20`.

## Когда не использовать

- когда value не является правильной метрикой концентрации;
- когда нужен threshold-based grouping как в `ABC`, а не summary rule.

## Вход

- `Value`
- опциональные пороги `TopValueShare`, `TopItemShare`

## Выход

- `ValueShare` по товару
- `CumulativeValueShare`
- `CumulativeItemShare`
- флаги top-set
- summary с `ParetoPrincipleMet`

## Правила

- дефолтные пороги: `80%` value и `20%` товаров
- число top-items округляется вверх
- summary проверяет, достигает ли выбранная доля товаров целевой доли value

## Пример

```go
out, err := analytics.New().Inventory().Pareto(ctx, pareto.Input{
    Items: []pareto.Item{
        {SKU: "A", Name: "Item A", Value: 1000},
    },
})
```

## Интерпретация

Это анализ концентрации, а не прямая operational classification.

## Частые ошибки

- трактовать результат как замену `ABC`;
- забывать, что пороги настраиваются и не обязаны быть ровно `80/20`.

## Связанные анализы

- [ABC](abc.md)
- [GM/Contribution](gmcontribution.md)
