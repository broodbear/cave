package commands

import (
	"github.com/broodbear/cave/internal/datastore"
	"github.com/broodbear/cave/internal/services"
	"github.com/urfave/cli/v2"
)

var fileFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "separator",
		Usage: "character(s) to separate values",
		Value: ",",
	},
	&cli.StringFlag{
		Name:  "filename",
		Usage: "where to write the data",
		Value: "cave.csv",
	},
}

var (
	credentialsAddCommand = &cli.Command{
		Name:  "add",
		Usage: "add new credentials",
		Action: func(c *cli.Context) error {
			creds, err := wireCredentials(c.String("project"), c.String("database"))
			if err != nil {
				return err
			}

			return creds.Add()
		},
	}

	credentialsImportCommand = &cli.Command{
		Name:  "import",
		Usage: "import existing credentials from a file",
		Flags: fileFlags,
		Action: func(c *cli.Context) error {
			creds, err := wireCredentials(c.String("project"), c.String("database"))
			if err != nil {
				return err
			}

			return creds.Import(c.String("filename"), c.String("separator"))
		},
	}

	credentialsExportCommand = &cli.Command{
		Name:  "export",
		Usage: "export credentials to a file",
		Flags: fileFlags,
		Action: func(c *cli.Context) error {
			creds, err := wireCredentials(c.String("project"), c.String("database"))
			if err != nil {
				return err
			}

			return creds.Export(c.String("filename"), c.String("separator"))
		},
	}

	credentialsPrintCommand = &cli.Command{
		Name:  "print",
		Usage: "print the credentials to the screen",
		Action: func(c *cli.Context) error {
			creds, err := wireCredentials(c.String("project"), c.String("database"))
			if err != nil {
				return err
			}

			return creds.Print()
		},
	}
)

func Credentials() *cli.Command {
	return &cli.Command{
		Name:    "credentials",
		Aliases: []string{"c"},
		Usage:   "manage known credentials",
		Action: func(c *cli.Context) error {
			cli.ShowSubcommandHelpAndExit(c, 0)

			return nil
		},
		Subcommands: []*cli.Command{
			credentialsAddCommand,
			credentialsImportCommand,
			credentialsExportCommand,
			credentialsPrintCommand,
		},
	}
}

func wireCredentials(project, database string) (services.Credentials, error) {
	db, err := datastore.NewDatastore(database)
	if err != nil {
		return services.Credentials{}, err
	}

	credStore := datastore.NewCredentials(db)

	credService := services.NewCredentials(project, credStore)

	return credService, nil
}
