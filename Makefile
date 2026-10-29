BIN ?= bin/app

.PHONY: build test check fmt lint vuln generate generate-check clean

build:
	go build ./...

test:
	go test -race -count=1 ./...

check:
	test -z "$$(gofmt -l .)"
	go vet ./...

fmt:
	gofmt -w .

lint:
	golangci-lint config verify
	golangci-lint run ./...

vuln:
	govulncheck ./...

generate:
	buf lint
	buf generate

generate-check:
	buf generate
	git diff --exit-code

clean:
	rm -rf bin
