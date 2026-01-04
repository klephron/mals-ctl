package config

type Config struct {
	Servers []*Server
	Context *Context
}

type Server struct {
	Name string
	Url  string
}

type Context struct {
	Server string
}
