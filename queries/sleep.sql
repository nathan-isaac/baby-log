-- name: CreateSleep :exec
INSERT INTO sleep (sleep_id, person_id, started_at, ended_at, notes, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: FilterSleep :many
SELECT *
FROM sleep
WHERE person_id = ?
ORDER BY started_at DESC;

-- name: FindSleep :one
SELECT *
FROM sleep
WHERE sleep_id = ?;
