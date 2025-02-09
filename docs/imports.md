#### sources

https://go.dev/wiki/SettingGOPATH

#### general problem

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
