package init

import (
	"github.com/urfave/cli/v2"
	"github.com/yydsorg/yyds/cmd"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init a yyds frame",
			Subcommands: []*cli.Command{
				{
					Name:   "go",
					Usage:  "init a goyyds frame",
					Action: initialize,
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

func initialize(ctx *cli.Context) error {

	return nil
}
