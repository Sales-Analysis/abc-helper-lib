# Обзор документации

`abc-helper-lib` — библиотека на Go для товарной и клиентской аналитики.

Текущая публичная точка входа — корневой facade `analytics`:

```go
svc := analytics.New()
inv := svc.Inventory()
cust := svc.Customer()
```

## Структура

- facade `inventory` объединяет товарные и supply-chain анализы;
- facade `customer` объединяет клиентские анализы;
- каждый анализ живет в отдельном пакете и отдает stateless API
  `Analyze(ctx, input)`.

## Доступные документы

- [Товарные анализы](inventory.md)
- [Клиентские анализы](customer.md)
- [Соглашения API и технические детали](api.md)

## Быстрая навигация

### Товарный блок

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

### Клиентский блок

- [RFM](customer/rfm.md)
- [CLV](customer/clv.md)
- [Churn / Retention](customer/churn.md)

## Соглашения по пакетам

- входные данные передаются через явные `Input`-структуры;
- результаты возвращаются через явные `Output`-структуры;
- составные анализы вроде `ABC-XYZ` собираются из нижележащих пакетов, а не
  прячутся в монолитный facade;
- legacy API `abc.New()` и `(*ABC).Calculate(...)` оставлен для обратной
  совместимости, но помечен как deprecated.

## Версионирование

- репозиторий использует semver-теги вида `v0.2.2`;
- история релизов ведется в корневом `CHANGELOG.md`.
