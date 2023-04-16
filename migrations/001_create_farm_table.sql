-- +goose Up
create table public.farm
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

alter table farm
    add constraint farm_name_location_unique
        unique (name, location);

-- +goose Down
drop table if exists farm cascade;