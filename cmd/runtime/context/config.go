package context

import (
	"mals-ctl/pkg/config"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func (s *Context) configEnsure() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	dir := filepath.Dir(s.options.ConfigPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	if _, err := os.Stat(s.options.ConfigPath); os.IsNotExist(err) {
		f, err := os.Create(s.options.ConfigPath)
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}

func (s *Context) configStore(cfg *config.Config) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	viper.Set(config.KeyServers, cfg.Servers)
	viper.Set(config.KeyContext, cfg.Context)

	return viper.WriteConfig()
}

func (s *Context) ConfigLoad() (*config.Config, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	viper.SetConfigFile(s.options.ConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (s *Context) ConfigServerAdd(name, url string) (*config.Server, error) {
	cfg, err := s.ConfigLoad()
	if err != nil {
		return nil, err
	}

	for _, srv := range cfg.Servers {
		if srv.Name == name {
			return nil, nil
		}
	}

	server := &config.Server{Name: name, Url: url}

	cfg.Servers = append(cfg.Servers, server)

	return server, s.configStore(cfg)
}

func (s *Context) ConfigServerRemove(name string) (*config.Server, error) {
	cfg, err := s.ConfigLoad()
	if err != nil {
		return nil, err
	}

	for i, srv := range cfg.Servers {
		if srv.Name == name {
			cfg.Servers = append(cfg.Servers[:i], cfg.Servers[i+1:]...)
			return srv, s.configStore(cfg)
		}
	}
	return nil, nil
}

func (s *Context) ConfigContextServerSet(name string) (bool, error) {
	if err := s.configEnsure(); err != nil {
		return false, err
	}

	cfg, err := s.ConfigLoad()
	if err != nil {
		return false, err
	}

	for _, srv := range cfg.Servers {
		if srv.Name == name {
			if cfg.Context == nil {
				cfg.Context = &config.Context{}
			}
			cfg.Context.Server = name
			return true, s.configStore(cfg)
		}
	}
	return false, nil
}
