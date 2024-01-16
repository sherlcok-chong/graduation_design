create table if not exists user # 用户信息
(
    id       bigint primary key auto_increment,
    name     varchar(40)  not null unique,
    email    varchar(40)  not null unique,
    password varchar(128) not null,
    sign     varchar(128) not null,
    gender   varchar(8)   not null,
    birthday varchar(20)  not null
) default charset = utf8;

