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
- Inventory analyses use `Items` as the preferred collection field in `Input`.
- Customer analyses use `Customers` as the collection field in `Input`.
- `Results` is the standard top-level collection field in `Output`.
- `abc.Input.Products` remains supported as a legacy alias, but new code should use `abc.Input.Items`.
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
- If an analyzer reorders results relative to the input, result items should
  preserve `OriginalIndex`.
- Classification-style analyzers expose `Summary` when group counts or top-level
  totals materially help interpretation.

## Defaulting Rules

- `ABC`: defaults to `80/95` thresholds.
- `XYZ`: defaults to coefficient of variation thresholds `10/25`.
- `VED`: defaults to criticality thresholds `70/40`.
- `FSN`: defaults to activity thresholds `70/1`.
- `Safety Stock`: defaults `ServiceFactor` to `1.65`.
- `Churn`: defaults to `30` days at risk and `90` days churned.
- `HML` and `SDE`: derive thresholds from data if explicit thresholds are not provided.

## Validation Policy

- Invalid configuration returns `error`.
- Threshold ordering mistakes now return `error` instead of silently resetting to defaults.
- Auto-defaulting remains available only for omitted optional configuration.
- For `HML` and `SDE`, explicit threshold overrides must be complete and ordered correctly.
- Composite analyses propagate validation errors from their underlying analyzers.
- Structurally impossible bounded inputs now return `error`, for example:
  negative order counts, negative lead times, invalid ratio fields, and
  impossible service-level cycle relationships.

## Validation Behavior Matrix

| Case | Behavior | Notes |
| :--- | :--- | :--- |
| Optional configuration is omitted | Default | Examples: `ABC` thresholds, `XYZ` thresholds, `Safety Stock` service factor |
| Threshold value is out of range | `error` | Examples: percent > `100`, negative threshold, invalid ratio |
| Threshold ordering is invalid | `error` | Examples: `B <= A`, `Y <= X`, `ChurnDays <= AtRiskDays` |
| Explicit threshold override is incomplete where full override is required | `error` | Applies to `HML` and `SDE` |
| Bounded numeric input is negative | `error` | Examples: negative order counts, lead times, demand deviation |
| Float input is `NaN` or `Inf` | `error` | Shared validation rejects non-finite values |
| Cross-field bounded relationship is impossible | `error` | Examples: `FulfilledUnits > DemandedUnits`, `StockoutCycles > TotalCycles` |
| Formula input is non-positive but structurally allowed | Derived metric becomes `0` | Examples: `EOQ`, `Safety Stock`, `Reorder Point` helper math |
| Commercial scalar input is used directly by ranking or formula logic | Usually `error` when negative or non-finite | Examples: `Pareto.Value`, `GM/Contribution` money fields, `RFM.MonetaryValue`, `CLV.Revenue` |
| Dataset is empty | Empty result | `Results` stays empty and summaries remain zeroed where applicable |
| Timestamp is omitted | Default interpretation | `RFM` treats it as very old, `Churn` maps it to churn-threshold age |
| Timestamp is in the future | `error` | `RFM` and `Churn` reject future `LastOrderAt` values |
| Rate is supplied as ratio or percent where supported | Normalize | `CLV` accepts both `0..1` and `0..100` forms |

## Edge Case Policy

- Non-positive numeric business inputs usually produce `0` for derived metrics
  instead of hard validation errors.
- Commercial scalar fields that directly affect ranking or financial formulas
  are increasingly validated as finite and non-negative.
- Empty datasets return empty results.
- Missing timestamps may be mapped to a default interpretation rather than
  rejected:
  `RFM` treats missing recency as very old,
  `Churn` treats missing last order as churn-threshold age.
- Future timestamps are rejected where the analyzer contract depends on
  historical recency, currently in `RFM` and `Churn`.
- Some rates accept both ratio form (`0..1`) and percent form (`0..100`) where
  it is meaningful, for example in `CLV`.

## Error Shape

- The current validation contract returns regular `error` values.
- Validation errors currently use the textual prefix `invalid input: ...`.
- Typed validation errors are intentionally not part of the stable public API yet.
- New integrations should depend on analyzer behavior and documented rules, not on concrete error types.

## Deprecated API

- `abc.New()` and `(*ABC).Calculate(...)` are still present for backward
  compatibility.
- New integrations should prefer `abc.Analyze(...)` or the root/inventory
  facade.

## Compatibility Notes

- Current module path: `github.com/Sales-Analysis/abc-helper-lib`
- Current Go version in `go.mod`: `1.23.4`
- Version tags follow semantic versioning such as `v0.3.0`
