# Documentation Overview

`abc-helper-lib` is a Go library for inventory and customer analytics.

The current public entry point is the root `analytics` facade:

```go
svc := analytics.New()
inv := svc.Inventory()
cust := svc.Customer()
```

## Structure

- `inventory` facade groups inventory and supply-chain analyses.
- `customer` facade groups customer lifecycle analyses.
- Each analysis lives in its own package and exposes a stateless
  `Analyze(ctx, input)` API.

## Available Documentation

- [Decision guide](decision-guide.md)
- [Inventory analyses](inventory.md)
- [Customer analyses](customer.md)
- [API conventions and technical notes](api.md)
- [v1 API review](v1-api-review.md)

## Quick Navigation

### Inventory

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

### Customer

- [RFM](customer/rfm.md)
- [CLV](customer/clv.md)
- [Churn / Retention](customer/churn.md)

## Package Conventions

- Inputs are passed via explicit `Input` structs.
- Outputs are returned via explicit `Output` structs.
- Composite analyses such as `ABC-XYZ` are assembled from lower-level
  analyses instead of hiding all logic in a monolithic facade.
- The legacy `abc.New()` and `(*ABC).Calculate(...)` API remains available for
  backward compatibility, but it is deprecated.

## Versioning

- The repository uses semantic version tags such as `v0.3.0`.
- Release history is recorded in the root `CHANGELOG.md`.
