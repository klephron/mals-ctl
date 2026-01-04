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
	configPath := filepath.Join(configDir, info.CtlName, "config.yaml")

	o := Options{
		ConfigPath: configPath,
	}

	iostreams := runtime.IOStreams{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}

	cmd := &cobra.Command{
		Use:   info.CtlName,
		Short: info.CtlDescriptionShort,
	}

	cmd.PersistentFlags().StringVarP(&o.ConfigPath, "config", "c", o.ConfigPath, "Path to the config file")
	cmd.PersistentFlags().StringVar(&o.ContextServer, "context-server", o.ContextServer, "Context server override")

	c := newContext(&o)

	cmd.AddCommand(config.NewCommand(c, iostreams))
	cmd.AddCommand(listener.NewCommand(c, iostreams))
	cmd.AddCommand(log.NewCommand(c, iostreams))
	cmd.AddCommand(lsp.NewCommand(c, iostreams))
	cmd.AddCommand(model.NewCommand(c, iostreams))
	cmd.AddCommand(scope.NewCommand(c, iostreams))
	cmd.AddCommand(usage.NewCommand(c, iostreams))

	return cmd
}
