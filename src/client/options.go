package client

import (
	"context"
	"time"
)

type Options struct {
	// Used to select codec
	ContentType string

	// Connection Pool
	PoolSize int
	PoolTTL  time.Duration

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

func newOptions(options ...Option) Options {

	opts := Options{
		PoolSize: DefaultPoolSize,
		PoolTTL:  DefaultPoolTTL,
		Context:  context.Background(),
	}
	for _, o := range options {
		o(&opts)
	}

	return opts
}
