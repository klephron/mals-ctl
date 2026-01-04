package config

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type Store struct {
	mu   sync.RWMutex
	path string
}

func NewStore(path string) *Store {
	return &Store{path: path}
}

func (s *Store) ensure() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	if _, err := os.Stat(s.path); os.IsNotExist(err) {
		f, err := os.Create(s.path)
		if err != nil {
			return err
		}
		f.Close()
	}

	return nil
}

func (s *Store) Load() (*Config, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

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
	s.mu.Lock()
	defer s.mu.Unlock()

	viper.Set(KeyServers, cfg.Servers)
	viper.Set(KeyContext, cfg.Context)

	return viper.WriteConfig()
}

func (s *Store) AddServer(name, url string) (*Server, error) {
	cfg, err := s.Load()
	if err != nil {
		return nil, err
	}

	for _, srv := range cfg.Servers {
		if srv.Name == name {
			return nil, nil
		}
	}

	server := &Server{Name: name, Url: url}

	cfg.Servers = append(cfg.Servers, server)

	return server, s.store(cfg)
}

func (s *Store) RemoveServer(name string) (*Server, error) {
	cfg, err := s.Load()
	if err != nil {
		return nil, err
	}

	for i, srv := range cfg.Servers {
		if srv.Name == name {
			cfg.Servers = append(cfg.Servers[:i], cfg.Servers[i+1:]...)
			return srv, s.store(cfg)
		}
	}
	return nil, nil
}

func (s *Store) SetContextServer(name string) (bool, error) {
	if err := s.ensure(); err != nil {
		return false, err
	}

	cfg, err := s.Load()
	if err != nil {
		return false, err
	}

	for _, srv := range cfg.Servers {
		if srv.Name == name {
			if cfg.Context == nil {
				cfg.Context = &Context{}
			}
			cfg.Context.Server = name
			return true, s.store(cfg)
		}
	}
	return false, nil
}
