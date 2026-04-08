# ABC Helper Library

| English | Русский |
| :---: | :---: |
| [**README.md**](README.md) | [**README.ru.md**](README.ru.md) |

---

## English

### Overview

`abc-helper-lib` is a Go library for inventory and customer analytics.

The current foundation includes:

- a thin root facade for orchestration;
- domain separation into `inventory` and `customer`;
- implemented inventory analyses: `ABC`, `XYZ`, `ABC-XYZ`, `VED`, `FSN`, `HML`, `SDE`, `EOQ`, `Reorder Point`, `Safety Stock`, `Pareto`, `GM/Contribution`, `Service Level`;
- implemented customer analyses: `RFM`, `CLV`, `Churn`.

### Compatibility

- Go version: `1.23.4` (from `go.mod`)
- Versioning: semantic version tags such as `v0.2.2`

### Installation

```bash
go get github.com/Sales-Analysis/abc-helper-lib@latest
```

### Usage

Create the root facade once and then call inventory and customer analyses through
their domain facades:

```go
ctx := context.Background()
svc := analytics.New()
inv := svc.Inventory()
cust := svc.Customer()
```

Inventory analyses:

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

Customer analyses:

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

Minimal import set for these examples:

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

### Analysis Status

All analyses below are implemented in `v0.2.2`.

| Analysis | Domain | Package | Facade Method | Status |
| :--- | :--- | :--- | :--- | :--- |
| `ABC` | Inventory | `abc` | `Inventory().ABC` | Implemented |
| `XYZ` | Inventory | `xyz` | `Inventory().XYZ` | Implemented |
| `ABC-XYZ` | Inventory | `abcxyz` | `Inventory().ABCXYZ` | Implemented |
| `VED` | Inventory | `ved` | `Inventory().VED` | Implemented |
| `FSN` | Inventory | `fsn` | `Inventory().FSN` | Implemented |
| `HML` | Inventory | `hml` | `Inventory().HML` | Implemented |
| `SDE` | Inventory | `sde` | `Inventory().SDE` | Implemented |
| `EOQ` | Inventory | `eoq` | `Inventory().EOQ` | Implemented |
| `Reorder Point` | Inventory | `reorderpoint` | `Inventory().ReorderPoint` | Implemented |
| `Safety Stock` | Inventory | `safetystock` | `Inventory().SafetyStock` | Implemented |
| `Pareto 80/20` | Inventory | `pareto` | `Inventory().Pareto` | Implemented |
| `GM/Contribution` | Inventory | `gmcontribution` | `Inventory().GMContribution` | Implemented |
| `Service Level` | Inventory | `servicelevel` | `Inventory().ServiceLevel` | Implemented |
| `RFM` | Customer | `rfm` | `Customer().RFM` | Implemented |
| `CLV` | Customer | `clv` | `Customer().CLV` | Implemented |
| `Churn/Retention` | Customer | `churn` | `Customer().Churn` | Implemented |

Legacy note: `abc.New()` and `(*ABC).Calculate(...)` remain available for backward compatibility, but they are deprecated.

### Project Structure

```
abc-helper-lib/
│
├── abc/                # ABC analysis
├── xyz/                # XYZ analysis
├── abcxyz/             # Combined ABC-XYZ analysis
├── ved/                # VED analysis
├── fsn/                # FSN analysis
├── hml/                # HML analysis
├── sde/                # SDE analysis
├── eoq/                # EOQ calculation
├── reorderpoint/       # Reorder point calculation
├── safetystock/        # Safety stock calculation
├── pareto/             # Pareto 80/20 analysis
├── gmcontribution/     # Gross margin and contribution analysis
├── servicelevel/       # Service level analysis
├── inventory/          # Inventory facade
├── rfm/                # RFM analysis
├── clv/                # Customer lifetime value analysis
├── churn/              # Churn and retention analysis
├── customer/           # Customer facade
├── facade.go           # Root analytics facade
├── go.mod
├── README.md           # Main project description
└── LICENSE
```

### Project Docs

- [**docs/README.md**](docs/README.md) for detailed English and Russian analysis docs
- [**CHANGELOG.md**](CHANGELOG.md) for release history
- [**CONTRIBUTING.md**](CONTRIBUTING.md) for development and release workflow
- [**SECURITY.md**](SECURITY.md) for vulnerability reporting guidance
- [**LICENSE**](LICENSE) for licensing terms
