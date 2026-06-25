package logcapitalization_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/sila-chain/Sila-Consensus-Core/v7/build/bazel"
	"github.com/sila-chain/Sila-Consensus-Core/v7/tools/analyzers/logcapitalization"
)

func init() {
	if bazel.BuiltWithBazel() {
		bazel.SetGoEnv()
	}
}

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, logcapitalization.Analyzer, "a")
}
