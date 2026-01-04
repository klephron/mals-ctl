package config

import (
	"mals-ctl/cmd/runtime"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, streams runtime.IOStreams) *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "Manage configuration",
	}

	return command
}
