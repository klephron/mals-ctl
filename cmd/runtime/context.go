package runtime

import (
	"mals-ctl/internal/api"
	"mals-ctl/pkg/config"
)

type Context interface {
	Client() (api.ClientWithResponsesInterface, error)

	ConfigLoad() (*config.Config, error)
	ConfigServerAdd(name string, url string) (*config.Server, error)
	ConfigServerRemove(name string) (*config.Server, error)
	ConfigContextServerSet(name string) (bool, error)
}

type ContextOptions struct {
	ConfigPath    string
	ContextServer string
}
