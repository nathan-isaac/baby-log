package main

import (
	"baby-log/internal/gateway"
	"baby-log/internal/migrations"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()

	if err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}

	db, err := gateway.NewConnection()

	if err != nil {
		panic(err)
	}

	migrator, err := migrations.New(db)

	if err != nil {
		panic(err)
	}

	app := &cli.App{
		Name:  "cli",
		Usage: "CLI for the baby-log application",
		Action: func(*cli.Context) error {
			fmt.Println("Welcome to the baby-log CLI")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "migration",
				Aliases: []string{"m"},
				Usage:   "manage migrations",
				Subcommands: []*cli.Command{
					{
						Name:  "status",
						Usage: "show the status of the migrations",
						Action: func(c *cli.Context) error {
							return migrator.Status()
						},
					},
					{
						Name:  "create",
						Usage: "create a new migration",
						Action: func(c *cli.Context) error {
							return migrator.Create(c.Args().First())
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
