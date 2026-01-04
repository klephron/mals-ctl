package runtime

import "mals-ctl/cmd/config"

type Context interface {
	Config() (*config.Config, error)
}
