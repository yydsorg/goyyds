package cmd

import (
	"github.com/spf13/cobra"
)

var DefaultCmd = newCmd()

type Option func(o *Options)

type Cmd interface {
	// The cli app within this cmd
	//App() *cli.App
	// Adds options, parses flags and initialise
	// exits on error
	Init(opts ...Option) error
	// Options set within this command
	Options() Options
}

type cmd struct {
	*cobra.Command
}

func (c cmd) Init(opts ...Option) error {
	panic("implement me")
}

func (c cmd) Options() Options {
	panic("implement me")
}

func newCmd() Cmd {
	cmd := new(cmd)

	return cmd
}
