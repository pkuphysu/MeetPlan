create table if not exists orders
(
    id          bigint auto_increment primary key,
    create_time timestamp  not null default current_timestamp,
    update_time timestamp  not null default current_timestamp on update current_timestamp,
    status      tinyint(2) not null,
    message     longtext   null,
    plan_id     bigint     not null,
    student_id  bigint     not null,
    constraint orders_plan_id_fk_plans_id
        foreign key (plan_id) references plans (id),
    constraint orders_student_id_fk_users_id
        foreign key (student_id) references users (id)
);