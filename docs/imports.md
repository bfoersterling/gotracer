When you install this program as a self contained binary \
where shall it search for the imported libraries?

For example the `fmt.Printf` function could bin in \
`/usr/lib/go/src/fmt/print.go`.

`Go` seems to search in whatever is in `go env GOROOT`:
```
$ go env GOROOT
/usr/lib/go
```

#### TODO: difference between GOPATH and GOROOT
