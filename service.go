package goyyds

import (
	"github.com/goyyds/goyyds/v1/src/client"
	"github.com/goyyds/goyyds/v1/src/server"
	"sync"
)

type service struct {
	opts Options

	once sync.Once
}

var (
	DefaultName = "GOYYDS"
)

func (s *service) Client() client.Client {
	panic("implement me")
}

func (s *service) Server() server.Server {
	panic("implement me")
}

func (s *service) Init(option ...Option) {
	panic("implement me")
}

func (s *service) Options() Options {
	panic("implement me")
}

func (s *service) Run() error {

	panic("implement me")
}

func (s *service) String() string {
	panic("implement me")
}

func (s *service) Name() string {
	return s.opts.Name
}
