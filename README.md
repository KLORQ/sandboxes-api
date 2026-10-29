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

## Development

Requires Go 1.27.

```sh
make test
make check
```
