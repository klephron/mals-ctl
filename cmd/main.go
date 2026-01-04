package main

import (
	"fmt"
	"mals-ctl/cmd/command"
	"os"
)

func main() {
	cmd := command.NewCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
