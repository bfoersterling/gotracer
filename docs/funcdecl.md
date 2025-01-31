## ast.FuncDecl

#### Connection to types.Info Map Keys

#### position matching (get from Uses to FuncDecl) (TLDR)

`my_funcdecl.Name.Pos()` matches the `value.Pos()` of the `Uses` types map.

This allows you to get directly from **Uses to FuncDecl**.

#### .Pos()

```
|
V
func main() {
...
```

#### .End()

The `End()` position is at the character after the closing squirly bracket.

```
func main() {
    ...
}
 ^
 |
```

#### .Name.Pos()


```
     |
     V
func main() {
...
```

#### .Name.End()

`.Name.End()` points to the opening parenthesis:
```
         |
         V
func main() {
...
```

#### .Type.End()

```
           |
           V
func main() {
...
```
