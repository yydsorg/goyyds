package config

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yydsorg/yyds/cmd"
	"log"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config a yyds frame",
			Subcommands: []*cli.Command{
				{
					Name:   "get",
					Usage:  "get config",
					Action: getConfig,
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "namespace",
							Aliases: []string{"ns"},
							Usage:   "get a namespace",
						},
						&cli.BoolFlag{
							Name:    "name",
							Aliases: []string{"n"},
							Usage:   "get a name",
						},
					},
				},
				{
					Name:   "set",
					Usage:  "set config",
					Action: setConfig,
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "namespace",
							Aliases: []string{"n"},
							Usage:   "set a namespace",
						},
					},
				},
				{
					Name:   "del",
					Usage:  "del config",
					Action: delConfig,
				},
			},
		})
}

func getConfig(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() == 0 {
		return cli.ShowSubcommandHelp(ctx)
	}
	key := args.Get(0)
	if len(key) == 0 {
		return fmt.Errorf("key cannot be blank")
	}
	log.Println("key:", key)

	return nil
}
func setConfig(ctx *cli.Context) error {
	args := ctx.Args()

	key := args.Get(0)
	value := args.Get(1)

	log.Println("key:", key, "value:", value)
	return nil
}
func delConfig(ctx *cli.Context) error {

	return nil
}
