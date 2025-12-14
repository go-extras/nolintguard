package g

// Test revive directive edge cases with require-justification enabled

// Test case: revive:disable-line without justification
func reviveDisableLine() {
	//revive:disable-line // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:disable-line with justification
func reviveDisableLineJustified() {
	//revive:disable-line Temporary workaround
	x := 1
	_ = x
}

// Test case: revive:disable-next-line without justification
func reviveDisableNextLine() {
	//revive:disable-next-line // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:disable-next-line with justification
func reviveDisableNextLineJustified() {
	//revive:disable-next-line Will be fixed in next release
	x := 1
	_ = x
}

// Test case: revive:enable without justification
func reviveEnableNoJustification() {
	//revive:enable // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:enable with justification
func reviveEnableJustified() {
	//revive:enable Issues have been resolved
	x := 1
	_ = x
}

// Test case: revive:enable-line without justification
func reviveEnableLine() {
	//revive:enable-line // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:enable-line with justification
func reviveEnableLineJustified() {
	//revive:enable-line Re-enabling after fix
	x := 1
	_ = x
}

// Test case: revive:enable-next-line without justification
func reviveEnableNextLine() {
	//revive:enable-next-line // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:enable-next-line with justification
func reviveEnableNextLineJustified() {
	//revive:enable-next-line Safe to re-enable
	x := 1
	_ = x
}

// Test case: revive:disable with rule and no justification
func reviveDisableRuleNoJustification() {
	//revive:disable:var-naming // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:disable with rule and justification
func reviveDisableRuleJustified() {
	//revive:disable:var-naming Legacy code compatibility
	x := 1
	_ = x
}

// Test case: revive:enable with rule and no justification
func reviveEnableRuleNoJustification() {
	//revive:enable:var-naming // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive:enable with rule and justification
func reviveEnableRuleJustified() {
	//revive:enable:var-naming Fixed naming issues
	x := 1
	_ = x
}

// Test case: revive directive with extra spaces
func reviveExtraSpaces() {
	//revive:disable    This has extra spaces
	x := 1
	_ = x
}

// Test case: revive directive with tabs
func reviveTabs() {
	//revive:disable	This has tabs
	x := 1
	_ = x
}

// Test case: revive directive with only whitespace after (no justification)
func reviveOnlyWhitespace() {
	//revive:disable    // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive directive with rule and only whitespace (no justification)
func reviveRuleOnlyWhitespace() {
	//revive:disable:exported    // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: revive with Unicode in justification
func reviveUnicode() {
	//revive:disable 临时解决方案 (temporary solution)
	x := 1
	_ = x
}

// Test case: revive with emoji in justification
func reviveEmoji() {
	//revive:disable TODO: fix this 🚧
	x := 1
	_ = x
}

// Test case: revive with very long justification
func reviveLongJustification() {
	//revive:disable This is a very long justification that explains in great detail why we need to disable this particular revive rule for this specific line of code because of legacy compatibility requirements
	x := 1
	_ = x
}

// Test case: Block comment with revive directive
func reviveBlockComment() {
	/* revive:disable */ // want "nolintguard: //revive: directive must include justification \\(reason\\)"
	x := 1
	_ = x
}

// Test case: Block comment with revive directive and justification
func reviveBlockCommentJustified() {
	/* revive:disable Justified in block comment */
	x := 1
	_ = x
}

