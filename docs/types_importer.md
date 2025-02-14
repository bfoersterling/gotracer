## types.Importer (interface)

```
$ go doc types.Importer
package types // import "go/types"

type Importer interface {
	// Import returns the imported package for the given import path.
	// The semantics is like for ImporterFrom.ImportFrom except that
	// dir and mode are ignored (since they are not present).
	Import(path string) (*Package, error)
}
```

#### srcimporter.New()

This function provides the ability to pass a `build.Context` which contains \
`GOROOT` and `GOPATH` fields.

```
$ go doc srcimporter.New
package srcimporter // import "go/internal/srcimporter"

func New(ctxt *build.Context, fset *token.FileSet, packages map[string]*types.Package) *Importer
    New returns a new Importer for the given context, file set, and map of
    packages. The context is used to resolve import paths to package paths,
    and identifying the files belonging to the package. If the context provides
    non-nil file system functions, they are used instead of the regular package
    os functions. The file set is used to track position information of package
    files; and imported packages are added to the packages map.
```
