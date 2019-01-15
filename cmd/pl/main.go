package main

import (
	"os"

	"github.com/yoheimuta/protolint/internal/cmd"
)

// DEPRECATED: Use cmd/protolint. See https://github.com/yoheimuta/protolint/issues/20.
func main() {
	os.Exit(int(
		cmd.Do(
			os.Args[1:],
			os.Stdout,
			os.Stderr,
		),
	))
}
