.PHONY: build install test lint clean help

# Build the standalone binary
build:
	@echo "Building nolintguard..."
	@mkdir -p bin
	@go build -o bin/nolintguard ./cmd/nolintguard
	@echo "Binary built at bin/nolintguard"

# Install the binary to GOPATH/bin
install:
	@echo "Installing nolintguard..."
	@go install ./cmd/nolintguard
	@echo "Installed to $(shell go env GOPATH)/bin/nolintguard"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run linter
lint:
	@echo "Running golangci-lint..."
	@golangci-lint run ./...

# Run linter and fix issues
lint-fix:
	@echo "Running golangci-lint with autofix..."
	@golangci-lint run ./... --fix

# Run formatters
fmt:
	@echo "Running formatters..."
	@golangci-lint run --fix ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@echo "Clean complete"

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the standalone binary to bin/nolintguard"
	@echo "  install  - Install the binary to GOPATH/bin"
	@echo "  test     - Run all tests"
	@echo "  lint     - Run golangci-lint"
	@echo "  fmt      - Run formatters (auto-fix)"
	@echo "  clean    - Remove build artifacts"
	@echo "  help     - Show this help message"

