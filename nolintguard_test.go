package nolintguard_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/go-extras/nolintguard"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	t.Run("default configuration", func(t *testing.T) {
		// Test with default flags (require-justification=false)
		analyzer := nolintguard.NewAnalyzer()
		analysistest.Run(t, testdata, analyzer, "a")
	})

	t.Run("with justification required", func(t *testing.T) {
		// Create fresh analyzer and enable justification requirement
		analyzer := nolintguard.NewAnalyzer()
		err := analyzer.Flags.Set("require-justification", "true")
		if err != nil {
			t.Fatal(err)
		}
		// Test gosec justification
		analysistest.Run(t, testdata, analyzer, "b")
		// Test revive justification
		analysistest.Run(t, testdata, analyzer, "d")
	})

	t.Run("justification not required", func(t *testing.T) {
		// Test that #nosec and //revive: without justification are allowed when flag is false
		analyzer := nolintguard.NewAnalyzer()
		// require-justification defaults to false
		analysistest.Run(t, testdata, analyzer, "c")
	})

	t.Run("with forbidden linters", func(t *testing.T) {
		// Test custom forbidden linters
		analyzer := nolintguard.NewAnalyzer()
		err := analyzer.Flags.Set("forbidden-linters", "staticcheck,unused")
		if err != nil {
			t.Fatal(err)
		}
		analysistest.Run(t, testdata, analyzer, "e")
	})

	t.Run("edge cases", func(t *testing.T) {
		// Test edge cases: duplicates, trailing commas, whitespace, etc.
		analyzer := nolintguard.NewAnalyzer()
		analysistest.Run(t, testdata, analyzer, "f")
	})

	t.Run("revive edge cases with justification required", func(t *testing.T) {
		// Test revive directive variations with justification requirement
		analyzer := nolintguard.NewAnalyzer()
		err := analyzer.Flags.Set("require-justification", "true")
		if err != nil {
			t.Fatal(err)
		}
		analysistest.Run(t, testdata, analyzer, "g")
	})

	t.Run("combined flags", func(t *testing.T) {
		// Test combination of require-justification and forbidden-linters
		analyzer := nolintguard.NewAnalyzer()
		err := analyzer.Flags.Set("require-justification", "true")
		if err != nil {
			t.Fatal(err)
		}
		err = analyzer.Flags.Set("forbidden-linters", "staticcheck,unused")
		if err != nil {
			t.Fatal(err)
		}
		analysistest.Run(t, testdata, analyzer, "h")
	})

	t.Run("gosec and nosec edge cases with justification required", func(t *testing.T) {
		// Test gosec and #nosec directive edge cases
		analyzer := nolintguard.NewAnalyzer()
		err := analyzer.Flags.Set("require-justification", "true")
		if err != nil {
			t.Fatal(err)
		}
		analysistest.Run(t, testdata, analyzer, "i")
	})
}
