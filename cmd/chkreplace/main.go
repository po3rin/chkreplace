package main

import (
	checkreplace "github.com/po3rin/chkreplace"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func main() {
	multichecker.Main(
		// skeletonで生成したAnalyzerの初期コードの場合はinspect.Analyzerの結果に依存している
		inspect.Analyzer,
		checkreplace.Analyzer,
	)
}
