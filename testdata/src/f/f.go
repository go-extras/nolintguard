package f

// Edge cases and corner cases

import (
	"crypto/md5"
)

// Test case: Duplicate linters in nolint directive
func duplicateLinters() {
	//nolint:gosec,gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead" "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Trailing comma
func trailingComma() {
	//nolint:gosec, // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Leading comma (should be ignored as invalid format)
func leadingComma() {
	//nolint:,gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Multiple consecutive commas
func multipleCommas() {
	//nolint:gosec,,errcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Tab character in linter list
func tabCharacter() {
	//nolint:	gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Lots of whitespace
func lotsOfWhitespace() {
	//nolint:   gosec   ,   errcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: nolint with only colon (no linters)
func onlyColon() {
	//nolint:
	x := 1
	_ = x
}

// Test case: nolint with colon and spaces only
func colonAndSpaces() {
	//nolint:   
	x := 1
	_ = x
}

// Test case: Very long linter list
func longLinterList() {
	//nolint:errcheck,ineffassign,staticcheck,unused,deadcode,varcheck,structcheck,gosec,typecheck,bodyclose // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: gosec at the end of a long list
func gosecAtEnd() {
	//nolint:errcheck,ineffassign,staticcheck,unused,gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Multiple forbidden linters in one directive
func multipleForbidden() {
	//nolint:gosec,revive,errcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead" "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	h := md5.New()
	_ = h
}

// Test case: nolint with inline comment containing colon
func inlineCommentWithColon() {
	//nolint:gosec // TODO: fix this later // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Block comment with nolint
func blockCommentNolint() {
	/* nolint:gosec */ // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: Block comment with nolint at start
func blockCommentNolintAtStart() {
	/* nolint:gosec */ // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: #nosec with multiple spaces before rule
func nosecMultipleSpaces() {
	// #nosec    G401 -- Justified
	h := md5.New()
	_ = h
}

// Test case: #nosec with tab before rule
func nosecTab() {
	// #nosec	G401 -- Justified
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with extra spaces
func gosecIgnoreSpaces() {
	//gosec:ignore    G401   --   Justified
	h := md5.New()
	_ = h
}

// Test case: Plain nolint with spaces
func plainNolintSpaces() {
	// nolint
	x := 1
	_ = x
}

// Test case: nolint mentioned in regular comment (should be ignored)
func nolintInComment() {
	// This function uses nolint directives
	x := 1
	_ = x
}

// Test case: Comment that starts with "nolint" but isn't a directive
func nolintNotDirective() {
	// nolinting is important for code quality
	x := 1
	_ = x
}

// Test case: Empty linter name in list
func emptyLinterName() {
	//nolint:gosec,,, // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

