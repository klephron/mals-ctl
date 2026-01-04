package main

import (
	"mals-ctl/cmd/command"
	"os"
)

func main() {
	cmd := command.NewCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
