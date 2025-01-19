package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func find_makefiles(dir_path string) []string {
	var matched_files []string
	err := filepath.Walk(dir_path, func(walk_path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", walk_path, err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(path.Base(walk_path), "Makefile") {
			matched_files = append(matched_files, walk_path)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir_path, err)
		return nil
	}

	return matched_files
}

func pretty_print_makefiles(files []string, args cli_args) {
	if args.verbose_flag {
		stdout_length, _ := fmt.Println("| The following Makefiles were found in this directory |")
		for i := 0; i < (stdout_length - 1); i++ {
			fmt.Printf("-")
		}
		fmt.Println()
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func search_result(args cli_args) {
	makefiles := find_makefiles(".")
	pretty_print_makefiles(makefiles, args)
	os.Exit(0)
}
