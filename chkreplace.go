package chkreplace

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer - chkreplace analyzer checks Replace methods use negative number as arg.
var Analyzer = &analysis.Analyzer{
	Name: "checkreplace",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var name = "Replace"

// Doc - this analyzer description.
const Doc = "checkreplace validates Replace methods to use negative number"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if call, ok := n.(*ast.CallExpr); ok {
			var id *ast.SelectorExpr
			var idx *ast.Ident

			switch fun := call.Fun.(type) {
			case *ast.SelectorExpr:
				id = fun
			default:
				return
			}

			switch x := id.X.(type) {
			case *ast.Ident:
				idx = x
			default:
				return
			}

			if id != nil && idx != nil && !pass.TypesInfo.Types[id].IsType() {
				if id.Sel.Name == "Replace" && (idx.Name == "strings" || idx.Name == "bytes") {
					var expr *ast.UnaryExpr
					if len(call.Args) != 4 {
						return
					}
					switch arg := call.Args[3].(type) {
					case *ast.UnaryExpr:
						expr = arg
					default:
						return
					}

					switch a := expr.X.(type) {
					case *ast.BasicLit:
						if a.Kind == token.INT && expr.Op == token.SUB {
							pass.Reportf(call.Lparen, "detect %s.%s(...) use negative value -%v", idx.Name, id.Sel.Name, a.Value)
						}
					}

				}
			}
		}
	})

	return nil, nil
}
