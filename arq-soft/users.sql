create table users
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    password   longtext    null,
    email      longtext    null,
    role       bigint      null,
    name       longtext    null,
    avatar     longtext    null
);

create index idx_users_deleted_at
    on users (deleted_at);

