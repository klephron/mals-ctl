package config

import (
	"mals-ctl/cmd/runtime"

	"github.com/spf13/cobra"
)

func NewServerCommand(c runtime.Context, streams runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Manage server configurations defined in config",
	}

	return cmd
}
