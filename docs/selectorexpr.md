```
$ go doc ast.SelectorExpr
package ast // import "go/ast"

type SelectorExpr struct {
	X   Expr   // expression
	Sel *Ident // field selector
}
```
