package listener

import (
	"context"
	"fmt"
	"mals-ctl/cmd/runtime"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listener",
		Short: "Manage listeners",
	}

	cmd.AddCommand(newLsCommand(c, io))

	return cmd
}

func newLsCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List listeners",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}
			res, err := api.ListenerGetAllWithResponse(context.TODO())
			if err != nil {
				return err
			}
			listeners := *res.JSON200
			for _, listener := range listeners {
				fmt.Fprintf(io.Out, "%v\n", listener.Name)
			}
			return nil
		},
	}
}
