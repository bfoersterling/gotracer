## types.Info.Uses Map

#### General Map Overview

Definition:
```
Uses map[*ast.Ident]Object
```
(`Object` is `types.Object`)

Put simply, the `Uses` map has the function call as a key (first part) and \
the function definition in the value (second part) of the map.
```
Key: cli_args (cli_args.go:11:13)	Value: type main.cli_args struct{verbose bool} (cli_args.go:3:6)
Key: cli_args (main.go:4:14)		Value: type main.cli_args struct{verbose bool} (cli_args.go:3:6)
Key: true (main.go:5:3)		        Value: const true untyped bool (-)
Key: cli_args (main.go:8:2)	    	Value: var cli_args main.cli_args (main.go:4:2)
Key: parse (main.go:8:11)		    Value: func (*main.cli_args).parse() (cli_args.go:11:22)
Key: parse (main.go:10:2)		    Value: func main.parse() string (cli_args.go:7:6)
Key: bool (cli_args.go:4:10)		Value: type bool (-)
Key: string	(cli_args.go:7:14)		Value: type string (-)
```

#### main function

Since the `main` function is never explicitly called in a Go program \
it is not part of the `Uses` map.

#### key.Pos()

The `Pos()` value of the map key `*ast.Ident`:

```
          |
          V
	flags.parse()

    |
    V
	parse()
```

#### value.Pos()

The `Pos()` value of the map value `Object` goes to the \
function declaration:
```
                     |
                     V
func (args *cli_args)parse() {
```

So it might not be necessary to retrieve information from the `Defs` map.

#### key.End()

The `End()` value of the key:
```
                  |
                  V
    cli_args.parse()
```

#### value.End()

Lead to panic due to nil pointer.
