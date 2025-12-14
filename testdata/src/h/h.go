package h

// Test combination of require-justification and forbidden-linters flags
// This file is tested with both flags enabled:
// - require-justification=true
// - forbidden-linters=staticcheck,unused

import (
	"crypto/md5"
)

// Test case: #nosec without justification (should fail)
func nosecNoJustification() {
	// #nosec G401 // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification (should pass)
func nosecWithJustification() {
	// #nosec G401 -- Using MD5 for checksums only
	h := md5.New()
	_ = h
}

// Test case: forbidden linter staticcheck (should fail)
func forbiddenStaticcheck() {
	//nolint:staticcheck // want "nolintguard: //nolint:staticcheck is forbidden"
	x := 1
	_ = x
}

// Test case: forbidden linter unused (should fail)
func forbiddenUnused() {
	//nolint:unused // want "nolintguard: //nolint:unused is forbidden"
	x := 1
	_ = x
}

// Test case: allowed linter errcheck (should pass)
func allowedErrcheck() {
	//nolint:errcheck
	x := 1
	_ = x
}

// Test case: gosec in nolint (always forbidden)
func gosecInNolint() {
	//nolint:gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: revive in nolint (always forbidden)
func reviveInNolint() {
	//nolint:revive // want "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	x := 1
	_ = x
}

// Test case: Multiple forbidden linters including gosec
func multipleForbiddenWithGosec() {
	//nolint:gosec,staticcheck,unused // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead" "nolintguard: //nolint:staticcheck is forbidden" "nolintguard: //nolint:unused is forbidden"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore without justification (should fail)
func gosecIgnoreNoJustification() {
	//gosec:ignore G401 // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with justification (should pass)
func gosecIgnoreWithJustification() {
	//gosec:ignore G401 -- Safe usage
	h := md5.New()
	_ = h
}

// Test case: revive directive without justification (should fail)
func reviveNoJustification() {
	//revive:disable // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive directive with justification (should pass)
func reviveWithJustification() {
	//revive:disable Temporary workaround
	x := 1
	_ = x
}

// Test case: Plain nolint (always allowed)
func plainNolint() {
	//nolint
	x := 1
	_ = x
}

// Test case: Mix of allowed and forbidden linters
func mixedLinters() {
	//nolint:errcheck,staticcheck,ineffassign // want "nolintguard: //nolint:staticcheck is forbidden"
	x := 1
	_ = x
}

// Test case: All three types of forbidden (gosec, revive, custom)
func allThreeForbidden() {
	//nolint:gosec,revive,staticcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead" "nolintguard: //nolint:revive is forbidden; use native revive directives instead" "nolintguard: //nolint:staticcheck is forbidden"
	h := md5.New()
	_ = h
}

