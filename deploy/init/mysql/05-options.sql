create table options
(
    id          int auto_increment primary key,
    create_time timestamp    not null default current_timestamp,
    update_time timestamp    not null default current_timestamp on update current_timestamp,
    name        varchar(128) not null,
    value       json         not null,
    constraint options_name_uniq unique (name)
);