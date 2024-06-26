package commands

import (
	"github.com/broodbear/cave/cmd/cave/config"
	"github.com/broodbear/cave/internal/datastore"
	"github.com/urfave/cli/v2"
)

func Initialize() *cli.Command {
	return &cli.Command{
		Name:    "initialize",
		Aliases: []string{"i"},
		Usage:   "initialize the database",
		Action: func(c *cli.Context) error {
			err := datastore.Migrate(config.GetConfigPath(), "default.db")
			if err != nil {
				return err
			}

			return nil
		},
	}
}
