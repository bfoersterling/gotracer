#### sources

https://go.dev/wiki/SettingGOPATH

#### general problem

Binaries built with:
```
go build -trimpath
```
will not work.
```
2025/02/09 18:19:11 new_func_center() failed with err:
./ast_conversions.go:4:2: could not import errors (can't find import: "errors": cannot find package "errors" in any of:
	($GOROOT not set)
	/home/user/go/src/errors (from $GOPATH))
```


When you install this program as a self contained binary \
where shall it search for the imported libraries?

For example the `fmt.Printf` function could bin in \
`/usr/lib/go/src/fmt/print.go`.

`Go` seems to search standard library packages in whatever is in `go env GOROOT`:
```
$ go env GOROOT
/usr/lib/go
```

#### GOROOT

From "The Go Programming Language" (Alan A. A. Donovan, Brian W. Kernighan):
> GOROOT ... specifies the root directory of the Go distribution, which provides \
all the packages of the **standard library**.

> ... for example, the source files of the fmt package reside in the \
$GOROOT/src/fmt directory.

> Users never need to set GOROOT since, by default, the go tool will use \
the location where it was installed.

#### GOPATH

From "The Go Programming Language" (Alan A. A. Donovan, Brian W. Kernighan):
> GOPATH has three subdirectories.\
The `src` subdirectory holds source code.\
Each package resides in a directory whose name relative to $GOPATH/src is \
the package's import path, such as "gopl.io/ch1/helloworld".

> The `bin` subdirectory holds executable programs...

https://go.dev/wiki/SettingGOPATH :
> The GOPATH environment variable specifies the location of your workspace. If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows.

https://go.dev/wiki/SettingGOPATH#unix-systems :
> Note that GOPATH must not be the same path as your Go installation.

Setting the GOPATH:
```
go env -w GOPATH=$HOME/go
```

#### go env

From "The Go Programming Language" (Alan A. A. Donovan, Brian W. Kernighan):
> The `go env` command prints the effective values of the environment variables \
relevant to the toolchain, including the default values for the missing ones.

#### assumptions about the user

The user of `gotracer` will most likely be a Go developer and thus have \
a Go installation on their system.

Then it is possible to use `GOROOT` and `GOPATH` on the system.\
Otherwise you would need to download the entire standard library and \
external packages because `go/parser` can only parse local files and dirs, \
but not directly from a URL.

#### possible solutions

- set `GOROOT` and `GOPATH` environment variables on runtime to whatever \
`go env GOROOT` and `go env GOPATH` output
=> if Go is not installed then this operation should result in an error
