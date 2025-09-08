package main

import (
	"go-gin-gorm-starter/cmd/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	appCli := &cli.App{
		Name:  "server",
		Usage: "server",
		Commands: []*cli.Command{
			commands.RunCommand,
		},
	}

	if err := appCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
