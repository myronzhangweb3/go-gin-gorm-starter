package commands

import (
	"go-gin-gorm-starter/cmd/run"

	"github.com/urfave/cli/v2"
)

var RunCommand = &cli.Command{
	Name:  "server",
	Usage: "Start server service",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Value:   "config/config.toml",
			Usage:   "config file path",
		},
	},
	Action: func(ctx *cli.Context) error {
		return run.Server(ctx)
	},
}
