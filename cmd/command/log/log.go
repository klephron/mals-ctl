package log

import (
	"context"
	"fmt"
	"mals-ctl/cmd/runtime"
	"mals-ctl/internal/encoding/yaml"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Manage logs",
	}

	cmd.AddCommand(newLsCommand(c, io))
	cmd.AddCommand(newGetCommand(c, io))

	return cmd
}

func newLsCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}
			res, err := api.LogGetAllWithResponse(context.TODO())
			if err != nil {
				return err
			}
			if res.ApplicationproblemJSONDefault != nil {
				return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
			}
			logs := *res.JSON200
			for _, log := range logs {
				fmt.Fprintf(io.Out, "%v\n", log.Name)
			}
			return nil
		},
	}
}

func newGetCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get log(s) comprehensive info",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				res, err := api.LogGetAllWithResponse(context.TODO())
				if err != nil {
					return err
				}
				if res.ApplicationproblemJSONDefault != nil {
					return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
				}
				logs := *res.JSON200
				out, err := yaml.Marshal(logs)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			} else {
				name := args[0]
				res, err := api.LogGetWithResponse(context.TODO(), name)
				if err != nil {
					return err
				}
				if res.ApplicationproblemJSONDefault != nil {
					return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
				}
				log := *res.JSON200
				out, err := yaml.Marshal(log)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			}

			return nil
		},
	}
}
