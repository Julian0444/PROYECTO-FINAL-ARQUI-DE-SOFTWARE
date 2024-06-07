create table categories
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime(3) null,
    updated_at    datetime(3) null,
    deleted_at    datetime(3) null,
    category_name longtext    null
);

create index idx_categories_deleted_at
    on categories (deleted_at);

