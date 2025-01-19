package main

import (
	"log"
)

func main() {
	program_flags := parse_cli_args()

	err := program_flags.evaluate()

	if err != nil {
		log.Fatal(err)
	}
}
