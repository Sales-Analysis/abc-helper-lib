# ABC Helper Library

| English | Русский |
| :---: | :---: |
| [**README.md**](README.md) | [**README.ru.md**](README.ru.md) |

---

## Русский

### Описание

`abc-helper-lib` — библиотека на Go для товарной и клиентской аналитики.

Текущий фундамент включает:

- тонкий корневой facade для orchestration;
- доменное разделение на `inventory` и `customer`;
- реализованные товарные анализы: `ABC`, `XYZ`, `ABC-XYZ`, `VED`, `FSN`, `HML`, `SDE`, `EOQ`, `Reorder Point`, `Safety Stock`, `Pareto`, `GM/Contribution`, `Service Level`;
- реализованные клиентские анализы: `RFM`, `CLV`, `Churn`.

### Установка

```bash
go get gitlab.com/username/abc-helper-lib@latest
```

### Структура проекта

```
abc-helper-lib/
│
├── abc/                # ABC-анализ
├── xyz/                # XYZ-анализ
├── abcxyz/             # Комбинированный ABC-XYZ
├── ved/                # VED-анализ
├── fsn/                # FSN-анализ
├── hml/                # HML-анализ
├── sde/                # SDE-анализ
├── eoq/                # Расчет EOQ
├── reorderpoint/       # Расчет точки заказа
├── safetystock/        # Расчет страхового запаса
├── pareto/             # Анализ Pareto 80/20
├── gmcontribution/     # Анализ gross margin и contribution
├── servicelevel/       # Анализ service level
├── inventory/          # Facade товарной аналитики
├── rfm/                # RFM-анализ
├── clv/                # Анализ customer lifetime value
├── churn/              # Анализ churn и retention
├── customer/           # Facade клиентской аналитики
├── facade.go           # Корневой analytics facade
├── go.mod
├── README.md           # Общее описание проекта
└── LICENSE
```

### Документация

Подробное описание пакета abc находится в файле
[**abc/README.ru.md**](abc/README.ru.md)
