package server

type Server interface {
	Init(...Option) error
}

type Option func(*Options)

type server struct {
	Name string
}

func (s server) Init(option ...Option) error {
	panic("implement me")
}
