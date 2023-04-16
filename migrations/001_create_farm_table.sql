-- +goose Up
create table farm
(
    id         serial
        constraint farm_pk
            primary key,
    name       varchar                 not null,
    location   varchar                 not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp
);

create unique index farm_name_location_unique
    on farm (name, location)
    where deleted_at is null;

-- +goose Down
drop table if exists farm cascade;