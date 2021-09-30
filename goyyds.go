package goyyds

import (
	"github.com/goyyds/goyyds/v1/src/client"
	"github.com/goyyds/goyyds/v1/src/server"
)

type Service interface {
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
	// Run the service
	Run() error
	// The service implementation
	String() string
}

type Option func(*Options)

func NewService(opts ...Option) Service {
	return newService(opts...)
}
