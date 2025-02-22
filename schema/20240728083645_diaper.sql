-- +goose Up
-- +goose StatementBegin
CREATE TABLE diaper
(
    diaper_id   TEXT PRIMARY KEY,
    person_id   TEXT        NOT NULL,
    type TEXT        NOT NULL CHECK ( type IN ('PEE', 'POO', 'MIX') ),
    occurred_at TIMESTAMP   NOT NULL,
    notes       TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,

    foreign key (person_id) references person(person_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table diaper;
-- +goose StatementEnd
