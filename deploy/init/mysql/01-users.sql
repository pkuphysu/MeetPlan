create table if not exists users
(
    id           bigint auto_increment primary key,
    create_time  timestamp    not null default current_timestamp,
    update_time  timestamp    not null default current_timestamp on update current_timestamp,
    pku_id       varchar(10)  not null unique,
    name         varchar(255) not null,
    email        varchar(255) not null,
    is_active    tinyint(1)   not null default 0,
    is_teacher   tinyint(1)   not null default 0,
    is_admin     tinyint(1)   not null default 0,
    gender       tinyint(1)   not null default 0,
    avatar       varchar(255) null,
    department   varchar(255) null,
    phone        varchar(255) null,
    major        varchar(255) null,
    grade        tinyint      null,
    dorm         varchar(255) null,
    office       varchar(255) null,
    introduction longtext     null
) engine = innodb
  default charset = utf8;