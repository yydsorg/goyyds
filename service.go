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

	return nil
}

func (s *service) String() string {
	panic("implement me")
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}
