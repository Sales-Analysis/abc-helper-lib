# Contributing

## Scope

This repository contains a Go library for inventory and customer analytics.
Contributions should keep the public API coherent, keep analysis packages
focused, and preserve existing behavior unless a breaking change is explicitly
planned.

## Local Setup

1. Install the Go version declared in `go.mod`.
2. Clone the repository.
3. Run `go test ./...` from the repository root.

## Design Guidelines

- Keep the root `analytics` facade thin.
- Keep domain facades in `inventory` and `customer` thin.
- Prefer one package per analysis.
- Prefer stateless `Analyze(...)` entry points for new analyses.
- Keep backward compatibility for exported APIs unless a breaking release is
  intentionally planned.
- Add or update tests for every behavior change.
- Update `README.md`, `README.ru.md`, and `CHANGELOG.md` when the change is
  user-visible.

## Pull Requests

- Keep each pull request focused on one change area.
- Include tests or explain why tests are not applicable.
- Document assumptions for thresholds, formulas, and domain-specific defaults.
- Use small, descriptive commits when possible.

## Release Process

1. Update `CHANGELOG.md`.
2. Verify the repository with `go test ./...`.
3. Create a semantic version tag such as `v0.2.2`.
4. Push the branch and the tag to `origin`.
5. Create a GitHub Release and use the matching `CHANGELOG.md` section as the
   release notes source.
