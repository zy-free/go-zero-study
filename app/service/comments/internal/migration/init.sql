create table `comments` (
    `id` bigint(11) unsigned not null auto_increment comment '评论主键',
    `room_id` bigint(11) unsigned not null  default '0' comment '房间id',
    `user_id` bigint(11) unsigned not null default '0' comment  '用户id',
    `nickname` varchar(100) not null  default '' comment '用户昵称',
    `avatar` varchar(200) not null default '' comment '用户头衔',
    `content` varchar(700) not null  default '' comment '评论内存',
    `create_time` datetime not  null  default current_timestamp comment '创建时间',
    `update_time` datetime not null default current_timestamp on update  current_timestamp comment '更新时间',
    primary key (`id`),
    key `idx_room_id` (`room_id`),
    key `idx_user_id` (`user_id`)
) engine=InnoDb charset=utf8mb4 comment='评论表';
