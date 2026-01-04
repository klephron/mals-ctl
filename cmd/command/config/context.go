package config

import (
	"fmt"
	"mals-ctl/cmd/runtime"

	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

func NewContextCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "context",
		Short: "Manage context defaults",
	}

	cmd.AddCommand(newContextGetCommand(c, io))
	cmd.AddCommand(newContextSetCommand(c, io))

	return cmd
}

func newContextGetCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get context",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := c.Config()
			if err != nil {
				return err
			}

			if config.Context != nil {
				out, err := yaml.Marshal(config.Context)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			}

			return nil
		},
	}
}

func newContextSetCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set context values",
	}

	cmd.AddCommand(newContextSetServerCommand(c, io))

	return cmd
}

func newContextSetServerCommand(c runtime.Context, _ runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "server <name>",
		Short: "Set context server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Store().SetContextServer(args[0])
		},
	}
}
