create table if not exists friend_links
(
    id          int auto_increment primary key,
    create_time timestamp    not null default current_timestamp,
    update_time timestamp    not null default current_timestamp on update current_timestamp,
    name        varchar(32)  not null,
    url         varchar(200) not null,
    description longtext     null
) engine = innodb
  default charset = utf8mb4;

