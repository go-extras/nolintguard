package d

// Test revive justification requirement

// Without justification (will fail when require-revive-justification is enabled)
func reviveNoJustification() {
	//revive:disable // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// With justification (OK)
func reviveWithJustification() {
	//revive:disable Until the code is stable
	x := 1
	_ = x
}

// With rule-specific disable and justification (OK)
func reviveRuleWithJustification() {
	//revive:disable:exported Code is internal only
	x := 1
	_ = x
}

// With rule-specific disable but no justification (fail)
func reviveRuleNoJustification() {
	//revive:disable:exported // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// enable directive with justification (OK)
func reviveEnableWithJustification() {
	//revive:enable Re-enabling after fixing issues
	x := 1
	_ = x
}
