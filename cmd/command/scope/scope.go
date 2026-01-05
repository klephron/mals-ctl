package scope

import (
	"context"
	"fmt"
	"mals-ctl/cmd/runtime"
	"mals-ctl/internal/encoding/yaml"

	"github.com/spf13/cobra"
)

func NewCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scope",
		Short: "Manage scope",
	}

	cmd.AddCommand(newTreeCommand(c, io))

	return cmd
}

func newTreeCommand(c runtime.Context, io runtime.IOStreams) *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "Show scope tree",
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := c.Client()
			if err != nil {
				return err
			}
			res, err := api.ScopeTreeRootWithResponse(context.TODO())
			if err != nil {
				return err
			}
			if res.ApplicationproblemJSONDefault != nil {
				return fmt.Errorf("%v", *res.ApplicationproblemJSONDefault.Detail)
			}
			spaceRoot := *res.JSON200
			out, err := yaml.Marshal(spaceRoot)
			if err != nil {
				return err
			}
			fmt.Fprintf(io.Out, "%v", string(out))
			return nil
		},
	}
}
