package goyyds

import (
	"context"
	"github.com/goyyds/goyyds/v1/src/plugins/cmd"
	"github.com/goyyds/goyyds/v1/src/server"
	"log"
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
		Name:   server.DefaultName,
		Cmd:    cmd.DefaultCmd,
		Server: server.DefaultServer,

		Context: context.Background(),
		Signal:  true,
	}

	for i, o := range opts {
		log.Println(i)
		o(&opt)
	}
	log.Println(opt.Cmd.Options().Name)
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
