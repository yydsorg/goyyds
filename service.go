package goyyds

import (
	"github.com/goyyds/goyyds/v1/src/client"
	"github.com/goyyds/goyyds/v1/src/server"
	"log"
	"sync"
)

type service struct {
	opts Options

	once sync.Once
}

var (
	DefaultName = "GOYYDS"
)

func newService(opts ...Option) Service {
	s := new(service)
	options := newOptions(opts...)

	// service name
	//serviceName := options.Server.Options().Name

	log.Println(options.Server.Options().Version)

	s.opts = options
	return s
}

func (s *service) Client() client.Client {
	panic("implement me")
}

func (s *service) Server() server.Server {
	panic("implement me")
}

func (s *service) Init(opts ...Option) {
	// process options
	for _, o := range opts {
		o(&s.opts)
	}
	s.once.Do(func() {
		log.Println("111")
	})

}

func (s *service) Options() Options {
	panic("implement me")
}

func (s *service) Run() error {
	log.Println("run")

	return nil
}

func (s *service) String() string {
	panic("implement me")
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

// Name of the service
func Name(n string) Option {
	return func(o *Options) {
		_ = o.Server.Init(server.Name(n))
	}
}

// Version of the service
func Version(v string) Option {
	return func(o *Options) {
		_ = o.Server.Init(server.Version(v))
	}
}
