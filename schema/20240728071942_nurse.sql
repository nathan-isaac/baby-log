-- +goose Up
-- +goose StatementBegin
CREATE TABLE nurse
(
    nurse_id     TEXT PRIMARY KEY,
    person_id   TEXT        NOT NULL,
    right_duration_seconds INTEGER NOT NULL,
    left_duration_seconds INTEGER NOT NULL,
    right_volume_ml   INTEGER     NOT NULL,
    left_volume_ml   INTEGER     NOT NULL,
    started_at  TIMESTAMP   NOT NULL,
    ended_at    TIMESTAMP,
    notes       TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,

    FOREIGN KEY (person_id) REFERENCES person (person_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table nurse;
-- +goose StatementEnd
