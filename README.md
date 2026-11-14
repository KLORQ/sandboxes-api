# sandboxes-api

<img src="docs/assets/rheic-shark.png" alt="Rheic shark logo" width="180">


## Import path migration

The canonical Go module path is now `github.com/rheic/sandboxes-api` (previously `github.com/quoiner/sandboxes-api`). Update imports and `go.mod` requirements when moving to a release that includes this change. The old module path is not kept as a compatibility alias.

Status: pre-alpha.

Public wire contracts for the sandboxes platform. This repository is the single
source of truth for protobuf definitions, the cloud OpenAPI document, generated
Go code, stable error codes, and version/capability metadata.

## Boundaries

- no business logic;
- no runtime, storage, or network implementations;
- generated code is never edited by hand;
- breaking contract changes before `1.0.0` follow the project's alpha policy:
  they may ship with an explicit forward-migration note and contract tests;
  after `1.0.0`, a breaking change requires a major version and migration note.

Policy migration (unreleased): `UpdatePolicy` whole-policy replacement is
rejected by the updated daemon because concurrent replacements can restore a
revoked rule. Clients must use `MutatePolicy` for atomic `allow`, `deny`, or
`rm` operations. Old clients fail closed and must be upgraded; new clients
against older daemons receive an unsupported RPC error. This is a breaking
security-contract change and requires the project's major-version release gate.

## Contracts

```text
proto/sandboxes/local/v1   local daemon API
proto/sandboxes/guest/v1   guest agent API
openapi/cloud/v1.yaml      cloud API placeholder
gen/go/                    generated Go, checked in
```

Generated code is produced with `buf` and must not be edited by hand.

`RegistryCredentialService.Login/Logout` transports registry credentials only
between `rsbx` and the local daemon. The daemon stores them in the host OS
credential vault; responses never include secret values. This API does not
claim that registry auth has been verified or that guest pulls are enabled.

`AccountService.Login/Logout` takes a trusted provider name and never returns
OAuth credentials, device codes, or PKCE values. The daemon resolves endpoints
and the public client registration from host configuration. Login is a bounded
unary call; logout reports only whether local account state was removed.

Local `Workspace` mounts identify a host `path`, whether it is `read_only`, and
an optional absolute `guest_path` destination. An empty `guest_path` preserves
the existing same-path mount behavior. Daemons must validate the host and guest
paths and reject conflicting destinations before creating the sandbox.

All remote MCP request headers must use `secret_headers` templates containing
`${placeholder}` references. The references cross the local RPC; the daemon
resolves values from the host vault when sending each request. The legacy
`headers` field is deprecated and supported daemons reject it because a custom
header name cannot prove that its value is non-secret.

## Development

Requires Go 1.27 and `buf` v1.50.0.

```sh
make generate
make test
make check
make lint
```
