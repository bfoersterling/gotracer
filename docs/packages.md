From "The Go Programming Language" (Alan A. A. Donovan, Brian W. Kernighan) (p293 10.7.3 "Building Packages")
> Since each directory contains one package, each executable program, or command \
in Unix terminology, requires its own directory.

The parse_dir function will throw an error when go files from different packages are in the same dir.

If you take this yaml https://github.com/go-yaml/yaml lib for example.\
All files are package "yaml", but the test files are of package "yaml_test".
