package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

type cli_args struct {
	file_path    string
	search_flag  bool
	all_flag     bool
	verbose_flag bool
}

func parse_cli_args() cli_args {
	args := cli_args{}

	flag.StringVar(&args.file_path, "f", "", "print relevant targets of the given Makefile")
	flag.StringVar(&args.file_path, "file", "", "print relevant targets of the given Makefile")
	flag.BoolVar(&args.search_flag, "s", false, "search current dir for Makefiles")
	flag.BoolVar(&args.search_flag, "search", false, "search current dir for Makefiles")
	flag.BoolVar(&args.all_flag, "a", false, "print all targets")
	flag.BoolVar(&args.all_flag, "all", false, "print all targets")
	flag.BoolVar(&args.verbose_flag, "v", false, "verbose output")
	flag.BoolVar(&args.verbose_flag, "verbose", false, "verbose output")
	flag.Parse()

	return args
}

func (args cli_args) evaluate() error {
	var err error
	// need to use function to detect if a string flag is set
	file_flag := is_flag_passed("f") || is_flag_passed("file")

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if args.search_flag && file_flag {
		err = errors.New("Flags -s and -f are exclusive.")
		return err
	}

	// either search or file flag must be present to do anything useful
	if !args.search_flag && !file_flag {
		flag.Usage()
	}

	if args.search_flag {
		search_result(args)
	}

	if file_flag {
		err := single_file_result(args)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	return err
}

func is_flag_passed(flag_name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if flag_name == f.Name {
			found = true
		}
	})

	return found
}
