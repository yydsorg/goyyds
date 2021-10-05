package cmd

import (
	cmd2 "github.com/goyyds/goyyds/v1/plugins/cmd/cmd"
	"github.com/spf13/cobra"
	"log"
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

	Run() error
}

type cmd struct {
	opts    Options
	RootCmd *cobra.Command
}

func (c *cmd) Run() error {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Printf("%s\n", err)
	//	}
	//}()
	for _, s := range c.opts.Command {
		log.Println("命令:", s)
		cmd2.Execute()
	}
	return nil
}

func (c *cmd) Init(opts ...Option) error {

	for _, o := range opts {
		o(&c.opts)
	}

	return nil
}

func (c *cmd) Options() Options {
	return c.opts
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

func Description(d string) Option {
	return func(o *Options) {
		o.Description = d
	}
}

func Command(c []string) Option {
	return func(o *Options) {
		o.Command = append(o.Command, c...)
	}
}

func newCmd() Cmd {
	cmd := new(cmd)
	var options []Option
	options = append(options, Name("goyyds:cmd"), Version("0.0.1"), Description("cmd test version"), Command([]string{"go version"}))
	_ = cmd.Init(options...)
	return cmd
}
