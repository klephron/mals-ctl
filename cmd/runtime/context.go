package runtime

import (
	"mals-ctl/cmd/config"
	"mals-ctl/internal/api"
)

type Context interface {
	Client() (api.ClientWithResponsesInterface, error)
	Config() (*config.Config, error)
	Store() *config.Store
}
