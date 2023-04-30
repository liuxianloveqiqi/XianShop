CREATE TABLE `user`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `passWord`     varchar(50)     NOT NULL DEFAULT null COMMENT '用户密码，MD5加密',
    `user_Nick`     varchar(100)    NOT NULL DEFAULT NULL COMMENT '用户昵称',
    `user_Face`     varchar(255)    NOT NULL DEFAULT NULL COMMENT '用户头像地址',
    `User_Sex`      tinyint(1)      NOT NULL DEFAULT null COMMENT '用户性别：0男，1女，2保密',
    `user_Email`    varchar(255)    NOT NULL DEFAULT NULL COMMENT '用户邮箱',
    `user_Phone`    varchar(20)     NOT NULL DEFAULT NULL COMMENT '手机号',
    `github_id`     varchar(255)    NOT NULL DEFAULT NULL COMMENT 'github_id',
    `qq_id`         varchar(255)    NOT NULL DEFAULT NULL COMMENT 'qq_id',
    `login_Address` varchar(255)    NOT NULL DEFAULT NULL COMMENT '用户登录IP地址',
    `create_time`  timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `userPhone` (`user_Phone`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';
