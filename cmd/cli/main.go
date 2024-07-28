package main

import (
	"baby-log/internal/gateway"
	"baby-log/internal/migrations"
	"baby-log/schema"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()

	if err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}

	names, err := schema.ListMigrations()

	if err != nil {
		logger.Error("failed to list migrations", "error", err)
		os.Exit(1)
	}

	logger.Info("migrations", "names", names)

	db, err := gateway.NewConnection()

	if err != nil {
		panic(err)
	}

	migrator, err := migrations.New(db)

	if err != nil {
		panic(err)
	}

	if err := migrator.Status(); err != nil {
		panic(err)
	}
}
