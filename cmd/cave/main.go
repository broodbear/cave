package main

import (
	"log"
	"os"

	"github.com/broodbear/cave/cmd/cave/commands"
	"github.com/broodbear/cave/cmd/cave/config"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cave",
		Usage: "keep track of your targets",
		Action: func(c *cli.Context) error {

			cli.ShowAppHelpAndExit(c, 0)

			return nil
		},
		Suggest: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "project",
				Usage: "name of the project",
				Value: "default",
			},
			&cli.StringFlag{
				Name:  "database",
				Usage: "path to the sqlite3 database",
				Value: config.GetConfigPath() + "/default.db",
			},
		},
		Commands: []*cli.Command{
			commands.Initialize(),
			commands.Credentials(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
