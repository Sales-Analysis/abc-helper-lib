# Клиентские анализы

Это index page для клиентских анализов, реализованных в `abc-helper-lib`.

## Содержание

- [RFM](customer/rfm.md)
- [CLV](customer/clv.md)
- [Churn / Retention](customer/churn.md)
- [Гид по выбору анализа](decision-guide.md)
- [Соглашения API](api.md)

## Сводная таблица

| Анализ | Пакет | Facade-метод | Основные входы | Ключевой результат |
| :--- | :--- | :--- | :--- | :--- |
| [`RFM`](customer/rfm.md) | `rfm` | `Customer().RFM` | last order, число заказов, monetary value | RFM score и сегмент |
| [`CLV`](customer/clv.md) | `clv` | `Customer().CLV` | выручка, заказы, retention, discount | прогнозный customer lifetime value |
| [`Churn/Retention`](customer/churn.md) | `churn` | `Customer().Churn` | дата последнего заказа | статус клиента и summary rates |

## Связанные технические заметки

- [Гид по выбору анализа](decision-guide.md)
- [Соглашения API](api.md)
- `RFM` решает задачу сегментации
- `CLV` решает задачу оценки ценности
- `Churn` решает задачу статуса удержания
