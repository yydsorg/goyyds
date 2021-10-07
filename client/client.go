package client

import "time"

type Client interface {
	Init(...Option) error
	Options() Options
	String() string
}

type Option func(*Options)

var (

	// DefaultClient is a default client to use out of the box
	DefaultClient Client = newRpcClient()
	// DefaultPoolSize sets the connection pool size
	DefaultPoolSize = 100
	// DefaultPoolTTL sets the connection pool ttl
	DefaultPoolTTL = time.Minute
)
