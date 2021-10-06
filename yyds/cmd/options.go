package cmd

import (
	"context"
	"github.com/urfave/cli/v2"
	"log"
)

type Options struct {
	// Name of the application
	Name string
	// Description of the application
	Description string
	// Version of the application
	Version string
	// Action to execute when Run is called and there is no subcommand
	Action func(*cli.Context) error
	// TODO replace with built in command definition
	Commands []*cli.Command
	// TODO replace with built in flags definition
	Flags []cli.Flag
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type Option func(options *Options)

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func beforeFromContext(ctx context.Context, def cli.BeforeFunc) cli.BeforeFunc {

	log.Println("beforeFromContext done")
	return func(c *cli.Context) error {
		return nil
	}
}
