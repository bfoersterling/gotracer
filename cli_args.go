package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type cli_args struct {
	go_dir           string
	entrypoint       string
	list_entrypoints bool
	list_uncalled    bool
	version          bool
}

func get_cli_args() cli_args {
	args := cli_args{}

	flag.StringVar(&args.go_dir, "d", ".", "directory with go source files")
	flag.StringVar(&args.go_dir, "dir", ".", "directory with go source files")
	flag.StringVar(&args.entrypoint, "e", "main", "entry point for the call tree")
	flag.StringVar(&args.entrypoint, "entrypoint", "main", "entry point for the call tree")
	flag.BoolVar(&args.list_entrypoints, "l", false, "list all possible entry points")
	flag.BoolVar(&args.list_entrypoints, "list", false, "list all possible entry points")
	flag.BoolVar(&args.list_uncalled, "u", false, "list all uncalled functions")
	flag.BoolVar(&args.list_uncalled, "uncalled", false, "list all uncalled functions")
	flag.BoolVar(&args.version, "V", false, "print version")
	flag.BoolVar(&args.version, "version", false, "print version")
	flag.Parse()

	return args
}

func (args cli_args) evaluate() {
	if args.version {
		fmt.Printf("%s %s, commit: %s, build at: %s.\n", os.Args[0], version, commit, date)
		os.Exit(0)
	}

	if len(os.Args) > 1 && os.Args[1] == "." {
		flag.Usage()
		os.Exit(1)
	}

	afps, fset, err := parse_dir_afps(args.go_dir)

	if err != nil {
		log.Fatal(err)
	}

	if args.list_entrypoints {
		fmt.Printf("%v\n", list_all_entrypoints(fset, afps))
		os.Exit(0)
	}

	if args.list_uncalled {
		fmt.Printf("%v\n", list_uncalled_funcs(fset, afps))
		os.Exit(0)
	}

	verbose_calltree(fset, afps, args.entrypoint)
}
