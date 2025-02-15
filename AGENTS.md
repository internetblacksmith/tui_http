# AGENTS.md - Developer Guide for AI Coding Agents

## Build/Test Commands
- Build: `go build -o restless cmd/restless/main.go`
- Run: `./restless` or `go run cmd/restless/main.go`
- Test: `go test ./...` (run all tests) or `go test ./pkg/http` (single package)
- Lint: `go vet ./...` and `golangci-lint run` (if installed)
- Format: `go fmt ./...` or `gofmt -s -w .`
- Dependencies: `go mod tidy` to clean up dependencies

## Code Style Guidelines
### Go Specific
- Follow Go naming conventions (PascalCase for exports, camelCase for private)
- Use `gofmt` for consistent formatting
- Organize imports: standard library, external packages, internal packages
- Prefer explicit error handling over panics
- Use interfaces for testing and modularity

### Project Structure
- `cmd/` - Application entry points
- `pkg/` - Exportable packages (ui, http, models)  
- `internal/` - Private application code
- Keep UI components in `pkg/ui/`
- HTTP logic in `pkg/http/`
- Data models in `pkg/models/`

### TUI Development
- Use Bubble Tea framework patterns (Model, Update, View)
- Style definitions in separate files
- Keep components focused and composable
- Handle terminal resize events properly