#### Entrypoint As CLI Argument

The entrypoint is a string that matches exactly one FuncDecl.

There can be multiple calls to the same FuncDecl.\
So the first call found when traversing the tree \
is being used as the entrypoint.

#### Can The Entrypoint Be a Method?

The entrypoint can be a method, but the receiver type \
has to be specified.

```
gotracer -d test_files/makefile_parser/ -e '(*statement).read_statement'
```

#### Search types.Info Uses Field for Unique String

The `value.String()` part of the map contains this information:
```
key: parse	value:func (*main.cli_args).parse()
key: parse	value:func main.parse() string
```

You can trim the `func ` part \
and the return values (everything after the closing bracket) \
since you cannot redeclare the same function just with different return values.

Then you still have the package name/path.\
You should be able to cut it by using the `value.Pkg()` function.\
This is ok since parsing a dir with mixed packages will throw an error while parsing.

#### all possible entrypoints

How to get all possible entrypoints?

Maybe use the `Defs` field of the `types.Info` struct.

OR

The FuncDecl (`Name` + `Recv` fields) \
=> `Recv` is just a pointer to an `ast.FieldList` and contains a slice of `ast.Fields`.\
Those `ast.Fields` contain `Names` which are `ast.Idents` that can return a `Name` (string).\
So you would need two nested loops to get the receiver?
