package main

import (
	"sort"
	"strings"
)

func list_all_entrypoints(fc func_center) string {
	entrypoints := make([]string, 0)

	for _, value := range fc.func_defs {
		entrypoints = append(entrypoints, get_tree_string(value))
	}

	sort.Strings(entrypoints)

	return strings.Join(entrypoints, "\n")
}
