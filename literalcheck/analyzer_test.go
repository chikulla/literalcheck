package literalcheck_test

import (
	"testing"

	"github.com/chikulla/literalcheck/literalcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, literalcheck.Analyzer, "main")
}
