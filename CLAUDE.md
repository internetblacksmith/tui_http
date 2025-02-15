# Restless

Terminal-based HTTP client built with Go and Bubble Tea for making REST API calls from the command line.

## Build Commands

```bash
make test    # go test ./...
make lint    # go vet ./... + golangci-lint
make build   # go build -o restless cmd/restless/main.go
make run     # go run cmd/restless/main.go
```

## Critical Rules

- Pin dependencies to exact versions (e.g., `"github.com/charmbracelet/bubbletea v1.3.6"`)
- Keep docs updated with every code change
- Keep Makefile updated - add new tasks as project evolves
- Follow Bubble Tea patterns: Model/Update/View separation, styles in dedicated files
- Keep HTTP client logic isolated in `pkg/http/`, never import UI packages from it

## Detailed Guides

| Topic | Guide |
|-------|-------|
| Agent setup | [AGENTS.md](AGENTS.md) |
