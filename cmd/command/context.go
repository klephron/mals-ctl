package command

import (
	"mals-ctl/cmd/config"
	"mals-ctl/cmd/runtime"
)

type context struct {
	runtime.Context
	options *Options
	store   *config.Store
}

func newContext(options *Options) *context {
	return &context{
		options: options,
		store:   config.NewStore(options.ConfigPath),
	}
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
