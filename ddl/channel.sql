create table channel
(
    name       varchar(250)
        constraint channel_pk
            primary key,
    author_id     bigint             not null
        constraint channel_user_id_fk
            references "user",
    created_at date default now() not null,
    deleted_at date
);

