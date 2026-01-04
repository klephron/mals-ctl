package main

// import (
// 	"fmt"
// 	"os"
// )

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type contextKey string

const configKey contextKey = "config"

type Config struct {
	Database DatabaseConfig
	Math     MathConfig
	Debug    bool
}

type DatabaseConfig struct {
	Host string
	Port int
}

type MathConfig struct {
	A int
	B int
}

func loadConfig(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func newRootCommand() *cobra.Command {
	var cfgFile string

	cmd := &cobra.Command{
		Use:   "app",
		Short: "Example CLI application",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loadConfig(cfgFile)
			if err != nil {
				return err
			}

			ctx := context.WithValue(cmd.Context(), configKey, cfg)
			cmd.SetContext(ctx)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cfg := cmd.Context().Value(configKey).(*Config)
			fmt.Printf("Database: %s\n", cfg.Database.Host)
			fmt.Printf("Port: %d\n", cfg.Database.Port)
			fmt.Printf("Debug: %t\n", cfg.Debug)
		},
	}

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./config.toml)")
	cmd.AddCommand(newSumCommand())

	return cmd
}

func newSumCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "sum",
		Short: "Add two numbers from config",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := cmd.Context().Value(configKey).(*Config)
			fmt.Printf("%d + %d = %d\n", cfg.Math.A, cfg.Math.B, cfg.Math.A+cfg.Math.B)
		},
	}
}

func main() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// func main() {
// 	if err := rootNewCommand().Execute(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }
