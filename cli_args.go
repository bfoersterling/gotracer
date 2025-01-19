package main

import (
	"flag"
	"log"
	"os"
)

type cli_args struct {
	go_dir     string
	entrypoint string
}

func get_cli_args() cli_args {
	args := cli_args{}

	flag.StringVar(&args.go_dir, "d", ".", "directory with go source files")
	flag.StringVar(&args.go_dir, "dir", ".", "directory with go source files")
	flag.StringVar(&args.entrypoint, "e", "main", "entry point for the call tree")
	flag.StringVar(&args.entrypoint, "entrypoint", "main", "entry point for the call tree")
	flag.Parse()

	return args
}

func (args cli_args) evaluate() {
	if len(os.Args) > 1 && os.Args[1] == "." {
		flag.Usage()
		os.Exit(1)
	}

	afps, fset, err := parse_dir_afps(args.go_dir)

	if err != nil {
		log.Fatal(err)
	}

	verbose_calltree(fset, afps, args.entrypoint)
}
