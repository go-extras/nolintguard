# nolintguard

[![CI/CD Pipeline](https://github.com/go-extras/nolintguard/actions/workflows/ci.yml/badge.svg)](https://github.com/go-extras/nolintguard/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-extras/nolintguard)](https://goreportcard.com/report/github.com/go-extras/nolintguard)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`nolintguard` is a static analysis tool designed to enforce organizational policy around the usage of `//nolint` directives in Go source code. It is intended to be used as a **custom linter for golangci-lint**.

## Purpose

The tool helps enforce best practices for suppression directives by:

- Requiring `#nosec` instead of `//nolint:gosec` for security-related suppressions
- Requiring native revive directives instead of `//nolint:revive` for style suppressions
- Optionally restricting or forbidding arbitrary `//nolint` directives

This ensures that security suppressions are explicit and tool-appropriate, stylistic suppressions use native mechanisms, and arbitrary suppressions are disallowed unless explicitly permitted.

## Installation

### As a golangci-lint plugin

```bash
go get github.com/go-extras/nolintguard
```

### As a standalone tool

```bash
# Install latest release
go install github.com/go-extras/nolintguard/cmd/nolintguard@latest

# Install specific version
go install github.com/go-extras/nolintguard/cmd/nolintguard@v1.0.0

# Or download pre-built binaries from GitHub Releases
# https://github.com/go-extras/nolintguard/releases

# Or build locally
make build

# Or install locally
make install
```

## Usage

### Standalone Mode

Run the linter directly on your code:

```bash
# Analyze current package
nolintguard .

# Analyze all packages recursively
nolintguard ./...

# With justification requirement
nolintguard -require-justification ./...

# With forbidden linters
nolintguard -forbidden-linters=staticcheck,unused ./...

# Combine flags
nolintguard -require-justification -forbidden-linters=staticcheck ./...
```

**Available flags:**
- `-require-justification` - Require `#nosec`, `//gosec:`, and `//revive:` directives to include justification
- `-forbidden-linters=<list>` - Comma-separated list of linters to forbid in `//nolint` directives

### With golangci-lint

Add `nolintguard` to your `.golangci.yml`:

```yaml
linters:
  enable:
    - nolintguard

linters-settings:
  nolintguard:
    # Require security/style suppression directives to include justification
    require-justification: true  # default: false
    # Comma-separated list of linters to forbid in //nolint directives
    forbidden-linters: "staticcheck,unused"  # default: ""
```

## Rules

### 1. Forbidden: `//nolint:gosec`

Any usage of `//nolint:gosec` is **forbidden**. Use gosec's native suppression directives instead.

Gosec supports two official formats for suppressing warnings:

**Format 1: `#nosec`**
```go
// #nosec G401 -- Using MD5 for non-cryptographic checksums only
hash := md5.New()

// Multiple rules
// #nosec G201 G202 G203 -- SQL concatenation reviewed and safe
query := "SELECT * FROM " + tableName

// All rules
// #nosec -- Verified safe: input is sanitized upstream
result := exec.Command("ls", userInput).Run()
```

**Format 2: `//gosec:ignore` or `//gosec:disable`**
```go
//gosec:ignore G401 -- Using MD5 for non-cryptographic checksums only
hash := md5.New()
```

Or:
```go
//gosec:disable G101 -- This is a false positive
const apiKey = "placeholder"
```

**Bad:**
```go
//nolint:gosec
hash := md5.New()
```

**Error message:**
```
nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead
```

### 2. Forbidden: `//nolint:revive`

Any usage of `//nolint:revive` is **forbidden**. Use native revive suppression directives instead.

**Bad:**
```go
//nolint:revive
return errors.New("test")
```

**Good:**
```go
//revive:disable-next-line
return errors.New("test")

//revive:disable Until the code is stable
type MyType struct { ... }

//revive:disable:exported This is internal only
func helperFunc() { ... }
```

**Error message:**
```
nolintguard: //nolint:revive is forbidden; use native revive directives instead
```

### 3. Optional: Forbid Specific Linters

You can forbid specific linters in `//nolint` directives by listing them in `forbidden-linters`.

**Configuration:**
```yaml
linters-settings:
  nolintguard:
    forbidden-linters: "staticcheck,unused,gosimple"
```

**Example:**
```go
//nolint:staticcheck  // Error: staticcheck is forbidden
//nolint:errcheck     // OK: errcheck is not in the forbidden list
```

**Error message:**
```
nolintguard: //nolint:staticcheck is forbidden
```

### 4. Optional: Require Justification for Suppressions

When `require-justification` is enabled, security and style suppression directives (`#nosec`, `//gosec:`, `//revive:`) **must** include a justification.

**Configuration:**
```yaml
linters-settings:
  nolintguard:
    require-justification: true
```

**Bad:**
```go
// #nosec G401
hash := md5.New()

//gosec:ignore G101
const apiKey = "placeholder"

//revive:disable
type myType struct { ... }
```

**Good:**
```go
// Gosec uses -- marker for justification
// #nosec G401 -- Using MD5 for non-cryptographic checksums only
hash := md5.New()

//gosec:ignore G101 -- This is a placeholder, not a real credential
const apiKey = "placeholder"

// Revive uses space-separated justification (no -- marker)
//revive:disable Until the code is stable
type myType struct { ... }

//revive:disable:exported This is internal only
func helperFunc() { ... }
```

**Error messages:**
```
nolintguard: #nosec directive must include justification (-- reason)
nolintguard: //gosec: directive must include justification (-- reason)
nolintguard: //revive: directive must include justification (reason)
```

## Configuration Options

| Field                   | Type   | Default | Description                                                                      |
|------------------------|--------|---------|----------------------------------------------------------------------------------|
| `require-justification` | bool   | `false` | Require `#nosec`, `//gosec:`, and `//revive:` directives to include justification |
| `forbidden-linters`     | string | `""`    | Comma-separated list of linters to forbid in `//nolint` directives              |

## Examples

### Multi-linter directives

The tool correctly handles multi-linter suppressions:

```go
//nolint:gosec,errcheck  // Error: gosec is forbidden
//nolint:revive,errcheck // Error: revive is forbidden
//nolint:gosec,revive    // Two errors: both gosec and revive are forbidden
```

### Whitespace variations

The tool handles various whitespace formats:

```go
//nolint:gosec       // Detected
//nolint: gosec      // Detected
// nolint:gosec      // Detected
/* nolint:gosec */   // Detected
```

## Non-Goals

The linter explicitly does **not**:

- Validate correctness of `#nosec` usage
- Re-run or replace gosec or revive
- Attempt autofix or rewriting
- Infer intent from context

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## Development

The project includes a Makefile for common development tasks:

```bash
# Build the standalone binary
make build

# Install to GOPATH/bin
make install

# Run tests
make test

# Run tests with coverage
go test -coverprofile=coverage.out -covermode=atomic ./...
go tool cover -html=coverage.out

# Run linter
make lint

# Run formatters (auto-fix)
make fmt

# Clean build artifacts
make clean

# Show all available targets
make help
```

Or use Go commands directly:

```bash
# Run tests
go test ./...

# Build all packages
go build ./...

# Build standalone binary
go build -o bin/nolintguard ./cmd/nolintguard

# Test GoReleaser configuration
goreleaser check

# Build snapshot (local testing)
goreleaser build --snapshot --clean --single-target
```

### Releases

Releases are automated using GoReleaser:

- **Pull Requests**: Snapshot builds are created as artifacts for testing
- **Tagged Releases**: Production releases are published to GitHub Releases when a tag is pushed

To create a new release:

```bash
# Tag the release
git tag -a v1.0.0 -m "Release v1.0.0"

# Push the tag
git push origin v1.0.0
```

The CI/CD pipeline will automatically:
- Build binaries for all supported platforms (Linux, macOS, Windows, FreeBSD)
- Create archives (tar.gz for Unix, zip for Windows)
- Generate checksums
- Publish to GitHub Releases
