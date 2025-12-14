package b

import (
	"crypto/md5"
)

// Test case: #nosec without justification (will fail when require-justification is enabled)
func nosecNoJustification() {
	// #nosec G401 // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification (OK)
func nosecWithJustification() {
	// #nosec G401 -- Using MD5 for non-cryptographic checksums only
	h := md5.New()
	_ = h
}

// Test case: #nosec with empty justification (will fail)
func nosecEmptyJustification() {
	// #nosec G401 -- // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec with whitespace-only justification (will fail)
func nosecWhitespaceJustification() {
	// #nosec G401 --   // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore without justification (will fail)
func gosecIgnoreNoJustification() {
	//gosec:ignore G401 // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with justification (OK)
func gosecIgnoreWithJustification() {
	//gosec:ignore G401 -- Using MD5 for non-cryptographic checksums
	h := md5.New()
	_ = h
}

// Test case: gosec:disable without justification (will fail)
func gosecDisableNoJustification() {
	//gosec:disable G101 // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	const apiKey = "placeholder"
	_ = apiKey
}

// Test case: gosec:disable with justification (OK)
func gosecDisableWithJustification() {
	//gosec:disable G101 -- This is a false positive
	const apiKey = "placeholder"
	_ = apiKey
}

// Test case: #nosec without rules but with justification (OK)
func nosecNoRulesWithJustification() {
	// #nosec -- Verified safe after security review
	h := md5.New()
	_ = h
}

// Test case: Multiple rules with justification (OK)
func nosecMultipleRulesWithJustification() {
	// #nosec G201 G202 G203 -- SQL concatenation reviewed and safe
	h := md5.New()
	_ = h
}
