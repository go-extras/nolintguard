// Command nolintguard is a standalone runner for the nolintguard analyzer.
//
// It can be used to run the linter independently without golangci-lint.
//
// Usage:
//
//	nolintguard [flags] [packages]
//
// Examples:
//
//	# Analyze current package
//	nolintguard .
//
//	# Analyze specific packages
//	nolintguard ./...
//
//	# With justification requirement
//	nolintguard -require-justification ./...
//
//	# With forbidden linters
//	nolintguard -forbidden-linters=staticcheck,unused ./...
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/go-extras/nolintguard"
)

func main() {
	singlechecker.Main(nolintguard.Analyzer)
}
