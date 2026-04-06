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

### Installation

```bash
go get gitlab.com/username/abc-helper-lib@latest
```

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

### Documentation

A detailed description of the abc package can be found in the file
[**abc/README.md**](abc/README.md)
