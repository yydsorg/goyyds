package config

import (
	"github.com/urfave/cli/v2"
	"github.com/yydsorg/yyds/cmd"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config a yyds frame",
			Subcommands: []*cli.Command{
				{
					Name:   "set",
					Usage:  "set config",
					Action: setConfig,
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "aaa",
							Aliases: []string{"y"},
							Usage:   "",
						},
					},
				},
			},
		})
}

func setConfig(ctx *cli.Context) error {

	return nil
}
