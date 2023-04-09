create table if not exists plans
(
    id          bigint auto_increment primary key,
    create_time timestamp    not null default current_timestamp,
    update_time timestamp    not null default current_timestamp on update current_timestamp,
    teacher_id  bigint       not null,
    start_time  timestamp    not null,
    duration    bigint       not null,
    place       varchar(255) not null,
    quota       tinyint      not null,
    message     longtext     null,
    constraint plans_teacher_id_fk_users_id
        foreign key (teacher_id) references users (id)
);

create index plans_start_time_index on plans (start_time);