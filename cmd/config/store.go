package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Store struct {
	path string
}

func NewStore(path string) *Store {
	return &Store{
		path: path,
	}
}

func (s *Store) Load() (*Config, error) {
	viper.SetConfigFile(s.path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (s *Store) store(cfg *Config) error {
	viper.Set(KeyServers, cfg.Servers)
	viper.Set(KeyContext, cfg.Context)

	return viper.WriteConfig()
}

func (s *Store) AddServer(name, url string) error {
	cfg, err := s.Load()
	if err != nil {
		return err
	}
	for _, srv := range cfg.Servers {
		if srv.Name == name {
			return fmt.Errorf("server %q already exists", name)
		}
	}
	cfg.Servers = append(cfg.Servers, &Server{Name: name, Url: url})
	return s.store(cfg)
}

func (s *Store) RemoveServer(name string) error {
	cfg, err := s.Load()
	if err != nil {
		return err
	}
	for i, srv := range cfg.Servers {
		if srv.Name == name {
			cfg.Servers = append(cfg.Servers[:i], cfg.Servers[i+1:]...)
			return s.store(cfg)
		}
	}
	return fmt.Errorf("server %q not found", name)
}

func (s *Store) SetContextServer(name string) error {
	cfg, err := s.Load()
	if err != nil {
		return err
	}
	for _, srv := range cfg.Servers {
		if srv.Name == name {
			if cfg.Context == nil {
				cfg.Context = &Context{}
			}
			cfg.Context.Server = name
			return s.store(cfg)
		}
	}
	return fmt.Errorf("server %q not found", name)
}
