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

## Contracts

```text
proto/sandboxes/local/v1   local daemon API
proto/sandboxes/guest/v1   guest agent API
openapi/cloud/v1.yaml      cloud API placeholder
gen/go/                    generated Go, checked in
```

Generated code is produced with `buf` and must not be edited by hand.

## Development

Requires Go 1.27 and `buf` v1.50.0.

```sh
make generate
make test
make check
make lint
```
