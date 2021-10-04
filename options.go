package goyyds

import (
	"context"
	"github.com/goyyds/goyyds/v1/src/plugins/cmd"
	"github.com/goyyds/goyyds/v1/src/server"
)

type Options struct {
	Name   string
	Cmd    cmd.Cmd
	Server server.Server

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
		Name:   DefaultName,
		Cmd:    cmd.DefaultCmd,
		Server: server.DefaultServer,

		Context: context.Background(),
		Signal:  false,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt

}
