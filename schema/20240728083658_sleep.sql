-- +goose Up
-- +goose StatementBegin
create table sleep
(
    sleep_id    text primary key,
    person_id   text        not null,
    started_at  timestamp   not null,
    ended_at    timestamp   not null,
    notes       text        not null,
    created_at  timestamp   not null,
    updated_at  timestamp   not null,

    foreign key (person_id) references person(person_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table sleep;
-- +goose StatementEnd
