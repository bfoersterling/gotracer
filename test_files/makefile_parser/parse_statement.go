package main

import (
	"errors"
	"regexp"
	"strings"
)

// linked list or slice of structs?
// the slice of struct function needs to receive and return a struct instance
// the linked list should only need a pointer to the current statement
// but only if I have a previous pointer as well in the list
// so I probably need a doubly linked list

// what information do I need from the previous statement?
// - special variables like .RECIPEPREFIX
// - was the previous statement a "target line"
// - was the previous statement a recipe line?

type statement struct {
	content       string
	recipeprefix  rune
	has_target    bool
	is_recipe     bool
	is_assignment bool
}

// default constructor
func new_statement() statement {
	return statement{
		"",
		'\t',
		false,
		false,
		false,
	}
}

// this function should only read the first logical line
// lexical analysis or interpretation will be done by other functions
// read from a file directly or put file in a string and read from string?
// you dont want to open and close the file for every statement
func (sm *statement) read_statement(remaining_content *string) error {
	temp_sm := ""
	escape_next := false
	var err error
	cursor := 0

	if len(*remaining_content) <= 0 {
		err = errors.New("There is no content left to parse.")
		return err
	}

	for _, char := range *remaining_content {
		cursor++

		if char == '\n' && !escape_next {
			break
		}

		if char == '\n' && escape_next {
			// remove backslash before newline
			temp_sm = temp_sm[:len(temp_sm)-1]
			continue
		}

		if char == '\\' && !escape_next {
			escape_next = true
		}

		if char != '\\' {
			escape_next = false
		}

		temp_sm += string(char)
	}

	sm.content = temp_sm

	*remaining_content = (*remaining_content)[cursor:]

	return err
}

func (sm *statement) remove_comment() {
	temp_buf := ""
	escape_next := false

	for _, char := range sm.content {
		if char == '\\' && !escape_next {
			escape_next = true
		}
		if char == '#' && !escape_next {
			break
		}
		if char != '\\' {
			escape_next = false
		}
		temp_buf += string(char)
	}
	sm.content = temp_buf
}

func (sm *statement) parse() {
	if strings.HasPrefix(sm.content, string(sm.recipeprefix)) {
		sm.is_recipe = true
		return
	}

	simple_assignment := regexp.MustCompile(`:=`)

	if simple_assignment.MatchString(sm.content) {
		sm.is_assignment = true
		return
	}

	trimmed_sm := strings.TrimSpace(sm.content)
	prefix := ""
	is_escaped := false
	inside_quotes := false

	for _, char := range trimmed_sm {
		switch char {
		case '\\':
			if !is_escaped {
				is_escaped = true
			}
			continue
		case '"', '\'':
			if !is_escaped {
				inside_quotes = !inside_quotes
			}
		case '=':
			if !is_escaped && !inside_quotes {
				sm.is_assignment = true
				if strings.HasPrefix(trimmed_sm, ".RECIPEPREFIX") {
					_, prefix, _ = strings.Cut(trimmed_sm, "=")
					prefix = strings.TrimSpace(prefix)
					sm.recipeprefix = rune(prefix[0])
				}
				return
			}
		case ':':
			if !is_escaped && !inside_quotes {
				sm.has_target = true
				return
			}
		}
		is_escaped = false
	}
}
