# Changelog

## Unreleased

- Add optional `Workspace.guest_path` for mounting a validated host path at an
  explicit guest destination. Empty values retain same-path mount behavior.
- Advance the local RPC contract version to 0.0.25 so clients can reject older
  daemons that would ignore a non-empty guest destination.

## v0.0.24 — 2026-11-09

- Deprecate static remote MCP request headers; supported daemons reject them
  because sensitivity cannot be inferred from custom header names.
- Require host-vault-backed `secret_headers` templates for all remote MCP headers.
- Clarify that pre-1.0 breaking contract changes follow the documented alpha migration policy.

## v0.0.23 — 2026-11-08

- Bind M7 pre-tag approval to the exact reviewed commit.
- Add the `secret_headers` MCP registration field for host-vault placeholder
  templates, keeping credential values out of the registration RPC and state.

## v0.0.20 — 2026-11-07

- Add optional `url` and non-secret `headers` to `MCPRegisterServerRequest`
  for the remote streamable-HTTP MCP transport. Exactly one of `command` or
  `url` must be set; remote hosts are constrained by the daemon's fail-closed
  allowlist.

## v0.0.19 — 2026-11-07

- Add presence-only refresh metadata (`refresh_present`, `refresh_coordinated`)
  to `AccountService.Status`. No token material, expiry values or validity
  claims are exposed.

## v0.0.18 — 2026-11-07

- Add the local `SandboxService.MCPCall` contract for invoking a named MCP tool
  in a sandbox. Daemon-side registry lookup, policy enforcement, bounded
  payload validation, and approval remain required before execution.
- Add a redacted `SandboxService.ListMCPServers` inventory contract for local
  `mcp ls`; it exposes gateway/transport and usability flags, never credentials
  or OAuth identity details, and must not start servers or auth flows.
- Add explicit stdio registration/removal contracts and one-use registration approval challenges.
- Add one-use MCP call approval challenges. The approval RPC carries only the challenge ID; it is not proof of consent. The daemon must obtain approval through a trusted host approval provider before execution.
- Document that persisted stdio command arguments are non-secret configuration; credentials require a host credential broker.

## v0.0.16 — 2026-11-06

### Added

- Add an optional host-side dynamic secret source to service-secret set/create RPCs, including a 1Password/AWS reference or shell command, refresh policy, verification controls, and captured provider account/profile selectors.

### Compatibility

- Existing value-based secret requests are unchanged. Older daemons ignore the additive field and reject the empty value; SDK clients must check the daemon API version before sending a dynamic source.

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

