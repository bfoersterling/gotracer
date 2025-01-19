These are all the "builtin" functions:\
https://pkg.go.dev/builtin#pkg-functions

The builtin functions are not in the Go source code that is publicly available though.\
The purpose for the `builtin.go` file is for documentation only.\
The builtin functions are deeper in the compiler and you cant seem to get to them.

Since builtin function are unlikely to change due the backward compatibility of Go \
the builtin functions are hardcoded in this tool.

Additional builtin function might appear in future versions of Go that have to be added \
to this source code.
