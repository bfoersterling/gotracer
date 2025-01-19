## TLDR

The `Lparen` field from ast.CallExpr should match the \
`key.End()` of the info.Uses map.

## ast.CallExpr

#### Pos()

```
|
V
parse()
```

```
    |
    V
    cli_args.parse()
```

#### End()

`End()` points to the token after the closing bracket.\
In these cases it were the new line characters.

```
       |
       V
parse()
```

```
                    |
                    V
    cli_args.parse()
```

#### Lparen

```
     |
     V
parse()
```

```
                  |
                  V
    cli_args.parse()
```

#### Rparen

```
      |
      V
parse()
```

```
                   |
                   V
    cli_args.parse()
```


## ast.CallExpr.Fun

#### Fun.Pos()
```
|
V
parse()
```

```
    |
    V
    cli_args.parse()
```

#### Fun.End()

Fun.End() points to the opening bracket of the function call.
```
     |
     V
parse()
```

```
                  |
                  V
    cli_args.parse()
```
