create table "user"
(
    id         bigserial
        constraint user_pk
            primary key,
    username   varchar(100) not null unique,
    password   varchar(100) not null,
    created_at date         not null default now(),
    deleted_at date,
    pfp        varchar(64)
);