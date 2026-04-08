# API Conventions

This page documents implementation-wide rules that apply across packages.

## Public Entry Points

- Root facade: `analytics.New()`
- Inventory facade: `svc.Inventory()`
- Customer facade: `svc.Customer()`
- Package-level analyzers: each analysis package exposes `Analyze(ctx, input)`

## Common API Shape

- Each analysis accepts a dedicated `Input` struct.
- Each analysis returns a dedicated `Output` struct and `error`.
- Most analyses are stateless and safe to call through package functions.
- Composite analyses such as `ABC-XYZ` orchestrate lower-level packages instead
  of duplicating formulas.

## Context Handling

- All analyzers accept `context.Context`.
- Current implementations mostly use context for cancellation checks.
- If the context is canceled, the analyzer returns the context error.

## Sorting Conventions

- Most outputs are sorted by the primary business metric in descending order.
- Where relevant, ties are resolved by original input position through stable
  sort behavior.
- Some outputs preserve `OriginalIndex` so composite analyses can merge data
  safely.

## Defaulting Rules

- `ABC`: defaults to `80/95` thresholds.
- `XYZ`: defaults to coefficient of variation thresholds `10/25`.
- `VED`: defaults to criticality thresholds `70/40`.
- `FSN`: defaults to activity thresholds `70/1`.
- `Safety Stock`: defaults `ServiceFactor` to `1.65`.
- `Churn`: defaults to `30` days at risk and `90` days churned.
- `HML` and `SDE`: derive thresholds from data if explicit thresholds are not provided.

## Edge Case Policy

- Non-positive numeric business inputs usually produce `0` for derived metrics
  instead of hard validation errors.
- Empty datasets return empty results.
- Missing timestamps may be mapped to a default interpretation rather than
  rejected:
  `RFM` treats missing recency as very old,
  `Churn` treats missing last order as churn-threshold age.
- Some rates accept both ratio form (`0..1`) and percent form (`0..100`) where
  it is meaningful, for example in `CLV`.

## Deprecated API

- `abc.New()` and `(*ABC).Calculate(...)` are still present for backward
  compatibility.
- New integrations should prefer `abc.Analyze(...)` or the root/inventory
  facade.

## Compatibility Notes

- Current module path: `github.com/Sales-Analysis/abc-helper-lib`
- Current Go version in `go.mod`: `1.23.4`
- Version tags follow semantic versioning such as `v0.2.2`
