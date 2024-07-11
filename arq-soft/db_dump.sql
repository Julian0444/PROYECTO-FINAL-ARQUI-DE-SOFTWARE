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


create table courses
(
    id                 bigint unsigned auto_increment
        primary key,
    created_at         datetime(3)     null,
    updated_at         datetime(3)     null,
    deleted_at         datetime(3)     null,
    course_name        longtext        null,
    course_description longtext        null,
    course_price       double          null,
    course_duration    bigint          null,
    course_init_date   longtext        null,
    course_state       tinyint(1)      null,
    course_capacity    bigint          null,
    course_image       longtext        null,
    category_id        bigint unsigned null,
    constraint fk_courses_category
        foreign key (category_id) references categories (id)
);

create index idx_courses_deleted_at
    on courses (deleted_at);

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

create table inscriptos
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    course_id  bigint unsigned null,
    user_id    bigint unsigned null,
    constraint fk_inscriptos_course
        foreign key (course_id) references courses (id),
    constraint fk_inscriptos_user
        foreign key (user_id) references users (id)
);

create index idx_inscriptos_deleted_at
    on inscriptos (deleted_at);

