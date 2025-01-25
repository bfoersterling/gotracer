# gotracer

A calltree for local Go projects.\
Does not include library calls.

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
$ ./gotracer -d . -e calltree
calltree
|--get_fcall_from_slice
|--(fcall).get_children
|  `--get_builtin_funcs
|--calltree->calltree (recursive)
`--calltree->calltree (recursive)
```

For methods make sure to quote the name of the function call:
```
./gotracer -e '(cli_args).evaluate' -d test_files/makefile_parser/
```

#### performance

This program uses the type checker that is part of the Go standard library in `go/types`.

The `Check` method from the `go/types` package initially runs for all Go files.\
That is the performance bottleneck.

Writing a custom type checker for structs is currently not an option.

Another option might be to use [pointer analysis](https://en.wikipedia.org/wiki/Pointer_analysis).

## TODO

- detect unused functions
- option to show all (used) functions (all possible entrypoints)
- deal with init() functions \
(have to be artificially included when generating `fcall`s)
