When going from a function call to a function declaration \
you first need to get the type of the object that calls the method.

You cannot get the type of the caller from the `ast` package.\
So the `go/types` package is being used.

To match the ast.CallExpr to the corresponding ast.FuncDecl you need \
the Uses field of the types.Info object.\
You get there by matching certain token.Pos values which is described in other \
documents in this dir.
