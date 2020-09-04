package missinject_test

import (
	"testing"

	"github.com/theoden9014/miss-inject"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, missinject.Analyzer, "a")
}