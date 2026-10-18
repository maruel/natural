# natural

Run `make help` for the build and check targets. Run `make fix`, `make verify`,
and `make test` before committing. CI also runs the platform and compatibility
checks in `.github/workflows/test.yml`. Install versioned hooks with
`make git-hooks`.

`make test` runs coverage and benchmark smoke checks and writes ignored
`coverage.txt`. Use `make test-race` for race checks.

CI shares checkout and locked Python setup in one Ubuntu check job. Runtime
tests use the existing Go version before static checks switch to Go 1.27.

Tooling is pinned independently of the runtime Go version in go.mod. Static
checks use Go 1.27 in CI to support golangci-lint v2.14.0.

The shared script index is generated with
`uv run --locked python scripts/update_agents_file_index.py`; use `--check` to validate it.

Python scripts and Ruff run through `uv run --locked`; `make tools` syncs the
locked development environment. `pyproject.toml` and `uv.lock` define the tools,
and `.editorconfig` defines shared editor defaults.
