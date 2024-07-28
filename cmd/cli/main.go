package main

import (
	"baby-log/internal/gateway"
	"baby-log/internal/migrations"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
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

	defer db.Close()

	migrator, err := migrations.New(db)

	if err != nil {
		panic(err)
	}

	if err := migrator.Status(); err != nil {
		panic(err)
	}
}
