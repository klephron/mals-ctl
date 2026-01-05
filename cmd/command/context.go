package command

import (
	"fmt"
	"mals-ctl/cmd/config"
	"mals-ctl/internal/api"
)

type context struct {
	options *Options
	store   *config.Store
}

func newContext(options *Options) *context {
	return &context{
		options: options,
		store:   config.NewStore(options.ConfigPath),
	}
}

func (s *context) Client() (api.ClientWithResponsesInterface, error) {
	cfg, err := s.Config()
	if err != nil {
		return nil, err
	}

	if cfg.Context == nil || cfg.Context.Server == "" {
		return nil, fmt.Errorf("context server is not specified")
	}

	for _, server := range cfg.Servers {
		if server.Name == cfg.Context.Server {
			return api.NewClientWithResponses(server.Url)
		}
	}

	return nil, fmt.Errorf("context server %q is not present", cfg.Context.Server)
}

func (s *context) Config() (*config.Config, error) {
	cfg, err := s.store.Load()
	if err != nil {
		return nil, err
	}

	if s.options.ContextServer != "" {
		if cfg.Context == nil {
			cfg.Context = &config.Context{}
		}
		cfg.Context.Server = s.options.ContextServer
	}

	return cfg, nil
}

func (s *context) Store() *config.Store {
	return s.store
}
