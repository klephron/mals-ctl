package lsp

import (
	"context"
	"fmt"
	"mals-ctl/cmd/runtime"
	"mals-ctl/internal/encoding/yaml"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lsp",
		Short: "Manage lsps",
	}

	cmd.AddCommand(newLsCommand(c, io))
	cmd.AddCommand(newGetCommand(c, io))

	return cmd
}

func newLsCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List lsps",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}
			res, err := api.LspGetAllWithResponse(context.TODO())
			if err != nil {
				return err
			}
			if res.ApplicationproblemJSONDefault != nil {
				return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
			}
			lsps := *res.JSON200
			for _, lsp := range lsps {
				fmt.Fprintf(io.Out, "%v\n", lsp.Name)
			}
			return nil
		},
	}
}

func newGetCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get lsp(s) comprehensive info",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}

			if len(args) == 0 {
				res, err := api.LspGetAllWithResponse(context.TODO())
				if err != nil {
					return err
				}
				if res.ApplicationproblemJSONDefault != nil {
					return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
				}
				lsps := *res.JSON200
				out, err := yaml.Marshal(lsps)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			} else {
				name := args[0]
				res, err := api.LspGetWithResponse(context.TODO(), name)
				if err != nil {
					return err
				}
				if res.ApplicationproblemJSONDefault != nil {
					return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
				}
				lsp := *res.JSON200
				out, err := yaml.Marshal(lsp)
				if err != nil {
					return err
				}
				fmt.Fprintf(io.Out, "%v", string(out))
			}

			return nil
		},
	}
}
