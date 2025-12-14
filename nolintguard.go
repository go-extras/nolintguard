// Package nolintguard implements a static analysis tool that enforces
// organizational policy around the usage of //nolint directives in Go source code.
//
// The analyzer detects and reports violations of nolint usage policies, including:
//   - Forbidden usage of //nolint:gosec (requires #nosec or //gosec: directives instead)
//   - Forbidden usage of //nolint:revive (requires native revive directives)
//   - Optional restriction of specific linters via forbidden-linters configuration
//   - Optional justification requirements for security/style suppression directives
//
// This linter is designed to be used as a custom linter for golangci-lint.
package nolintguard

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer creates a new instance of the nolintguard analyzer.
// This function is useful for testing with different flag configurations.
func NewAnalyzer() *analysis.Analyzer {
	var (
		requireJustification bool
		forbiddenLinters     string // comma-separated list
	)

	a := &analysis.Analyzer{
		Name:             "nolintguard",
		Doc:              "enforces project policy for nolint directives",
		Run:              makeRun(&requireJustification, &forbiddenLinters),
		RunDespiteErrors: true,
	}

	a.Flags.BoolVar(&requireJustification, "require-justification", false, "require security suppression directives (#nosec, //gosec:, //revive:) to include justification")
	a.Flags.StringVar(&forbiddenLinters, "forbidden-linters", "", "comma-separated list of forbidden nolint linters (e.g., 'staticcheck,unused')")

	return a
}

// Analyzer is the nolintguard analyzer that enforces project policy
// for nolint directives.
var Analyzer = NewAnalyzer()

// Config holds the configuration options for the nolintguard analyzer.
type Config struct {
	// RequireJustification, when true, requires security suppression directives
	// (#nosec, //gosec:, //revive:) to include a justification comment.
	RequireJustification bool

	// ForbiddenLinters is a map of linter names that should be forbidden
	// in //nolint directives (e.g., staticcheck, unused).
	ForbiddenLinters map[string]bool
}

const (
	gosecMessage             = "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	reviveMessage            = "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	nosecNoJustificationMsg  = "nolintguard: #nosec directive must include justification (-- reason)"
	gosecNoJustificationMsg  = "nolintguard: //gosec: directive must include justification (-- reason)"
	reviveNoJustificationMsg = "nolintguard: //revive: directive must include justification (reason)"
)

// makeRun creates a run function with closure over flag pointers.
// This allows each analyzer instance to have its own configuration.
func makeRun(requireJustification *bool, forbiddenLinters *string) func(*analysis.Pass) (any, error) {
	return func(pass *analysis.Pass) (any, error) {
		// Parse forbidden linters from comma-separated string
		forbiddenMap := make(map[string]bool)
		if *forbiddenLinters != "" {
			for _, linter := range strings.Split(*forbiddenLinters, ",") {
				linter = strings.TrimSpace(linter)
				if linter != "" {
					forbiddenMap[linter] = true
				}
			}
		}

		config := Config{
			RequireJustification: *requireJustification,
			ForbiddenLinters:     forbiddenMap,
		}

		for _, file := range pass.Files {
			inspectComments(pass, file, config)
		}

		return nil, nil
	}
}

// inspectComments examines all comments in a file for nolint directive violations.
func inspectComments(pass *analysis.Pass, file *ast.File, config Config) {
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			checkComment(pass, comment, config)
		}
	}
}

// checkComment analyzes a single comment for nolint directive violations.
func checkComment(pass *analysis.Pass, comment *ast.Comment, config Config) {
	text := comment.Text

	// Remove comment markers
	text = strings.TrimPrefix(text, "//")
	text = strings.TrimPrefix(text, "/*")
	text = strings.TrimSuffix(text, "*/")
	text = strings.TrimSpace(text)

	// Strip any inline comment after the directive (e.g., "// want ..." in tests)
	if idx := strings.Index(text, "//"); idx != -1 {
		text = strings.TrimSpace(text[:idx])
	}

	// Check for #nosec directive
	if strings.HasPrefix(text, "#nosec") {
		if config.RequireJustification {
			checkNosecJustification(pass, comment, text)
		}
		return
	}

	// Check for //gosec: directive
	if strings.HasPrefix(text, "gosec:") {
		if config.RequireJustification {
			checkGosecDirectiveJustification(pass, comment, text)
		}
		return
	}

	// Check for //revive: directive
	if strings.HasPrefix(text, "revive:") {
		if config.RequireJustification {
			checkReviveJustification(pass, comment, text)
		}
		return
	}

	// Check if this is a nolint directive
	if !strings.HasPrefix(text, "nolint") {
		return
	}

	// Extract the part after "nolint"
	remainder := strings.TrimPrefix(text, "nolint")

	// Handle plain //nolint without arguments
	if remainder == "" {
		return // Plain nolint is allowed
	}

	// Parse linter names from //nolint:linter1,linter2,...
	// Skip any whitespace or colon
	remainder = strings.TrimSpace(remainder)
	if !strings.HasPrefix(remainder, ":") {
		// Not a proper nolint directive format
		return
	}

	linters := parseLinters(remainder[1:]) // Skip the ':'

	// Check each linter for policy violations
	for _, linter := range linters {
		switch linter {
		case "gosec":
			// Always forbidden - must use #nosec
			pass.Reportf(comment.Pos(), "%s", gosecMessage)
		case "revive":
			// Always forbidden - must use native revive directives
			pass.Reportf(comment.Pos(), "%s", reviveMessage)
		default:
			// Check if this linter is in the forbidden list
			if config.ForbiddenLinters[linter] {
				pass.Reportf(comment.Pos(), "nolintguard: //nolint:%s is forbidden", linter)
			}
		}
	}
}

// parseLinters extracts individual linter names from a comma-separated list.
// It handles whitespace and returns a slice of trimmed linter names.
// It stops parsing at the first occurrence of // (inline comment).
func parseLinters(lintersText string) []string {
	if lintersText == "" {
		return nil
	}

	// Stop at the first // (inline comment)
	if idx := strings.Index(lintersText, "//"); idx != -1 {
		lintersText = lintersText[:idx]
	}

	parts := strings.Split(lintersText, ",")
	linters := make([]string, 0, len(parts))

	for _, part := range parts {
		linter := strings.TrimSpace(part)
		if linter != "" {
			linters = append(linters, linter)
		}
	}

	return linters
}

// checkNosecJustification verifies that a #nosec directive includes a justification.
// Format: #nosec [rules] -- justification.
func checkNosecJustification(pass *analysis.Pass, comment *ast.Comment, text string) {
	if !hasGosecJustification(text) {
		pass.Reportf(comment.Pos(), "%s", nosecNoJustificationMsg)
	}
}

// checkGosecDirectiveJustification verifies that a //gosec: directive includes a justification.
// Format: //gosec:ignore [rules] -- justification or //gosec:disable [rules] -- justification.
func checkGosecDirectiveJustification(pass *analysis.Pass, comment *ast.Comment, text string) {
	if !hasGosecJustification(text) {
		pass.Reportf(comment.Pos(), "%s", gosecNoJustificationMsg)
	}
}

// checkReviveJustification verifies that a //revive: directive includes a justification.
// Format: //revive:disable justification (space-separated, not --).
func checkReviveJustification(pass *analysis.Pass, comment *ast.Comment, text string) {
	if !hasReviveJustification(text) {
		pass.Reportf(comment.Pos(), "%s", reviveNoJustificationMsg)
	}
}

// hasGosecJustification checks if a directive contains a justification marker (--).
// A justification should be in the format: "-- reason for suppression".
func hasGosecJustification(text string) bool {
	// Look for the justification marker "--"
	idx := strings.Index(text, "--")
	if idx == -1 {
		return false
	}

	// Check that there's actual content after the "--"
	justification := strings.TrimSpace(text[idx+2:])
	return len(justification) > 0
}

// hasReviveJustification checks if a revive directive contains a justification.
// Revive format: //revive:disable justification (space-separated).
// Example: //revive:disable Until the code is stable.
func hasReviveJustification(text string) bool {
	// revive directives have format: revive:disable[:rule-name] [justification]
	// Find the first space after the rule specification

	// Skip past "revive:"
	text = strings.TrimPrefix(text, "revive:")

	// Common patterns: disable, disable-line, disable-next-line, enable
	for _, prefix := range []string{"disable", "enable"} {
		if strings.HasPrefix(text, prefix) {
			text = strings.TrimPrefix(text, prefix)

			// Check for optional rule name: disable:rule-name
			if strings.HasPrefix(text, ":") {
				// Find the space after the rule name
				idx := strings.Index(text, " ")
				if idx == -1 {
					return false // No justification
				}
				text = text[idx:]
			}

			// Now check if there's justification text
			text = strings.TrimSpace(text)
			return len(text) > 0
		}
	}

	return false
}
