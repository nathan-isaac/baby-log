package migrations

import (
	"baby-log/schema"
	"database/sql"
	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db *sql.DB
}

func New(db *sql.DB) (Migrator, error) {
	return Migrator{db: db}, nil
}

func (it Migrator) Up() error {
	goose.SetBaseFS(schema.EmbedMigrations)

	if err := goose.SetDialect("turso"); err != nil {
		return err
	}

	return goose.Up(it.db, ".")
}

func (it Migrator) Status() error {
	goose.SetBaseFS(schema.EmbedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}
	return goose.Status(it.db, ".")
}

func (it Migrator) Create(name string) error {
	goose.SetBaseFS(schema.EmbedMigrations)

	if err := goose.SetDialect("turso"); err != nil {
		return err
	}

	return goose.Create(it.db, "./schema", name, "sql")
}
