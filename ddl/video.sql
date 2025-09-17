create table video
(
    name        varchar(100) not null,
    description varchar(500),
    id          bigserial
        constraint video_pk
            primary key,
    channel_id  integer      not null
        constraint video_channel_id_fk
            references channel,
    file        varchar(64)  not null
);