## types.Object (interface)

#### implemented by

```
*Builtin
*Const
*Func
*Label
*Nil
*PkgName
*TypeName
*Var
```

#### Type()

```
func (main.cli_args).evaluate() error
```
=> `func() error`


```
func main.is_flag_passed(flag_name string) bool
```
=> func(flag_name string) bool


The `Type()` function seems to return the `func` keyword, \
all arguments and all return values of a function.

It does NOT return information about the object of a method.
