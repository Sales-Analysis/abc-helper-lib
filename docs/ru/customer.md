# Клиентские анализы

На этой странице собраны клиентские анализы, которые сейчас реализованы в
`abc-helper-lib`.

## Сводная таблица

| Анализ | Пакет | Facade-метод | Основные входы | Ключевой результат |
| :--- | :--- | :--- | :--- | :--- |
| `RFM` | `rfm` | `Customer().RFM` | last order, число заказов, monetary value | RFM score и сегмент |
| `CLV` | `clv` | `Customer().CLV` | выручка, заказы, retention, discount | прогнозный customer lifetime value |
| `Churn/Retention` | `churn` | `Customer().Churn` | дата последнего заказа | статус клиента и summary rates |

## RFM

- Назначение: сегментация клиентов по recency, frequency и monetary value.
- Вход:
  `LastOrderAt`, `Orders`, `MonetaryValue`, опционально `AnalysisTime`.
- Скоринг:
  каждая ось ранжируется по шкале `1..5` на основе относительного места
  клиента в выборке.
- Текущие сегменты:
  `Champions`, `Loyal`, `Potential Loyalist`, `At Risk`, `Hibernating`,
  `Promising`.
- Спецслучай:
  отсутствие даты последнего заказа трактуется как очень старая recency.

## CLV

- Назначение: оценка прогнозного customer lifetime value.
- Вход:
  `Revenue`, `Orders`, `PeriodsObserved`, `GrossMarginRate`,
  `RetentionRate`, `DiscountRate`, `AcquisitionCost`.
- Производные метрики:
  average order value и purchase frequency.
- Формула:
  `PredictedCLV = (AOV * PurchaseFrequency * GrossMarginRate * (RetentionRate / (1 + DiscountRate - RetentionRate))) - AcquisitionCost`.
- Нормализация ratios:
  если `GrossMarginRate`, `RetentionRate` или `DiscountRate` переданы в
  диапазоне `0..100`, они переводятся в `0..1`.
- Если знаменатель неположительный или ключевые бизнес-входы отсутствуют,
  реализация возвращает `-AcquisitionCost`.

## Churn / Retention

- Назначение: классификация клиентов по окну неактивности и сводка по
  retention.
- Вход:
  `LastOrderAt`, опционально `AnalysisTime`, опционально свои пороги.
- Дефолтные пороги:
  `AtRisk = 30 days`, `Churned = 90 days`.
- Результат по клиенту:
  `DaysSinceLastOrder`, `Status`, `IsRetained`, `IsChurned`.
- Summary:
  общее число клиентов, число active, at-risk и churned, retention rate и
  churn rate.
