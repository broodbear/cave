package main

import (
	"log"
	"os"

	"github.com/broodbear/cave/cmd/commands"
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
				Value: "default.db",
			},
		},
		Commands: []*cli.Command{
			commands.Credentials(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
