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
			Name:  "get",
			Usage: "get information",
			Subcommands: []*cli.Command{
				{
					Name:   "config",
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
					Name:   "info",
					Usage:  "get info",
					Action: getInfo,
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "namespace",
							Aliases: []string{"n"},
							Usage:   "set a namespace",
						},
					},
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
func getInfo(ctx *cli.Context) error {
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
