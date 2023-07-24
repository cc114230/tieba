CREATE TABLE `user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id`     bigint(20) NOT NULL,
    `username`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password`    varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email`       varchar(64) COLLATE utf8mb4_general_ci,
    `gender`      tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


create table community
(
    id             int auto_increment
        primary key,
    community_id   int unsigned not null,
    community_name varchar(128)                        not null,
    introduction   varchar(256)                        not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    update_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint idx_community_id
        unique (community_id),
    constraint idx_community_name
        unique (community_name)
) collate = utf8mb4_general_ci;

INSERT INTO tieba.community (id, community_id, community_name, introduction, create_time, update_time)
VALUES (1, 1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO tieba.community (id, community_id, community_name, introduction, create_time, update_time)
VALUES (2, 2, 'C', 'C/C++', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO tieba.community (id, community_id, community_name, introduction, create_time, update_time)
VALUES (3, 3, 'CS:GO', 'Rush B。。。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO tieba.community (id, community_id, community_name, introduction, create_time, update_time)
VALUES (4, 4, 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');



create table post
(
    id           bigint auto_increment
        primary key,
    post_id      bigint              not null comment '帖子id',
    title        varchar(128)        not null comment '标题',
    content      varchar(8192)       not null comment '内容',
    author_id    bigint              not null comment '作者的用户id',
    community_id bigint              not null comment '所属社区',
    status       tinyint   default 1 not null comment '帖子状态',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_post_id
        unique (post_id)
) collate = utf8mb4_general_ci;

create index idx_author_id
    on post (author_id);

create index idx_community_id
    on post (community_id);



create table comment
(
    id           bigint auto_increment
        primary key,
    comment_id   bigint        not null comment '评论id',
    post_id      bigint        not null comment '评论的帖子id',
    commenter_id bigint        not null comment '评论者id',
    commenter    varchar(64)   not null comment '评论者',
    content      varchar(8192) not null comment '评论内容',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '评论时间',
    foreign key (post_id) references post (post_id)
)  collate = utf8mb4_general_ci;