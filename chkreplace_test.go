package chkreplace_test

import (
	"testing"

	"github.com/po3rin/chkreplace"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, chkreplace.Analyzer, "a")
}
