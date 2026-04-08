# V1 API Review

This page records the current public API posture of `abc-helper-lib` and the
remaining work required before a `v1.0.0` stability commitment.

## Current Stability Level

- The module is still in the `v0.x` phase.
- Public APIs are usable and documented, but they are not yet frozen.
- Minor releases may still refine exported structs, defaults, validation rules,
  and output conventions when the changes are documented in `CHANGELOG.md`.

## What Is Already In Good Shape

- Root orchestration through `analytics.New()`.
- Domain facades through `Inventory()` and `Customer()`.
- Stateless package-level entry points through `Analyze(ctx, input)`.
- Explicit `Input` and `Output` structs across packages.
- Consistent use of `Items` for inventory inputs and `Customers` for customer inputs.
- `OriginalIndex` preservation when outputs are reordered.
- `Summary` on classification-style analyzers and other analyses with clear dataset-level aggregates.

## Deprecated Surface

- `abc.New()`
- `(*abc.ABC).Calculate(...)`
- `abc.Input.Products` in favor of `abc.Input.Items`

These APIs remain available for backward compatibility, but new integrations
should not build on them.

## Compatibility Policy During v0.x

- Patch releases should not intentionally break exported APIs.
- Minor releases may still add fields, tighten validation, or clarify output
  ordering when those changes improve consistency and are documented.
- Deprecated APIs should not be removed silently.
- Every compatibility-affecting change should be reflected in `CHANGELOG.md`
  and the bilingual docs under `docs/`.

## Open Decisions Before v1.0.0

1. Decide whether the legacy stateful `abc` API will be removed or retained as a thin compatibility layer.
2. Decide whether `abc.TotalRevenue` remains as a long-term convenience alias beside `abc.Output.Summary`.
3. Decide how far `Summary` should go across non-classification outputs such as `EOQ`, `CLV`, `GM/Contribution`, and `Service Level`.
4. Decide whether generic `"invalid input: ..."` errors are sufficient or whether the module needs typed validation errors.
5. Finish robustness work:
   fuzz tests, property-based tests, and broader edge-case datasets.
6. Automate release flow from version tag to GitHub Release notes.
7. Publish an explicit backward compatibility statement for the post-`v1.0.0` line.

## Recommended v1 Readiness Checklist

- Finalize the fate of deprecated `abc` entry points.
- Freeze naming and output conventions package-wide.
- Expand edge-case coverage beyond nominal unit tests.
- Keep golden tests and benchmarks for representative analyzers green.
- Add compatibility and migration notes for any last pre-v1 API changes.
- Tag the first release only after the public API review items above are closed.
