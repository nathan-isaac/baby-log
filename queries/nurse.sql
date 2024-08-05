-- name: CreateNurse :exec
INSERT INTO nurse (nurse_id, person_id, right_duration_seconds, left_duration_seconds, right_volume_ml, left_volume_ml, started_at, ended_at, notes, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: FilterNurse :many
SELECT *
FROM nurse
WHERE person_id = ?
ORDER BY started_at DESC;

-- name: FindNurse :one
SELECT *
FROM nurse
WHERE nurse_id = ?;
