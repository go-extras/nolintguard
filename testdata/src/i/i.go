package i

// Test gosec and #nosec directive edge cases with require-justification enabled

import (
	"crypto/md5"
)

// Test case: #nosec with multiple rules and justification
func nosecMultipleRules() {
	// #nosec G201 G202 G203 -- SQL queries are parameterized
	h := md5.New()
	_ = h
}

// Test case: #nosec with no rules but with justification
func nosecNoRules() {
	// #nosec -- Reviewed by security team
	h := md5.New()
	_ = h
}

// Test case: #nosec with no rules and no justification
func nosecNoRulesNoJustification() {
	// #nosec // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification containing special characters
func nosecSpecialChars() {
	// #nosec G401 -- Safe: using MD5 for ETag generation (RFC 7232 §2.3)
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification containing Unicode
func nosecUnicode() {
	// #nosec G401 -- 安全：仅用于校验和
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification containing emoji
func nosecEmoji() {
	// #nosec G401 -- ✅ Approved by security team
	h := md5.New()
	_ = h
}

// Test case: #nosec with very long justification
func nosecLongJustification() {
	// #nosec G401 -- This is a very long justification that explains in great detail why we are using MD5 here despite it being cryptographically weak because we are only using it for non-security-critical checksums and ETags
	h := md5.New()
	_ = h
}

// Test case: #nosec with multiple dashes in justification
func nosecMultipleDashes() {
	// #nosec G401 -- Using MD5 -- not for crypto -- only for checksums
	h := md5.New()
	_ = h
}

// Test case: #nosec with empty justification (just --)
func nosecEmptyJustification() {
	// #nosec G401 -- // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec with whitespace-only justification
func nosecWhitespaceJustification() {
	// #nosec G401 --    // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with multiple rules
func gosecIgnoreMultipleRules() {
	//gosec:ignore G201 G202 -- SQL is safe
	h := md5.New()
	_ = h
}

// Test case: gosec:disable with justification
func gosecDisableJustified() {
	//gosec:disable G401 -- Legacy code
	h := md5.New()
	_ = h
}

// Test case: gosec:disable without justification
func gosecDisableNoJustification() {
	//gosec:disable G401 // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with empty justification
func gosecIgnoreEmptyJustification() {
	//gosec:ignore G401 -- // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore with whitespace justification
func gosecIgnoreWhitespaceJustification() {
	//gosec:ignore G401 --   // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec directive with extra spaces
func gosecExtraSpaces() {
	//gosec:ignore    G401    --    Justified
	h := md5.New()
	_ = h
}

// Test case: gosec directive with tabs
func gosecTabs() {
	//gosec:ignore	G401	--	Justified
	h := md5.New()
	_ = h
}

// Test case: #nosec in block comment
func nosecBlockComment() {
	/* #nosec G401 */ // want "nolintguard: #nosec directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: #nosec in block comment with justification
func nosecBlockCommentJustified() {
	/* #nosec G401 -- Safe usage */
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore in block comment
func gosecIgnoreBlockComment() {
	/* gosec:ignore G401 */ // want "nolintguard: //gosec: directive must include justification \\(-- reason\\)"
	h := md5.New()
	_ = h
}

// Test case: gosec:ignore in block comment with justification
func gosecIgnoreBlockCommentJustified() {
	/* gosec:ignore G401 -- Safe usage */
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification containing URL
func nosecWithURL() {
	// #nosec G401 -- See https://example.com/security-review for details
	h := md5.New()
	_ = h
}

// Test case: #nosec with justification containing code
func nosecWithCode() {
	// #nosec G401 -- Using `md5.New()` for checksums only
	h := md5.New()
	_ = h
}

