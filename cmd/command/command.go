package command

import (
	"mals-ctl/cmd/command/config"
	cfg "mals-ctl/cmd/config"
	"mals-ctl/cmd/runtime"
	"mals-ctl/pkg/info"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Options struct {
	ConfigFile string
}

type context struct {
	runtime.Context
	options *Options
}

func (s *context) Config() (*cfg.Config, error) {
	viper.SetConfigFile(s.options.ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg cfg.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewCommand() *cobra.Command {
	o := Options{}
	iostreams := runtime.IOStreams{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}

	cmd := &cobra.Command{
		Use:   info.CtlName,
		Short: info.CtlDescriptionShort,
	}

	cmd.PersistentFlags().StringVarP(&o.ConfigFile, "config", "c", o.ConfigFile, "Path to the config file.")

	c := &context{
		options: &o,
	}

	cmd.AddCommand(config.NewCommand(c, iostreams))

	return cmd
}
