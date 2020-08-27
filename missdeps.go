package missdeps

import (
	"go/ast"
	"go/types"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "missdeps",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "missdeps is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CompositeLit:
			t, ok := pass.TypesInfo.TypeOf(n).Underlying().(*types.Struct)
			if !ok {
				return
			}
			if len(n.Elts) == 0 {
				return
			}
			if _, ok := n.Elts[0].(*ast.KeyValueExpr); !ok {
				return
			}

			used := map[string]bool{}
			for _, elt := range n.Elts {
				key := elt.(*ast.KeyValueExpr).Key
				used[key.(*ast.Ident).Name] = true
			}

			for i := 0; i < t.NumFields(); i++ {
				f := t.Field(i)
				fname := f.Name()
				if len(fname) != 0 {
					// skip public fields
					if unicode.IsUpper(rune(fname[0])) {
						continue
					}
				}

				// skip not type interface
				if !types.IsInterface(f.Type()) {
					continue
				}


				if _, ok := used[fname]; !ok {
					pass.Reportf(n.Pos(), "find missing dependency: %v", fname)
				}
			}
		}
	})

	return nil, nil
}
