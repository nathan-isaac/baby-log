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
	goose.SetBaseFS(schema.EmbedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		return Migrator{}, err
	}

	return Migrator{
		db: db,
	}, nil
}

func (it Migrator) Up() error {
	return goose.Up(it.db, ".")
}

func (it Migrator) Status() error {
	return goose.Status(it.db, ".")
}

func (it Migrator) Create(name string) error {
	return goose.Create(it.db, "./schema", name, "sql")
}

func (it Migrator) Version() error {
	return goose.Version(it.db, ".")
}
