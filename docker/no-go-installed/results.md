As expected `gotracer` will not work if `go` is not installed:
```
/workspace # gotracer 
panic: conf.Check failed: ./ast_conversions.go:4:2: could not import errors (can't find import: "errors": cannot find package "errors" in any of:
		/usr/lib/go/src/errors (from $GOROOT)
		/root/go/src/errors (from $GOPATH))

goroutine 1 [running]:
main.cli_args.evaluate({{0x689099, 0x1}, {0x6892ec, 0x4}, 0x0, 0x0, 0x0})
	gotracer/cli_args.go:64 +0x44a
main.main()
	gotracer/main.go:12 +0x7e
```

Maybe moving to googles `x/tools` libraries will solve this problem.
