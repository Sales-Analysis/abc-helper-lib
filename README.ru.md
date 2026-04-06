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

### Совместимость

- Версия Go: `1.23.4` (из `go.mod`)
- Версионирование: semver-теги вида `v0.2.2`

### Установка

```bash
go get github.com/Sales-Analysis/abc-helper-lib@latest
```

### Использование

Создай корневой facade один раз, а затем вызывай товарные и клиентские анализы
через доменные facade:

```go
ctx := context.Background()
svc := analytics.New()
inv := svc.Inventory()
cust := svc.Customer()
```

Товарные анализы:

```go
_, _ = inv.ABC(ctx, abc.Input{
    Products: []abc.Product{
        {SKU: "A", Name: "Item A", Quantity: 10, Price: 100},
    },
})

_, _ = inv.XYZ(ctx, xyz.Input{
    Items: []xyz.Item{
        {SKU: "A", Name: "Item A", Demands: []float64{10, 12, 9}},
    },
})

_, _ = inv.ABCXYZ(ctx, abcxyz.Input{
    Items: []abcxyz.Item{
        {SKU: "A", Name: "Item A", Quantity: 10, Price: 100, Demands: []float64{10, 12, 9}},
    },
})

_, _ = inv.VED(ctx, ved.Input{
    Items: []ved.Item{
        {SKU: "A", Name: "Item A", CriticalityScore: 85},
    },
})

_, _ = inv.FSN(ctx, fsn.Input{
    Items: []fsn.Item{
        {SKU: "A", Name: "Item A", Movements: []float64{3, 0, 2, 1}},
    },
})

_, _ = inv.HML(ctx, hml.Input{
    Items: []hml.Item{
        {SKU: "A", Name: "Item A", UnitCost: 250},
    },
})

_, _ = inv.SDE(ctx, sde.Input{
    Items: []sde.Item{
        {SKU: "A", Name: "Item A", LeadTimeDays: 45},
    },
})

_, _ = inv.EOQ(ctx, eoq.Input{
    Items: []eoq.Item{
        {SKU: "A", Name: "Item A", AnnualDemand: 1200, OrderingCost: 50, HoldingCost: 2},
    },
})

_, _ = inv.ReorderPoint(ctx, reorderpoint.Input{
    Items: []reorderpoint.Item{
        {SKU: "A", Name: "Item A", AverageDemandPerPeriod: 20, LeadTimePeriods: 4, SafetyStock: 33},
    },
})

_, _ = inv.SafetyStock(ctx, safetystock.Input{
    Items: []safetystock.Item{
        {SKU: "A", Name: "Item A", DemandStdDev: 10, LeadTimePeriods: 4, ServiceFactor: 1.65},
    },
})

_, _ = inv.Pareto(ctx, pareto.Input{
    Items: []pareto.Item{
        {SKU: "A", Name: "Item A", Value: 1000},
    },
})

_, _ = inv.GMContribution(ctx, gmcontribution.Input{
    Items: []gmcontribution.Item{
        {SKU: "A", Name: "Item A", Revenue: 1000, COGS: 600, VariableCost: 700, FixedCost: 100},
    },
})

_, _ = inv.ServiceLevel(ctx, servicelevel.Input{
    Items: []servicelevel.Item{
        {SKU: "A", Name: "Item A", DemandedUnits: 100, FulfilledUnits: 95, TotalCycles: 20, StockoutCycles: 2},
    },
})
```

Клиентские анализы:

```go
_, _ = cust.RFM(ctx, rfm.Input{
    AnalysisTime: time.Now(),
    Customers: []rfm.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -10), Orders: 5, MonetaryValue: 1200},
    },
})

_, _ = cust.CLV(ctx, clv.Input{
    Customers: []clv.Customer{
        {CustomerID: "C1", Name: "Customer 1", Revenue: 1200, Orders: 6, PeriodsObserved: 3, GrossMarginRate: 0.4, RetentionRate: 0.8, DiscountRate: 0.1, AcquisitionCost: 50},
    },
})

_, _ = cust.Churn(ctx, churn.Input{
    AnalysisTime: time.Now(),
    Customers: []churn.Customer{
        {CustomerID: "C1", Name: "Customer 1", LastOrderAt: time.Now().AddDate(0, 0, -45)},
    },
})
```

Минимальный набор импортов для этих примеров:

```go
import (
    "context"
    "time"

    analytics "github.com/Sales-Analysis/abc-helper-lib"
    "github.com/Sales-Analysis/abc-helper-lib/abc"
    "github.com/Sales-Analysis/abc-helper-lib/abcxyz"
    "github.com/Sales-Analysis/abc-helper-lib/churn"
    "github.com/Sales-Analysis/abc-helper-lib/clv"
    "github.com/Sales-Analysis/abc-helper-lib/eoq"
    "github.com/Sales-Analysis/abc-helper-lib/fsn"
    "github.com/Sales-Analysis/abc-helper-lib/gmcontribution"
    "github.com/Sales-Analysis/abc-helper-lib/hml"
    "github.com/Sales-Analysis/abc-helper-lib/pareto"
    "github.com/Sales-Analysis/abc-helper-lib/reorderpoint"
    "github.com/Sales-Analysis/abc-helper-lib/rfm"
    "github.com/Sales-Analysis/abc-helper-lib/safetystock"
    "github.com/Sales-Analysis/abc-helper-lib/sde"
    "github.com/Sales-Analysis/abc-helper-lib/servicelevel"
    "github.com/Sales-Analysis/abc-helper-lib/ved"
    "github.com/Sales-Analysis/abc-helper-lib/xyz"
)
```

### Статус анализов

Все анализы ниже реализованы в `v0.2.2`.

| Анализ | Домен | Пакет | Facade-метод | Статус |
| :--- | :--- | :--- | :--- | :--- |
| `ABC` | Товарный | `abc` | `Inventory().ABC` | Реализован |
| `XYZ` | Товарный | `xyz` | `Inventory().XYZ` | Реализован |
| `ABC-XYZ` | Товарный | `abcxyz` | `Inventory().ABCXYZ` | Реализован |
| `VED` | Товарный | `ved` | `Inventory().VED` | Реализован |
| `FSN` | Товарный | `fsn` | `Inventory().FSN` | Реализован |
| `HML` | Товарный | `hml` | `Inventory().HML` | Реализован |
| `SDE` | Товарный | `sde` | `Inventory().SDE` | Реализован |
| `EOQ` | Товарный | `eoq` | `Inventory().EOQ` | Реализован |
| `Reorder Point` | Товарный | `reorderpoint` | `Inventory().ReorderPoint` | Реализован |
| `Safety Stock` | Товарный | `safetystock` | `Inventory().SafetyStock` | Реализован |
| `Pareto 80/20` | Товарный | `pareto` | `Inventory().Pareto` | Реализован |
| `GM/Contribution` | Товарный | `gmcontribution` | `Inventory().GMContribution` | Реализован |
| `Service Level` | Товарный | `servicelevel` | `Inventory().ServiceLevel` | Реализован |
| `RFM` | Клиентский | `rfm` | `Customer().RFM` | Реализован |
| `CLV` | Клиентский | `clv` | `Customer().CLV` | Реализован |
| `Churn/Retention` | Клиентский | `churn` | `Customer().Churn` | Реализован |

Legacy-заметка: `abc.New()` и `(*ABC).Calculate(...)` оставлены для обратной совместимости, но помечены как deprecated.

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

### Документы проекта

- [**CHANGELOG.md**](CHANGELOG.md) с историей релизов
- [**CONTRIBUTING.md**](CONTRIBUTING.md) с процессом разработки и релизов
- [**SECURITY.md**](SECURITY.md) с правилами сообщения об уязвимостях
- [**LICENSE**](LICENSE) с условиями лицензии
