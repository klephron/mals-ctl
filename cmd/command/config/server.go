package config

import (
	"fmt"
	"mals-ctl/cmd/config"
	"mals-ctl/cmd/runtime"
	"mals-ctl/internal/encoding/yaml"

	"github.com/spf13/cobra"
)

func NewServerCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Manage server configurations defined in config",
	}

	cmd.AddCommand(newServerLsCommand(c, io))
	cmd.AddCommand(newServerGetCommand(c, io))
	cmd.AddCommand(newServerAddCommand(c, io))
	cmd.AddCommand(newServerRemoveCommand(c, io))

	return cmd
}

func newServerLsCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := c.Config()
			if err != nil {
				return err
			}
			for _, srv := range cfg.Servers {
				fmt.Fprintf(io.Out, "%v\n", string(srv.Name))
			}
			return nil
		},
	}
}

func newServerGetCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "get [name]",
		Short: "Get server(s) comprehensive info",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := c.Config()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				out, err := yaml.Marshal(cfg.Servers)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			} else {
				var server *config.Server
				name := args[0]

				for _, srv := range cfg.Servers {
					if srv.Name == name {
						server = srv
					}
				}

				if server == nil {
					err = fmt.Errorf("server %q does not exist\n", name)
					return err
				}

				out, err := yaml.Marshal(server)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			}

			return nil
		},
	}
}

func newServerAddCommand(c runtime.Context, _ runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "add <name> <url>",
		Short: "Add server",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			url := args[1]

			server, err := c.Store().AddServer(name, url)
			if err != nil {
				return err
			}
			if server == nil {
				return fmt.Errorf("server %q already exists", name)
			}

			return nil
		},
	}
}

func newServerRemoveCommand(c runtime.Context, _ runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "remove <name>",
		Short: "Remove server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			server, err := c.Store().RemoveServer(args[0])
			if err != nil {
				return err
			}

			if server == nil {
				return fmt.Errorf("server %q does not exist", name)
			}

			return nil
		},
	}
}
