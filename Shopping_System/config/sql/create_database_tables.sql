drop database shopping;
Create Database If Not Exists shopping;
use shopping;
Create table users(
    rank int not null auto_increment comment '序号',
    uid varchar(30) not null comment '用户id',
    username varchar(50) not null comment '用户名',
    password varchar(50) not null comment '用户密码',
    sex int not null comment '用户性别',
    securityquestion varchar(100) not null comment '密保问题',
    answer varchar (100) not null comment '密保答案',
    primary key (rank)
)engine =InnoDB
 character set = utf8mb4
 auto_increment=1
 collate = utf8mb4_general_ci;

create table goods(
    gid bigint not null AUTO_INCREMENT comment '商品编号',
    goodname varchar(50) not null comment '商品名称',
    price varchar(10) not null comment '商品单价',
    introduction varchar(100) not null comment '商品简介',
    size varchar(50) not null comment '商品尺寸',
    uid varchar(30) not null comment '卖家id',
    primary key (gid)
)engine =InnoDB charset =utf8mb4 auto_increment=1 collate =utf8mb4_general_ci;

create table likes(
    lid bigint not null auto_increment comment '收藏商品编号',
    gid bigint not null comment '商品编号',
    uid varchar(30) not null comment '收藏该商品的用户',
    primary key (lid)
)engine =InnoDB charset =utf8mb4 auto_increment=1 collate =utf8mb4_general_ci;

create table shoppingcart(
    sid bigint not null auto_increment comment '购物车商品编号',
    gid bigint not null comment '商品编号',
    uid varchar(30) not null comment '将该商品添入购物车的用户id',
    price varchar(10) not null comment '商品单价',
    number int not null comment '数量',
    size varchar(50) not null comment '商品尺寸',
    primary key (sid)
)engine =InnoDB charset =utf8mb4 auto_increment=1 collate =utf8mb4_general_ci;

create table comments(
    cid bigint not null auto_increment comment '评论编号',
    gid bigint not null comment '商品编号',
    uid varchar(30) not null comment '买家id',
    star tinyint not null comment '卖家评分',
    comment varchar(100) not null comment '买家评论',
    primary key (cid)
)engine =InnoDB charset =utf8mb4 auto_increment=1 collate =utf8mb4_general_ci;

create table orders(
    oid bigint not null auto_increment comment '订单编号',
    gid bigint not null comment '商品编号',
    uid varchar(30) not null comment '买家id',
    number int not null comment '数量',
    size varchar(50) not null comment '商品尺寸',
    address varchar(100) not null comment '地址',
    phonenumber varchar(20) not null comment '电话号码',
    realname varchar(20) not null comment '真实姓名',
    primary key (oid)
)engine =InnoDB charset =utf8mb4 auto_increment=1 collate =utf8mb4_general_ci;
