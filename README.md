# sandboxes-api

Status: pre-alpha.

Public wire contracts for the sandboxes platform. This repository is the single
source of truth for protobuf definitions, the cloud OpenAPI document, generated
Go code, stable error codes, and version/capability metadata.

## Boundaries

- no business logic;
- no runtime, storage, or network implementations;
- generated code is never edited by hand;
- breaking contract changes require a major version and a migration note.

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
between `ksbx` and the local daemon. The daemon stores them in the host OS
credential vault; responses never include secret values. This API does not
claim that registry auth has been verified or that guest pulls are enabled.

`AccountService.Login/Logout` takes a trusted provider name and never returns
OAuth credentials, device codes, or PKCE values. The daemon resolves endpoints
and the public client registration from host configuration. Login is a bounded
unary call; logout reports only whether local account state was removed.

## Development

Requires Go 1.27 and `buf` v1.50.0.

```sh
make generate
make test
make check
make lint
```
