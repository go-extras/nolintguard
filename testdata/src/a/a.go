package a

import (
	"crypto/md5"
	"errors"
)

// Test case: forbidden gosec directive
func useGosec() {
	//nolint:gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: forbidden revive directive
func useRevive() error {
	//nolint:revive // want "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	return errors.New("test")
}

// Test case: multiple linters including forbidden ones
func useMultipleLinters() {
	//nolint:gosec,errcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: revive with other linters
func useReviveWithOthers() error {
	//nolint:revive,errcheck // want "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	return errors.New("test")
}

// Test case: both gosec and revive
func useBothForbidden() error {
	//nolint:gosec,revive // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead" "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	h := md5.New()
	_ = h
	return errors.New("test")
}

// Test case: allowed - #nosec is OK
func useNosec() {
	// #nosec G401 -- Using MD5 for non-cryptographic checksums only
	h := md5.New()
	_ = h
}

// Test case: allowed - #nosec with multiple rules
func useNosecMultiple() {
	// #nosec G201 G202 G203 -- Reviewed and safe
	h := md5.New()
	_ = h
}

// Test case: allowed - gosec:ignore directive
func useGosecIgnore() {
	//gosec:ignore G401 -- Using MD5 for non-cryptographic checksums only
	h := md5.New()
	_ = h
}

// Test case: allowed - gosec:disable directive
func useGosecDisable() {
	//gosec:disable G401 -- This is a false positive
	h := md5.New()
	_ = h
}

// Test case: allowed - native revive directive is OK
func useNativeRevive() error {
	//revive:disable-next-line
	return errors.New("test")
}

// Test case: allowed - other nolint directives are OK by default
func useOtherLinters() error {
	//nolint:errcheck
	return errors.New("test")
}

// Test case: plain nolint without arguments (allowed by default)
func usePlainNolint() error {
	//nolint
	return errors.New("test")
}

// Test case: gosec in the middle of list
func useGosecMiddle() {
	//nolint:errcheck,gosec,staticcheck // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: whitespace variations
func useWhitespace() {
	//nolint: gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: block comment
func useBlockComment() {
	/* nolint:gosec */ // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: space after // (both formats should be detected)
func useSpaceAfterSlash() {
	// nolint:gosec // want "nolintguard: //nolint:gosec is forbidden; use #nosec or //gosec:ignore instead"
	h := md5.New()
	_ = h
}

// Test case: space after // for revive
func useSpaceAfterSlashRevive() error {
	// nolint:revive // want "nolintguard: //nolint:revive is forbidden; use native revive directives instead"
	return errors.New("test")
}

// Test case: just mentioning nolint in text is OK
// This is a comment about nolint directives in general.
func mentionNolint() {
	// We should use nolint carefully
}
