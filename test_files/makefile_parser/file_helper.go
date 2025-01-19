package main

import (
	"bufio"
	"log"
	"os"
)

func file_to_string(file_path string) string {
	bytes, err := os.ReadFile(file_path)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func file_to_string_slice(file_path string) []string {
	file, err := os.Open(file_path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
