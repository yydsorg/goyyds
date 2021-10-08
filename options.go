package goyyds

import (
	"context"
	"github.com/yydsorg/goyyds/client"
	"github.com/yydsorg/goyyds/plugins/cmd"
	"github.com/yydsorg/goyyds/server"
	"github.com/yydsorg/goyyds/web"
	"log"
)

type Options struct {
	Name   string
	Genre  string
	Cmd    cmd.Cmd
	Server server.Server
	Client client.Client
	Web    web.Web

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func newOptions(opts ...Option) Options {

	opt := Options{
		Name: server.DefaultName,
		Cmd:  cmd.DefaultCmd,
		//Server: server.DefaultServer,
		//Client: client.DefaultClient,
		//Web:    web.DefaultWeb,

		Context: context.Background(),
		Signal:  true,
	}

	for _, o := range opts {
		o(&opt)
	}

	log.Println(opt.Genre)
	switch opt.Genre {
	case "web":
		opt.Web = web.DefaultWeb
	case "service":
		opt.Server = server.DefaultServer
	case "client":
		opt.Client = client.DefaultClient
	default:
		opt.Web = web.DefaultWeb
	}

	return opt
}

func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}
func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}
func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}
func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
