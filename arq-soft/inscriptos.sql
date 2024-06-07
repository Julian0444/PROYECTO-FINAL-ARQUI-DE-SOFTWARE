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

