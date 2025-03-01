package main

import (
	"sort"
	"strings"
)

func list_uncalled_funcs(fc func_center) string {
	uncalled_funcs := make([]string, 0)

	fcalls, err := fc.get_fcalls()

	if err != nil {
		panic(err)
	}

	for _, fcall := range fcalls {
		if fcall.call_name == "main" {
			continue
		}
		if fcall.uses_key == nil {
			uncalled_funcs = append(uncalled_funcs, fcall.call_name)
		}
	}

	sort.Strings(uncalled_funcs)

	return strings.Join(uncalled_funcs, "\n")
}
