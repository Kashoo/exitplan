package main

import (
	"os"

	cmd2 "github.com/Kashoo/exitplan/exitplan-test/cmd"
)

func main() {
	if err := cmd2.Execute(); err != nil {
		os.Exit(1)
	}
}
