# chkreplace

<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/> [![CircleCI](https://circleci.com/gh/po3rin/chkreplace.svg?style=shield)](https://circleci.com/gh/po3rin/gotree) [![GolangCI](https://golangci.com/badges/github.com/po3rin/chkreplace.svg)](https://golangci.com)

Go 1.12 released !!!
https://golang.org/doc/go1.12

ReplaceAll Methods is available from Go 1.12. Let's find a method passing negative number to Replace methods

## Instalation

```
$ go get github.com/po3rin/chkreplace/cmd/chkreplace
```

## Usage

as CLI

```bash
chkreplace github.com/po3rin/chkreplace/example
example/replace.go:11:17: detect strings.Replace(...) use negative value -1
example/replace.go:23:15: detect bytes.Replace(...) use negative value -1
```

as Analyzer

```go
func main() {
	multichecker.Main(
		// skeletonで生成したAnalyzerの初期コードの場合はinspect.Analyzerの結果に依存している
		inspect.Analyzer,
		checkreplace.Analyzer,
	)
}
```