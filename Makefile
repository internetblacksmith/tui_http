.PHONY: build run test clean fmt lint deps build-all menu help list

.DEFAULT_GOAL := menu

menu:
	@echo "╔══════════════════════════════════════════════════════╗"
	@echo "║           Restless - Command Menu                    ║"
	@echo "╚══════════════════════════════════════════════════════╝"
	@echo ""
	@echo "  === Development ==="
	@echo "  1) Build binary"
	@echo "  2) Run application"
	@echo ""
	@echo "  === Testing ==="
	@echo "  3) Run all tests"
	@echo "  4) Run linter"
	@echo "  5) Format code"
	@echo ""
	@echo "  === Build ==="
	@echo "  6) Build all platforms"
	@echo "  7) Clean artifacts"
	@echo "  8) Tidy dependencies"
	@echo ""
	@read -p "Enter choice: " choice; \
	case $$choice in \
		1) $(MAKE) build ;; \
		2) $(MAKE) run ;; \
		3) $(MAKE) test ;; \
		4) $(MAKE) lint ;; \
		5) $(MAKE) fmt ;; \
		6) $(MAKE) build-all ;; \
		7) $(MAKE) clean ;; \
		8) $(MAKE) deps ;; \
		*) echo "Invalid choice" ;; \
	esac

# Build the application
build:
	go build -o restless cmd/restless/main.go

# Run the application
run:
	go run cmd/restless/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f restless restless-*

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	go vet ./...
	@which golangci-lint > /dev/null && golangci-lint run || echo "golangci-lint not installed, skipping"

# Install dependencies
deps:
	go mod tidy

# Build for different platforms
build-all:
	GOOS=linux GOARCH=amd64 go build -o restless-linux-amd64 cmd/restless/main.go
	GOOS=darwin GOARCH=amd64 go build -o restless-darwin-amd64 cmd/restless/main.go
	GOOS=windows GOARCH=amd64 go build -o restless-windows-amd64.exe cmd/restless/main.go

help:
	@echo "Available commands:"
	@echo "  make build      - Build the restless binary"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run all tests"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make fmt        - Format Go source code"
	@echo "  make lint       - Run linter and vet"
	@echo "  make deps       - Tidy module dependencies"
	@echo "  make build-all  - Cross-compile for linux, darwin, windows"
	@echo "  make help       - Show this help"
	@echo "  make list       - Quick command reference"

list:
	@echo "build run test clean fmt lint deps build-all"
