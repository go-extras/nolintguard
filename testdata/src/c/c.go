package c

import (
	"crypto/md5"
)

// With default config (require-justification=false),
// #nosec without justification should be allowed (no diagnostics)
func nosecWithoutJustification() {
	// #nosec G401
	h := md5.New()
	_ = h
}

// This should always be allowed
func nosecWithJustification() {
	// #nosec G401 -- Using MD5 for checksums only
	h := md5.New()
	_ = h
}

// This should always be allowed
func gosecIgnoreWithJustification() {
	//gosec:ignore G401 -- Safe usage
	h := md5.New()
	_ = h
}
