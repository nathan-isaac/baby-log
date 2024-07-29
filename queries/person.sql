-- name: CreatePerson :exec
INSERT INTO person (person_id, name, birth_date, birth_time, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: FilterPeople :many
SELECT *
FROM person
ORDER BY name;

-- name: FindPerson :one
SELECT *
FROM person
WHERE person_id = ?;
