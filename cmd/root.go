package main

import (
	"fmt"
	"mals-ctl/pkg/info"

	"github.com/spf13/cobra"
)

func rootNewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   info.CtlName,
		Short: info.CtlDescriptionShort,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello, world!")
		},
	}

	return cmd
}
