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
	"flag"
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/go-extras/nolintguard"
)

// Build information. Populated at build-time via ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Check for version flag before singlechecker processes flags.
	for _, arg := range os.Args[1:] {
		if arg == "-version" || arg == "--version" || arg == "-V" {
			fmt.Printf("nolintguard version %s (commit: %s, built: %s)\n", version, commit, date)
			os.Exit(0)
		}
	}

	// Add custom version flag.
	flag.Bool("version", false, "print version and exit")

	singlechecker.Main(nolintguard.Analyzer)
}
