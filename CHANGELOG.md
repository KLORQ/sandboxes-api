# Changelog

## v0.0.15 — 2026-11-05

### Added

- Add sandbox-scoped OAuth metadata to local Create using nested messages aligned with the Docker Kits v2 declaration: token endpoint, resource hosts, sentinels, credential file, response fields, passthrough, and `skipIfEnv`.

### Compatibility

- This contract carries kit metadata only. Daemons must fail closed when custom OAuth interception is unavailable; the contract alone does not imply runtime support.

## v0.0.12 — 2026-11-03

### Added

- Add the account device-login stream and credential-presence status RPCs.
- Publish a CycloneDX module SBOM as a CI artifact.

### Changed

- Normalize the Go module and generated protobuf package paths to the lowercase `github.com/quoiner/sandboxes-api` namespace.
- Update `golang.org/x/net` to v0.62.0, including fixes for GO-2026-6860 and GO-2026-6866.

