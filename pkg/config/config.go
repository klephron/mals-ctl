package config

const (
	KeyServers = "servers"
	KeyContext = "context"
)

type Config struct {
	Servers []*Server `mapstructure:"servers"`
	Context *Context  `mapstructure:"context"`
}

type Server struct {
	Name string `mapstructure:"name"`
	Url  string `mapstructure:"url"`
}

type Context struct {
	Server string `mapstructure:"server"`
}
