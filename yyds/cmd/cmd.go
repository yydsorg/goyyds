package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var (
	DefaultCmd  Cmd = New()
	version         = "1.0.0"
	name            = "yyds"
	description     = "yyds is a client for goyyds,use `yyds [command] --help` to see more detail"
	defaultFlag     = []cli.Flag{
		&cli.StringFlag{
			Name:    "init",
			Usage:   "initialize a goyyds frame",
			EnvVars: nil,
		},
		&cli.StringFlag{
			Name:    "create",
			Usage:   "create a goyyds frame",
			EnvVars: nil,
		},
		&cli.StringFlag{
			Name:    "remove",
			Usage:   "remove a goyyds frame",
			EnvVars: nil,
		},
	}
)

type Cmd interface {
	// Init initialises options
	// Note: Use Run to parse command line
	Init(opts ...Option) error
	// Options set within this command
	Options() Options
	// The cli app within this cmd
	App() *cli.App
	// Run executes the command
	Run() error
	// Implementation
	String() string
}

func action(c *cli.Context) error {

	fmt.Println("action done")
	return nil
}

type command struct {
	opts Options
	app  *cli.App

	// before is a function which should
	// be called in Before if not nil
	before cli.ActionFunc

	// indicates whether this is a service
	service bool
}

func (c *command) Init(opts ...Option) error {
	for _, o := range opts {
		o(&c.opts)
	}

	if len(c.opts.Name) > 0 {
		c.app.Name = c.opts.Name
	}
	if len(c.opts.Version) > 0 {
		c.app.Version = c.opts.Version
	}
	if len(c.opts.Description) > 0 {
		c.app.Usage = c.opts.Description
	}
	c.app.HideVersion = len(c.opts.Version) == 0

	if len(c.opts.Flags) > 0 {
		c.app.Flags = append(c.app.Flags, c.opts.Flags...)
	}
	if c.opts.Action != nil {
		c.app.Action = c.opts.Action
	}
	return nil
}

func (c *command) Options() Options {
	panic("implement me")
}

func (c *command) App() *cli.App {

	return c.app
}

func (c *command) Run() error {
	defer func() {
		if r := recover(); r != nil {
			// send report TODO
			//panic(r)
			log.Println(r)
		}
	}()

	return c.app.Run(os.Args)
}

func (c *command) String() string {
	panic("implement me")
}

func (c *command) Before(ctx *cli.Context) error {

	return nil
}

func New(opts ...Option) *command {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	cmd := new(command)
	cmd.app = cli.NewApp()
	cmd.app.Version = version
	cmd.app.Description = description
	cmd.app.Name = name
	cmd.app.Flags = defaultFlag
	cmd.app.Action = action
	cmd.app.Before = beforeFromContext(options.Context, cmd.Before)
	cmd.opts = options

	// append flags
	if len(options.Flags) > 0 {
		cmd.app.Flags = append(cmd.app.Flags, options.Flags...)
	}

	//action to replace
	if options.Action != nil {
		cmd.app.Action = options.Action
	}

	return cmd
}

func Register(cmds ...*cli.Command) {
	app := DefaultCmd.App()
	app.Commands = append(app.Commands, cmds...)

	// sort
	sort.Slice(app.Commands, func(i, j int) bool {
		return app.Commands[i].Name < app.Commands[j].Name
	})
}

func Run() {
	options := []Option{
		func(o *Options) {
			fmt.Println("options1")
		},
	}
	if err := DefaultCmd.Init(options...); err != nil {
		fmt.Println("初始化失败")
		os.Exit(1)
	}
	if err := DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
