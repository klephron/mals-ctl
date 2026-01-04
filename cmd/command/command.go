package command

import (
	"mals-ctl/cmd/command/config"
	"mals-ctl/cmd/command/listener"
	"mals-ctl/cmd/command/log"
	"mals-ctl/cmd/command/lsp"
	"mals-ctl/cmd/command/model"
	"mals-ctl/cmd/command/scope"
	"mals-ctl/cmd/command/usage"
	"mals-ctl/cmd/runtime"
	"mals-ctl/pkg/info"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Options struct {
	ConfigPath    string
	ContextServer string
}

func NewCommand() *cobra.Command {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		home, _ := os.UserHomeDir()
		configDir = filepath.Join(home, ".config")
	}
	configPath := filepath.Join(configDir, info.CtlName, "config.toml")

	o := Options{
		ConfigPath: configPath,
	}

	io := runtime.IOStreams{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}

	cmd := &cobra.Command{
		Use:   info.CtlName,
		Short: info.CtlDescriptionShort,
	}

	cmd.SetIn(io.In)
	cmd.SetOut(io.Out)
	cmd.SetErr(io.Err)

	cmd.PersistentFlags().StringVarP(&o.ConfigPath, "config", "c", o.ConfigPath, "Path to the config file")
	cmd.PersistentFlags().StringVar(&o.ContextServer, "context-server", o.ContextServer, "Context server override")

	c := newContext(&o)

	cmd.AddCommand(config.NewCommand(c, io))
	cmd.AddCommand(listener.NewCommand(c, io))
	cmd.AddCommand(log.NewCommand(c, io))
	cmd.AddCommand(lsp.NewCommand(c, io))
	cmd.AddCommand(model.NewCommand(c, io))
	cmd.AddCommand(scope.NewCommand(c, io))
	cmd.AddCommand(usage.NewCommand(c, io))

	return cmd
}
