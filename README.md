# gotracer

A calltree for local Go projects.\
Does not support calls to external libraries.

## usage

Show tree of Go function calls in the current dir:
```
$ ./gotracer
main
|--parse_dir_afps
|  `--get_gofiles
`--verbose_calltree
   |--new_func_center
   |  |--get_fds_from_afps
   |  `--get_func_uses
   |     `--get_type_info
   |--(func_center).get_fcalls
   |  |--get_tree_string
   |  `--is_method
   `--calltree
      |--get_fcall_from_slice
      |--(fcall).get_children
      |  `--get_builtin_funcs
      |--calltree->calltree (recursive)
      `--calltree->calltree (recursive)
```

The default entry point is the `main` function.

To specify another entry point.\
Which is the first call that is found when traversing function calls:
```
$ gotracer -d . -e calltree
calltree
|--get_fcall_from_slice
|--(fcall).get_children
|  `--get_builtin_funcs
|--calltree->calltree (recursive)
`--calltree->calltree (recursive)
```

For methods make sure to quote the name of the function call:
```
gotracer -e '(cli_args).evaluate' -d test_files/makefile_parser/
```

To list all possible entryponts:
```
gotracer -l
```

#### implementation

This program exclusively uses `Go`'s standard library.\
There are no external dependencies.

#### performance

`gotracer` uses the type checker that is part of the Go standard library in `go/types`.

The `Check` method from the `go/types` package initially runs for all Go files.\
That is the performance bottleneck.

Writing a custom type checker for is currently not an option.

Another option would be to use [pointer analysis](https://en.wikipedia.org/wiki/Pointer_analysis).

## TODO

- option to list uncalled functions
- option to list unreachable functions (needs correct entrypoint and will not work with external libs)
- maybe include external calls
