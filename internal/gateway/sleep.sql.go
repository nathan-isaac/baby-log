// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sleep.sql

package gateway

import (
	"context"
	"database/sql"
	"time"
)

const createSleep = `-- name: CreateSleep :exec
INSERT INTO sleep (sleep_id, person_id, started_at, ended_at, notes, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateSleepParams struct {
	SleepID   string
	PersonID  string
	StartedAt time.Time
	EndedAt   sql.NullTime
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateSleep(ctx context.Context, arg CreateSleepParams) error {
	_, err := q.db.ExecContext(ctx, createSleep,
		arg.SleepID,
		arg.PersonID,
		arg.StartedAt,
		arg.EndedAt,
		arg.Notes,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const filterSleep = `-- name: FilterSleep :many
SELECT sleep_id, person_id, started_at, ended_at, notes, created_at, updated_at
FROM sleep
WHERE person_id = ?
ORDER BY started_at DESC
`

func (q *Queries) FilterSleep(ctx context.Context, personID string) ([]Sleep, error) {
	rows, err := q.db.QueryContext(ctx, filterSleep, personID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sleep
	for rows.Next() {
		var i Sleep
		if err := rows.Scan(
			&i.SleepID,
			&i.PersonID,
			&i.StartedAt,
			&i.EndedAt,
			&i.Notes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findSleep = `-- name: FindSleep :one
SELECT sleep_id, person_id, started_at, ended_at, notes, created_at, updated_at
FROM sleep
WHERE sleep_id = ?
`

func (q *Queries) FindSleep(ctx context.Context, sleepID string) (Sleep, error) {
	row := q.db.QueryRowContext(ctx, findSleep, sleepID)
	var i Sleep
	err := row.Scan(
		&i.SleepID,
		&i.PersonID,
		&i.StartedAt,
		&i.EndedAt,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
