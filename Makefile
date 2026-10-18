# Copyright 2026 Marc-Antoine Ruel. All rights reserved.
# Use of this source code is governed under the Apache License, Version 2.0
# that can be found in the LICENSE file.

# Repository development targets and pinned static analysis tools.
.DEFAULT_GOAL := help
.PHONY: build fix git-hooks help test test-race tools verify

GOLANGCI_LINT := go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.14.0
ACTIONLINT := go run github.com/rhysd/actionlint/cmd/actionlint@v1.7.12
ADDLICENSE := go run github.com/google/addlicense@v1.2.0

help:
	@printf '  %-14s - %s\n' \
		'make build' 'Build all Go packages' \
		'make fix' 'Apply formatting and safe lint fixes' \
		'make verify' 'Run static checks' \
		'make test' 'Run Go tests with coverage and benchmark smoke checks' \
		'make test-race' 'Run Go tests with the race detector' \
		'make git-hooks' 'Install versioned Git hooks' \
		'make tools' 'Sync locked development tools'

fix: tools
	@uv run --locked ruff check --fix scripts
	@uv run --locked ruff format scripts
	@uv run --locked python scripts/update_agents_file_index.py
	@$(GOLANGCI_LINT) fmt
	@$(GOLANGCI_LINT) run --show-stats=false --fix ./...

verify: tools
	@uv run --locked ruff check scripts
	@uv run --locked ruff format --check scripts
	@$(ACTIONLINT)
	@$(GOLANGCI_LINT) fmt --diff
	@$(GOLANGCI_LINT) run --show-stats=false ./...
	@go vet -unsafeptr=false ./...
	@go test -run '^$$' ./...
	@$(ADDLICENSE) -ignore '.venv/**' -ignore '.ruff_cache/**' -ignore 'third_party/**' -ignore 'vendor/**' -check .
	@uv run --locked python scripts/lint_binaries.py
	@uv run --locked python scripts/update_agents_file_index.py --check
	@git ls-files -z | xargs -0 go run github.com/client9/misspell/cmd/misspell@v0.3.4 -error
	@go mod tidy -diff

build:
	@go build ./...

test:
	@go test -timeout=120s -covermode=count -coverprofile coverage.txt -bench=. -benchtime=1x ./...

test-race:
	@go test -timeout=120s -race -bench=. -benchtime=1x ./...

git-hooks:
	@./scripts/install-git-hooks.sh

tools:
	@uv sync --locked --quiet
