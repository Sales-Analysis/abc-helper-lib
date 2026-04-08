# Гид По Выбору Анализа

Используй эту страницу, когда главный вопрос не "как работает пакет?", а
"какой анализ лучше взять первым?"

## Товарные Вопросы

| Бизнес-вопрос | С чего начать | Что часто комбинируют |
| :--- | :--- | :--- |
| Какие SKU дают основную выручку? | [ABC](inventory/abc.md) | [Pareto 80/20](inventory/pareto.md), [ABC-XYZ](inventory/abcxyz.md) |
| Правда ли, что малое число товаров дает большую часть value? | [Pareto 80/20](inventory/pareto.md) | [ABC](inventory/abc.md) |
| Какие товары имеют стабильный, а какие волатильный спрос? | [XYZ](inventory/xyz.md) | [FSN](inventory/fsn.md), [ABC-XYZ](inventory/abcxyz.md) |
| Нужны одновременно важность по value и стабильность спроса. | [ABC-XYZ](inventory/abcxyz.md) | [Safety Stock](inventory/safetystock.md), [Reorder Point](inventory/reorderpoint.md) |
| Какие товары критичны операционно? | [VED](inventory/ved.md) | [SDE](inventory/sde.md), [HML](inventory/hml.md) |
| Какие товары fast-moving, slow-moving или non-moving? | [FSN](inventory/fsn.md) | [XYZ](inventory/xyz.md), [Pareto 80/20](inventory/pareto.md) |
| Какие товары дороги в расчете на единицу? | [HML](inventory/hml.md) | [ABC](inventory/abc.md), [VED](inventory/ved.md) |
| Какие товары сложны в снабжении? | [SDE](inventory/sde.md) | [Safety Stock](inventory/safetystock.md), [VED](inventory/ved.md) |
| Какой размер заказа экономически оптимален? | [EOQ](inventory/eoq.md) | [Reorder Point](inventory/reorderpoint.md) |
| Когда нужно размещать следующий заказ? | [Reorder Point](inventory/reorderpoint.md) | [Safety Stock](inventory/safetystock.md), [Service Level](inventory/servicelevel.md) |
| Какой страховой буфер нужен? | [Safety Stock](inventory/safetystock.md) | [Reorder Point](inventory/reorderpoint.md), [Service Level](inventory/servicelevel.md) |
| Насколько хорошо мы сейчас исполняем спрос? | [Service Level](inventory/servicelevel.md) | [Safety Stock](inventory/safetystock.md), [Reorder Point](inventory/reorderpoint.md) |
| Какие товары дают gross margin или contribution? | [GM/Contribution](inventory/gmcontribution.md) | [ABC](inventory/abc.md), [Pareto 80/20](inventory/pareto.md) |

## Клиентские Вопросы

| Бизнес-вопрос | С чего начать | Что часто комбинируют |
| :--- | :--- | :--- |
| Какие клиенты сейчас наиболее вовлечены? | [RFM](customer/rfm.md) | [CLV](customer/clv.md), [Churn / Retention](customer/churn.md) |
| Какие клиенты наиболее ценны экономически? | [CLV](customer/clv.md) | [RFM](customer/rfm.md), [Churn / Retention](customer/churn.md) |
| Какие клиенты удержаны, под риском или уже churned? | [Churn / Retention](customer/churn.md) | [RFM](customer/rfm.md), [CLV](customer/clv.md) |

## Рекомендуемые Комбинации

- `ABC + XYZ + Safety Stock`: выделить важные и нестабильные товары и подобрать им защитный буфер.
- `VED + SDE + Safety Stock`: защищать критичные позиции, которые еще и трудно пополнять.
- `ABC + Pareto`: проверить, действительно ли малая доля SKU дает основную выручку.
- `RFM + CLV`: объединить вовлеченность и экономическую ценность перед маркетинговыми решениями.
- `Churn + CLV`: направлять retention-усилия туда, где потенциальная потеря дороже всего.

## Быстрые Эвристики

- Выбирай [ABC](inventory/abc.md), когда язык задачи про "value", "выручку" или "долю вклада".
- Выбирай [XYZ](inventory/xyz.md), когда задача про "стабильность", "вариативность" или "прогнозируемость".
- Выбирай [FSN](inventory/fsn.md), когда речь о частоте движения, а не о статистической вариативности.
- Выбирай [VED](inventory/ved.md), когда бизнес-критичность важнее чистой экономики.
- Используй [EOQ](inventory/eoq.md), [Safety Stock](inventory/safetystock.md) и [Reorder Point](inventory/reorderpoint.md) как операционную тройку, а не как изолированные формулы.

## Частые Ошибки Выбора

- Использовать [ABC](inventory/abc.md) для вывода о стабильности спроса. Обычно это задача [XYZ](inventory/xyz.md).
- Использовать [XYZ](inventory/xyz.md) как proxy для важности бизнеса. Обычно это задача [ABC](inventory/abc.md) или [Pareto 80/20](inventory/pareto.md).
- Использовать [EOQ](inventory/eoq.md) без проверки reorder timing и service assumptions.
- Использовать [CLV](customer/clv.md) отдельно от модели текущего состояния, например [RFM](customer/rfm.md) или [Churn](customer/churn.md).
