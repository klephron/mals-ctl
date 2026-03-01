package context

import (
	"fmt"
	"mals-ctl/cmd/runtime"
	"mals-ctl/internal/api"
	"sync"
)

type Context struct {
	mu      sync.RWMutex
	options *runtime.ContextOptions
}

func New(options *runtime.ContextOptions) *Context {
	return &Context{
		options: options,
	}
}

func (s *Context) Client() (api.ClientWithResponsesInterface, error) {
	cfg, err := s.ConfigLoad()
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
