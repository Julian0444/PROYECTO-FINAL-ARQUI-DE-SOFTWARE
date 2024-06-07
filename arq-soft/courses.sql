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

