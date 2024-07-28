-- +goose Up
-- +goose StatementBegin
create table weight
(
    weight_id   text primary key,
    person_id   text        not null,
    weight_ounces   integer        not null,
    occurred_at timestamp   not null,
    notes       TEXT        NOT NULL,
    created_at  timestamp   not null,
    updated_at  timestamp   not null,

    foreign key (person_id) references person(person_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table weight;
-- +goose StatementEnd
