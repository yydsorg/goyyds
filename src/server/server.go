package server

type Server interface {
	Init(...Option) error
	Options() Options
}

var (
	DefaultServer Server = newRpcServer()
)

type Option func(*Options)

type server struct {
	Name string
}

func Init(option ...Option) error {

	return nil
}
