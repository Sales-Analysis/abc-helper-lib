# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

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

[Unreleased]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.2...HEAD
[v0.2.2]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/Sales-Analysis/abc-helper-lib/compare/v0.1.0...v0.2.0
[v0.1.0]: https://github.com/Sales-Analysis/abc-helper-lib/releases/tag/v0.1.0
