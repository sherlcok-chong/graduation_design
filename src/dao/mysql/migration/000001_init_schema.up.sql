create table if not exists user # 用户信息
(
    id       bigint primary key auto_increment,
    name     varchar(40)  not null unique,
    email    varchar(40)  not null unique,
    password varchar(128) not null,
    avatar   varchar(128) not null,
    sign     varchar(128) not null default '',
    gender   varchar(8)   not null default '',
    birthday varchar(20)  not null default '',
    address  varchar(20)
) default charset = utf8mb4;

create table if not exists file #媒体文件
(
    id        bigint primary key auto_increment,
    filename  varchar(255) not null,
    file_key  varchar(255) not null,
    url       varchar(266) not null,
    userid    bigint       not null,
    create_at timestamp    not null default now()
);

create table if not exists commodity #商品
(
    id      bigint primary key auto_increment,
    name    varchar(20)  not null,
    user_id bigint       not null,
    price   varchar(20)  not null,
    texts   varchar(255) not null,
    is_free bool         not null,
    is_lend bool         not null
);

create index user_com on commodity (user_id);
-- 获取商品的自增主键值
SET @commodity_id = LAST_INSERT_ID();
create table if not exists commodity_media #商品媒体
(
    id           bigint primary key auto_increment,
    commodity_id bigint not null,
    file_id      bigint not null
);

create index comm on commodity_media (commodity_id);

-- 标签表
CREATE TABLE tags
(
    tag_id   BIGINT PRIMARY KEY auto_increment,
    tag_name VARCHAR(50) not null
);
SET @commodity_id = LAST_INSERT_ID();
-- 关系表，用于关联商品和标签
CREATE TABLE product_tags
(
    product_id bigint not null,
    tag_id     BIGINT not null,
    PRIMARY KEY (product_id, tag_id)
);

-- 评论表
create table if not exists comment
(
    id         bigint primary key not null auto_increment,
    user_id    bigint             not null,
    product_id bigint             not null,
    texts      varchar(255)       not null default ''
);

create table if not exists comment_media #评论媒体
(
    id         bigint primary key not null key auto_increment,
    comment_id bigint             not null,
    file_id    bigint             not null
);

create table if not exists orders
(
    id              bigint primary key not null auto_increment,
    lend_user_id    bigint             not null,
    borrow_user_id  bigint             not null,
    product_id      bigint             not null,
    unit_price      varchar(10)        not null,            #单日价格
    total_price     varchar(10)        not null,            #租赁总价
    completion_time varchar(10)        not null default '', #完成时间 归还
    product_status  int                not null default 0,  #当期状态 0，商家邮递中 1 租赁中 2 归还中 3 已归还
    express_number  varchar(20)        not null default ' ',
    start_time      varchar(10)        not null,
    end_time        varchar(10)        not null
);
create table if not exists user_like
(
    id         bigint primary key not null auto_increment,
    user_id    bigint             not null,
    product_id bigint             not null
);

create index user_like_user_id on user_like (user_id);

create table if not exists lend_time
(
    id         bigint primary key not null auto_increment,
    product_id bigint             not null,
    start_time varchar(10)        not null default '2020-11-22',
    end_time   varchar(10)        not null default '2020-11-22'
);

create index lend_time_product on lend_time (product_id);

create table if not exists message
(
    id       bigint primary key not null auto_increment,
    fid      bigint             not null,
    tid      bigint             not null,
    is_file  bool               not null default false,
    is_read  bool               not null default false,
    texts    varchar(255)       not null,
    createAt datetime           not null default now()
);
create index fid_read_message on message (fid, is_read);





