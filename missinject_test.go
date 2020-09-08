package missinject_test

import (
	"testing"

	"github.com/theoden9014/missinject"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, missinject.Analyzer, "a")
}

func TestValidate(t *testing.T) {
	if err := analysis.Validate([]*analysis.Analyzer{missinject.Analyzer}); err != nil {
		t.Errorf("failed to validate missinject.Analyzer: %v", err)
	}
}
