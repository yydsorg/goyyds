package server

type Server interface {
	// Initialise options
	Init(...Option) error
	// Retrieve the options
	Options() Options
	// Start the server
	Start() error
	// Stop the server
	Stop() error
	// Server implementation
	String() string
	// version
	Version() string
}

var (
	DefaultServer  Server = newRpcServer()
	DefaultVersion string = "0.0.1"
	DefaultName    string = "GOYYDS"
)

type Option func(*Options)

func Init(opt ...Option) {
	if DefaultServer == nil {
		DefaultServer = newRpcServer(opt...)
	}
	_ = DefaultServer.Init(opt...)
}
