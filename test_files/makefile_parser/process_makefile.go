package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strings"
)

func get_special_targets() []string {
	return []string{
		".PHONY",
		".SUFFIXES",
		".DEFAULT",
		".PRECIOUS",
		".INTERMEDIATE",
		".NOTINTERMEDIATE",
		".SECONDARY",
		".SECONDEXPANSION",
		".DELETE_ON_ERROR",
		".IGNORE",
		".LOW_RESOLUTION_TIME",
		".SILENT",
		".EXPORT_ALL_VARIABLES",
		".NOTPARALLEL",
		".ONESHELL",
		".POSIX",
	}
}

type makefile struct {
	path    string
	targets []string
}

func has_relevant_target(unfolded_line string, all_flag bool) bool {
	target, _, _ := strings.Cut(unfolded_line, ":")
	variable_regex := regexp.MustCompile(`\$(.*)`)

	if all_flag {
		return true
	}

	// exclude anything that looks like a file
	if strings.Contains(target, "/") || strings.Contains(target, ".") {
		return false
	}

	// exclude pattern rules
	if strings.Contains(target, "%") {
		return false
	}

	for _, special_target := range get_special_targets() {
		trimmed_line := strings.TrimSpace(unfolded_line)
		if strings.HasPrefix(trimmed_line, special_target) {
			return false
		}
	}

	// exclude targets with variables in them
	if variable_regex.MatchString(target) {
		return false
	}

	return true
}

func parse_makefile(args cli_args) makefile {
	mf := makefile{args.file_path, make([]string, 0)}

	mf_content := file_to_string(args.file_path)

	mf.path = args.file_path

	for len(mf_content) > 0 {
		sm := new_statement()
		err := sm.read_statement(&mf_content)

		if err != nil {
			log.Fatal("read_statement() failed")
		}

		sm.remove_comment()

		sm.parse()

		if sm.has_target {
			mf.targets = append(mf.targets, sm.content)
		}
	}

	return mf
}

func single_file_result(args cli_args) error {
	var err error
	if args.file_path == "" {
		err = errors.New("File path may not be an empty string!")
		return err
	}
	mf := parse_makefile(args)

	if args.verbose_flag {
		name_length, _ := fmt.Printf("| Targets found in Makefile \"%s\" |\n", mf.path)

		for i := 0; i < (name_length - 1); i++ {
			fmt.Printf("-")
		}
		fmt.Printf("\n")
	}

	ordered_targets := mf.targets

	slices.Sort(ordered_targets)

	for _, target := range ordered_targets {
		if !args.all_flag {
			target, _, _ = strings.Cut(target, ":")
		}
		if has_relevant_target(target, args.all_flag) {
			fmt.Println(target)
		}
	}

	return err
}
