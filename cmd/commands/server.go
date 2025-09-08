package commands

import (
	"github.com/urfave/cli/v2"
	"go-gin-gorm-starter/cmd/run"
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
