package log

import (
	"mals-ctl/cmd/runtime"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, streams runtime.IOStreams) *cobra.Command {
	command := &cobra.Command{
		Use:   "log",
		Short: "Manage logs",
	}

	return command
}
