# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [v0.3.1] - 2026-04-09

### Changed

- Added row-level validation for core inventory analyses:
  `ABC`, `HML`, `SDE`, `XYZ`, `FSN`, `VED`, `Pareto`, and `GM/Contribution`.
- Tightened customer validation rules for `RFM`, `CLV`, and `Churn`.
- `RFM` and `Churn` now reject future `LastOrderAt` values instead of silently clipping them.
- `ABC` now returns zero shares for zero-revenue datasets instead of propagating `NaN` values.

### Documentation

- Clarified validation behavior in `docs/en/api.md` and `docs/ru/api.md`.
- Updated the documented validation matrix for commercial scalar inputs and timestamp handling.

### Testing

- Added regression coverage for the `ABC` zero-revenue edge case.
- Added negative tests for row-level validation across the updated inventory and customer analyses.

## [v0.3.0] - 2026-04-08

### Added

- Added preferred inventory-style `abc.Input.Items` input support while keeping `abc.Input.Products` as a legacy alias.
- Added `Summary` outputs for `ABC`, `XYZ`, `VED`, `FSN`, `HML`, `SDE`, and `ABC-XYZ`.
- Added `OriginalIndex` to the main `ABC` output so composite and direct consumers share the same indexing behavior.
- Added bilingual analysis selection guides in `docs/en/decision-guide.md` and `docs/ru/decision-guide.md`.
- Added bilingual `v1` API review and compatibility notes in `docs/en/v1-api-review.md` and `docs/ru/v1-api-review.md`.
- Added golden test infrastructure and golden fixtures for `ABC`, `XYZ`, `Pareto`, and `EOQ`.
- Added benchmark coverage for `ABC`, `XYZ`, `Pareto`, and `RFM`.
- Added fuzz tests for `EOQ`, `Reorder Point`, `Safety Stock`, and `Service Level`.
- Added property-based tests for `EOQ`, `Reorder Point`, `Safety Stock`, and `Service Level`.

### Changed

- Unified validation and error handling rules across threshold-based and bounded-input analyzers.
- Unified input and output conventions around `Items`, `Customers`, `Results`, `OriginalIndex`, and dataset-level `Summary` values.
- Expanded executable examples in `example_test.go` for key inventory and customer analyses.
- Shared float validation now rejects non-finite values such as `NaN` and `Inf`.

### Documentation

- Updated root README files and API docs to reflect the new input/output conventions and release documentation set.
- Linked the new decision guide and `v1` API review from the root documentation index and language-specific docs.

### Testing

- Added invalid-config and item-level validation coverage across updated analyzers.
- Established golden, benchmark, fuzz, and property-based test coverage as the new baseline quality layer.

## [v0.2.2] - 2026-04-06

### Documentation

- Updated the changelog with release history for `v0.1.0`, `v0.2.0`, and `v0.2.1`.
- Added GitHub compare links for released versions.

## [v0.2.1] - 2026-04-06

### Changed

- Aligned the `abc` package structure with the rest of the analyses.
- Consolidated usage documentation in the root README files.

### Documentation

- Added readable usage examples for all inventory and customer analyses in `README.md` and `README.ru.md`.
- Reformatted the root README usage section for better readability.

## [v0.2.0] - 2026-04-06

### Added

- Added the analytics library foundation with top-level `analytics`, `inventory`, and `customer` facades.
- Added inventory analyses: `ABC`, `XYZ`, `ABC-XYZ`, `VED`, `FSN`, `HML`, `SDE`, `EOQ`, `Reorder Point`, `Safety Stock`, `Pareto 80/20`, `GM/Contribution`, and `Service Level`.
- Added customer analyses: `RFM`, `CLV`, and `Churn/Retention`.
- Added stateless APIs to support composition across analyses.

### Changed

- Updated the root documentation to reflect the new library architecture.
- Kept the legacy `abc.New()` and `Calculate(...)` API for backward compatibility while moving the package to the new stateless flow.

### Deprecated

- Marked the legacy stateful `abc` API as deprecated.

### Removed

- Removed the unused generic `Analyzer` interface.

## [v0.1.0] - 2025-08-08

### Added

- Added the initial `ABC` analysis implementation.
- Added result building, structure refactoring, and initial tests.
- Added the first README and changelog documentation.

[Unreleased]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.3.1...HEAD
[v0.3.1]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.2...v0.3.0
[v0.2.2]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/Sales-Analysis/abc-helper-lib/releases/tag/v0.1.0
