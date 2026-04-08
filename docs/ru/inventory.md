# Товарные анализы

Это index page для товарных и supply-chain анализов, реализованных в
`abc-helper-lib`.

## Содержание

- [ABC](inventory/abc.md)
- [XYZ](inventory/xyz.md)
- [ABC-XYZ](inventory/abcxyz.md)
- [VED](inventory/ved.md)
- [FSN](inventory/fsn.md)
- [HML](inventory/hml.md)
- [SDE](inventory/sde.md)
- [EOQ](inventory/eoq.md)
- [Reorder Point](inventory/reorderpoint.md)
- [Safety Stock](inventory/safetystock.md)
- [Pareto 80/20](inventory/pareto.md)
- [GM/Contribution](inventory/gmcontribution.md)
- [Service Level](inventory/servicelevel.md)
- [Соглашения API](api.md)

## Сводная таблица

| Анализ | Пакет | Facade-метод | Основные входы | Ключевой результат |
| :--- | :--- | :--- | :--- | :--- |
| [`ABC`](inventory/abc.md) | `abc` | `Inventory().ABC` | количество, цена | доля выручки и группа |
| [`XYZ`](inventory/xyz.md) | `xyz` | `Inventory().XYZ` | ряд спроса | группа по вариативности |
| [`ABC-XYZ`](inventory/abcxyz.md) | `abcxyz` | `Inventory().ABCXYZ` | стоимость продаж и ряд спроса | комбинированная группа |
| [`VED`](inventory/ved.md) | `ved` | `Inventory().VED` | score критичности | группа критичности |
| [`FSN`](inventory/fsn.md) | `fsn` | `Inventory().FSN` | история движений | fast/slow/non-moving группа |
| [`HML`](inventory/hml.md) | `hml` | `Inventory().HML` | unit cost | группа по стоимости |
| [`SDE`](inventory/sde.md) | `sde` | `Inventory().SDE` | lead time в днях | группа по сложности снабжения |
| [`EOQ`](inventory/eoq.md) | `eoq` | `Inventory().EOQ` | спрос, ordering cost, holding cost | оптимальный размер заказа |
| [`Reorder Point`](inventory/reorderpoint.md) | `reorderpoint` | `Inventory().ReorderPoint` | средний спрос, lead time, safety stock | точка заказа |
| [`Safety Stock`](inventory/safetystock.md) | `safetystock` | `Inventory().SafetyStock` | отклонение спроса, lead time, service factor | страховой запас |
| [`Pareto 80/20`](inventory/pareto.md) | `pareto` | `Inventory().Pareto` | value по товару | накопленная доля и проверка 80/20 |
| [`GM/Contribution`](inventory/gmcontribution.md) | `gmcontribution` | `Inventory().GMContribution` | выручка и компоненты затрат | маржинальные метрики |
| [`Service Level`](inventory/servicelevel.md) | `servicelevel` | `Inventory().ServiceLevel` | исполнение спроса и stockout cycles | метрики уровня сервиса |

## Связанные технические заметки

- [Соглашения API](api.md)
- `ABC`, `XYZ` и `ABC-XYZ` связаны через `OriginalIndex`
- `Safety Stock` и `Reorder Point` обычно читаются вместе как пара расчетов
