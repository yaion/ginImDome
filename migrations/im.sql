
create database im;
    use im;
        create table user(
            id bigint(20) not null auto_increment,
            user_name varchar(20) not null comment '用户名',
            password varchar(20) not null comment '密码',
            email varchar(50) not null comment '邮箱',
            phone varchar(20) not null comment '手机号',
            sex int not null comment '性别',
            age int not null comment '年龄',
            address varchar(100) not null comment '地址',
            create_time datetime not null comment '创建时间',
            update_time datetime not null comment '更新时间',
            delete_time datetime not null comment '删除时间',
            primary key (id)
            )engine=innodb default charset=utf8mb4 comment='用户';
        create group_user(
            group_id bigint(20) not null comment'群组id',
            user_id  bigint(20) not null comment'用户id',
            type tinyint(1) not null comment'成员类型',
            primary key(group_id,user_id)
        )engine=innodb default charset=utf8mb4 comment='群组用户';
        create table group(
            id bigint(20) not null auto_increment,
            group_name varchar(20) not null comment '群组名',
            create_time datetime not null comment '创建时间',
            remember_num int not null comment '成员数量',
            describe varchar(100) not null comment '描述',
            create_user_id bigint(20) not null comment '创建者id',
            update_user_id bigint(20) not null comment '更新者id',
            delete_user_id bigint(20) not null comment '删除者id',
            update_time datetime not null comment '更新时间',
            delete_time datetime not null comment '删除时间',
            primary key (id)
            )engine=innodb default charset=utf8mb4 comment='群组';
        create table config(
            id bigint(20) not null auto_increment,
            config_name varchar(20) not null comment '配置名',
            config_value varchar(20) not null comment '配置值',
            describe varchar(100) not null comment '描述',
            create_time datetime not null comment '创建时间',
            update_time datetime not null comment '更新时间',
            delete_time datetime not null comment '删除时间',
            primary key (id)
            )engine=innodb default charset=utf8mb4 comment='配置';