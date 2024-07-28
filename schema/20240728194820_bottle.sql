-- +goose Up
-- +goose StatementBegin
CREATE TABLE bottle
(
    bottle_id   TEXT PRIMARY KEY,
    person_id   TEXT        NOT NULL,
    bottle_type TEXT        NOT NULL CHECK ( bottle_type IN ('BREAST', 'FORMULA') ),
    volume_ml   INTEGER     NOT NULL,
    occurred_at TIMESTAMP   NOT NULL,
    notes       TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,

    foreign key (person_id) references person(person_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table bottle;
-- +goose StatementEnd
