package goyyds

import (
	"github.com/yydsorg/goyyds/client"
	"github.com/yydsorg/goyyds/server"
	"github.com/yydsorg/goyyds/web"
)

type YYDS interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() client.Client
	// Server is for handling requests and events
	Server() server.Server
	// web
	Web() web.Web
	// Run the service
	Run() error
	// The service implementation
	String() string
	// version
	Version() string
}

type Option func(*Options)

func NewYYDS(opts ...Option) YYDS {
	return newService(opts...)
}
