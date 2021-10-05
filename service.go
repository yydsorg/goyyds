package goyyds

import (
	"github.com/goyyds/goyyds/v1/src/client"
	"github.com/goyyds/goyyds/v1/src/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type service struct {
	opts Options

	once sync.Once
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

func (s *service) Init(opts ...Option) {
	// process options

	for _, o := range opts {
		o(&s.opts)
	}
	s.once.Do(func() {
		log.Println("init do once")
	})

}

func (s *service) Options() Options {
	panic("implement me")
}

func (s *service) Client() client.Client {
	panic("implement me")
}

func (s *service) Server() server.Server {
	panic("implement me")
}

func (s *service) Run() error {
	log.Println("run")
	if err := s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)

	if s.opts.Signal {
		sos := []os.Signal{
			syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
		}
		signal.Notify(ch, sos...)
	}
	log.Println("waiting for signal")
	select {
	case <-ch:
	case <-s.opts.Context.Done():
	}

	return s.Stop()
}

func (s *service) String() string {
	panic("implement me")
}

func (s *service) Version() string {
	return s.opts.Server.Version()
}

func (s *service) Start() error {
	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := s.opts.Server.Start(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Stop() error {
	for _, fn := range s.opts.BeforeStop {
		if err := fn(); err != nil {
			return err
		}
	}

	if err := s.opts.Server.Stop(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStop {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}

func newService(opts ...Option) Service {
	s := new(service)
	options := newOptions(opts...)

	// service name
	//serviceName := options.Server.Options().Name

	s.opts = options
	return s
}

// Name of the service
func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(server.Name(n))
	}
}
