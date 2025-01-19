package main

import (
	"log"
	"os"
	"strings"
)

func assert_dir(dir string) {
	file_info, err := os.Stat(dir)

	if err != nil {
		log.Fatal(err)
	}

	if !file_info.IsDir() {
		log.Fatal(dir, " is not a directory.")
	}
}

// gets go files in a dir (nonrecursive)
func get_gofiles(dir string) []string {
	files := make([]string, 0)

	dir_entries, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range dir_entries {
		if !entry.Type().IsRegular() {
			continue
		}
		if strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".go") {
			files = append(files, dir+"/"+entry.Name())
		}
	}

	return files
}
