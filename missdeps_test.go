package missdeps_test

import (
	"testing"

	"github.com/theoden9014/missdeps"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, missdeps.Analyzer, "a")
}