-- +goose Up
-- +goose StatementBegin
CREATE TABLE person (
  person_id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  birth_date DATE,
  birth_time TIME,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE person;
-- +goose StatementEnd
