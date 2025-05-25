
create database gd_im;
use gd_im;
/**  用户表 **/
create table `user`(
    `id` bigint(20) not null auto_increment,
    `user_name` varchar(20) not null comment '用户名',
    `password` varchar(20) not null comment '密码',
    `email` varchar(50) not null default '' comment '邮箱',
    `phone` varchar(20) not null default '' comment '手机号',
    `sex` int(10) not null default 1 comment '性别',
    `age` int(10) not null default 18 comment '年龄',
    `address` varchar(225) not null comment '地址',
    `create_time` datetime not null comment '创建时间',
    `update_time` datetime not null comment '更新时间',
    `delete_time` datetime not null comment '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_user_name` (`user_name`) USING BTREE,
    KEY `idx_email` (`email`) USING BTREE,
    KEY `idx_phone` (`phone`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';
/** 群组-用户表 **/
create table `relation`(
    `id` bigint(20) not null auto_increment,
    `group_id` bigint(20) not null comment'群组id',
    `user_id`  bigint(20) not null comment'用户id',
    `type` tinyint(1) not null default 2 comment'成员类型:0:群主,1:管理员,2:普通成员,3:禁言成员,4:拉黑成员,10:表示好友关系',
    `create_time` datetime not null comment '创建时间',
    `update_time` datetime not null comment '更新时间',
    `delete_time` datetime not null comment '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_group_id` (`group_id`) USING BTREE,
    KEY `idx_user_id` (`user_id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群组用户';
/** 群组表 **/
create table `group`(
    `id` bigint(20) not null auto_increment,
    `group_name` varchar(64) not null comment '群组名',
    `remember_num` int not null comment '成员数量',
    `owner_id ` varchar(225) not null comment '群主id',
    `type` tinyint(1) not null comment '群类型：0:私有群,1:公开群,2:临时群,3:临时会话',
    `image` varchar(225) not null comment '头像url',
    `describe` varchar(225) not null comment '描述',
    `create_user_id` bigint(20) not null comment '创建者id',
    `update_user_id` bigint(20) not null comment '更新者id',
    `delete_user_id` bigint(20) not null comment '删除者id',
    `create_time` datetime not null comment '创建时间',
    `update_time` datetime not null comment '更新时间',
    `delete_time` datetime not null comment '删除时间',
    PRIMARY KEY (`id`) USING BTREE
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='群组';
/** 配置表 **/
create table `config`(
    `id` bigint(20) not null auto_increment,
    `config_name` varchar(64) not null default'' comment '配置名',
    `config_code`  varchar(64) not null default'' comment '配置编码',
    `config_value` varchar(64) not null default'' comment '配置值',
    `describe` varchar(225) not null comment '描述',
    `create_time` datetime not null comment '创建时间',
    `update_time` datetime not null comment '更新时间',
    `delete_time` datetime not null comment '删除时间',
    PRIMARY KEY (`id`) USING BTREE
 )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='配置';