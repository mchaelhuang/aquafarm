-- +goose Up
create table pond
(
    id         serial
        constraint pond_pk
            primary key,
    farm_id    integer                 not null
        constraint pond_farm_id_fk
            references farm
            on delete cascade,
    label      varchar                 not null,
    volume     integer                 not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp
);

create unique index pond_farm_id_label_unique
    on pond (farm_id, label)
    where deleted_at is null;

-- +goose Down
drop table if exists pond cascade;