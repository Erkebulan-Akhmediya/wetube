create table users_roles
(
    user_id   bigint      not null
        constraint users_roles_user_id_fk
            references "user",
    role_name varchar(10) not null
        constraint users_roles_role_name_fk
            references role,
    constraint users_roles_pk
        primary key (user_id, role_name)
);