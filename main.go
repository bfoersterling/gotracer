package main

func main() {
	program_flags := get_cli_args()

	program_flags.evaluate()
}
