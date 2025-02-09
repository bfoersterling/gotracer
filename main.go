package main

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	program_flags := get_cli_args()

	program_flags.evaluate()
}
