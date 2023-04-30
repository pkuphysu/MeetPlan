create table if not exists update_records
(
    id          int auto_increment primary key,
    create_time timestamp    not null default current_timestamp,
    update_time timestamp    not null default current_timestamp on update current_timestamp,
    time        timestamp    not null,
    author      varchar(512) not null,
    url         varchar(512) not null,
    info        longtext     not null
) engine = innodb
  default charset = utf8mb4;

