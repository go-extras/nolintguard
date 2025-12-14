package e

// Test forbidden linters feature

// staticcheck should be forbidden
func useStaticcheck() {
	//nolint:staticcheck // want "nolintguard: //nolint:staticcheck is forbidden"
	x := 1
	_ = x
}

// unused should be forbidden
func useUnused() {
	//nolint:unused // want "nolintguard: //nolint:unused is forbidden"
	x := 1
	_ = x
}

// Multiple linters including forbidden ones
func useMultiple() {
	//nolint:staticcheck,errcheck // want "nolintguard: //nolint:staticcheck is forbidden"
	x := 1
	_ = x
}

// errcheck is NOT forbidden (no diagnostic expected)
func useErrcheck() {
	//nolint:errcheck
	x := 1
	_ = x
}
