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
- implemented inventory analyses: `ABC`, `XYZ`, `ABC-XYZ`, `VED`, `FSN`, `HML`, `SDE`, `EOQ`, `Reorder Point`, `Safety Stock`.

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
├── inventory/          # Inventory facade
├── customer/           # Customer facade scaffold
├── facade.go           # Root analytics facade
├── go.mod
├── README.md           # Main project description
└── LICENSE
```

### Documentation

A detailed description of the abc package can be found in the file
[**abc/README.md**](abc/README.md)
