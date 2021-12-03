-- 建库
CREATE DATABASE IF NOT EXISTS `userdb` default charset utf8 COLLATE utf8_general_ci;
-- 切换数据库
use `userdb`;

-- 用户表
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `user_id`  INT UNSIGNED PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
    `uid`      VARCHAR(20)  NOT NULL UNIQUE COMMENT '认证码',
    `phone`    VARCHAR(20)  NOT NULL UNIQUE COMMENT '手机号',
    `stu_id`   VARCHAR(12)  NOT NULL UNIQUE COMMENT '学号',
    `name`     VARCHAR(20)  NOT NULL COMMENT '用户名',
    `gender`   VARCHAR(5)   NOT NULL COMMENT '性别',
    `password` VARCHAR(128) NOT NULL COMMENT 'MD5密码'
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- Session表
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`
(
    `session_id` VARCHAR(100) PRIMARY KEY,
    `user_id`    INT UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 宿舍楼表
DROP TABLE IF EXISTS `buildings`;
CREATE TABLE `buildings`
(
    `building_id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(50) NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 单元表
DROP TABLE IF EXISTS `units`;
CREATE TABLE `units`
(
    `unit_id`     INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`        VARCHAR(50) NOT NULL,
    `building_id` INT UNSIGNED NOT NULL,
    FOREIGN KEY (`building_id`) REFERENCES buildings (`building_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 宿舍表
DROP TABLE IF EXISTS `dorms`;
CREATE TABLE `dorms`
(
    `dorm_id`        INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `name`           VARCHAR(20) NOT NULL,
    `gender`         CHAR(10),
    `total_beds`     INT UNSIGNED NOT NULL,
    `available_beds` INT UNSIGNED NOT NULL,
    `unit_id`        INT UNSIGNED,
    FOREIGN KEY (`unit_id`) REFERENCES units (`unit_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 订单表
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `order_id`    VARCHAR(100) PRIMARY KEY,
    `user_id`     INT UNSIGNED NOT NULL,
    `count`       INT UNSIGNED NOT NULL,
    `building_id` INT UNSIGNED NOT NULL,
    `gender`      VARCHAR(5)  NOT NULL,
    `create_time` VARCHAR(50) NOT NULL,
    `state`       INT UNSIGNED NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`),
    FOREIGN KEY (`building_id`) REFERENCES buildings (`building_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 订单项表
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`
(
    `item_id`  INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `order_id` VARCHAR(100) NOT NULL,
    `user_id`  INT UNSIGNED NOT NULL,
    FOREIGN KEY (`order_id`) REFERENCES orders (`order_id`),
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 学生选宿舍表
DROP TABLE IF EXISTS `user_dorm`;
CREATE TABLE `user_dorm`
(
    `user_id` INT UNSIGNED PRIMARY KEY,
    `dorm_id` INT UNSIGNED,
    FOREIGN KEY (`user_id`) REFERENCES users (`user_id`),
    FOREIGN KEY (`dorm_id`) REFERENCES dorms (`dorm_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- 创建视图
DROP VIEW IF EXISTS
    dorm_count;
CREATE VIEW dorm_count AS
SELECT `joined`.`building_name`  AS `building_name`,
       `joined`.`available_beds` AS `available_beds`,
       `joined`.`gender`         AS `gender`,
       COUNT(0)                  AS `count`
FROM (
         SELECT `userdb`.`buildings`.`name`       AS `building_name`,
                `userdb`.`dorms`.`name`           AS `dorm_name`,
                `userdb`.`dorms`.`available_beds` AS `available_beds`,
                `userdb`.`dorms`.`gender`         AS `gender`
         FROM ((`userdb`.`buildings` JOIN `userdb`.`units`)
                  JOIN `userdb`.`dorms` )
         WHERE ((`userdb`.`buildings`.`building_id` = `userdb`.`units`.`building_id`) AND
                (`userdb`.`units`.`unit_id` = `userdb`.`dorms`.`unit_id`))
     ) `joined`
GROUP BY `joined`.`building_name`, `joined`.`available_beds`, `joined`.`gender`
ORDER BY `joined`.`building_name`, `joined`.`available_beds`;

-- 创建视图
DROP VIEW IF EXISTS
    dorm_list;
CREATE VIEW dorm_list AS
SELECT `userdb`.`buildings`.`name`       AS `building_name`,
       `userdb`.`units`.`name`           AS `unit_name`,
       `userdb`.`dorms`.`name`           AS `dorm_name`,
       `userdb`.`dorms`.`dorm_id`        AS `dorm_id`,
       `userdb`.`dorms`.`available_beds` AS `available_beds`,
       `userdb`.`dorms`.`gender`         AS `gender`
FROM ((`userdb`.`buildings` JOIN `userdb`.`units`)
         JOIN `userdb`.`dorms` )
WHERE ((`userdb`.`buildings`.`building_id` = `userdb`.`units`.`building_id`)
    AND (`userdb`.`units`.`unit_id` = `userdb`.`dorms`.`unit_id`));

-- 初始化数据
INSERT
INTO `buildings` (`building_id`, `name`)
VALUES (1, '5号楼'),
       (2, '8号楼'),
       (3, '9号楼'),
       (4, '13号楼'),
       (5, '14号楼');

INSERT INTO `units` (`unit_id`, `name`, `building_id`)
VALUES (1, '1层', '1'),
       (2, '2层', '1'),
       (3, '3层', '1'),
       (4, '4层', '1'),
       (5, '5层', '1'),
       (6, '1层', '2'),
       (7, '2层', '2'),
       (8, '3层', '2'),
       (9, '4层', '2'),
       (10, '1层', '3'),
       (11, '2层', '3'),
       (12, '3层', '3'),
       (13, '4层', '3'),
       (14, '1单元', '4'),
       (15, '2单元', '4'),
       (16, '3单元', '4'),
       (17, '1单元', '5'),
       (18, '2单元', '5'),
       (19, '3单元', '5');

INSERT INTO `dorms` (`dorm_id`, `name`, `gender`, `total_beds`, `available_beds`, `unit_id`)
VALUES (NULL, 'e1111', '女', '2', '1', '14'),
       (NULL, 'e1112', '女', '2', '2', '14'),
       (NULL, 'e1113', '女', '3', '2', '14'),
       (NULL, 'e1114', '女', '5', '4', '14'),
       (NULL, 'e1115', '女', '3', '1', '14'),
       (NULL, 'e1121', '女', '2', '1', '14'),
       (NULL, 'e1122', '女', '2', '2', '14'),
       (NULL, 'e1123', '女', '3', '2', '14'),
       (NULL, 'e1124', '女', '5', '4', '14'),
       (NULL, 'e1125', '女', '3', '0', '14'),
       (NULL, 'e1211', '女', '2', '1', '14'),
       (NULL, 'e1212', '女', '2', '2', '14'),
       (NULL, 'e1213', '女', '3', '2', '14'),
       (NULL, 'e1214', '女', '5', '4', '14'),
       (NULL, 'e1215', '女', '3', '1', '14'),
       (NULL, 'e1221', '女', '2', '1', '14'),
       (NULL, 'e1222', '女', '2', '2', '14'),
       (NULL, 'e1223', '女', '3', '0', '14'),
       (NULL, 'e1224', '女', '5', '4', '14'),
       (NULL, 'e1225', '女', '3', '1', '14'),
       (NULL, 'f1111', '女', '2', '0', '17'),
       (NULL, 'f1112', '女', '2', '2', '17'),
       (NULL, 'f1113', '女', '3', '2', '17'),
       (NULL, 'f1114', '女', '5', '4', '17'),
       (NULL, 'f1115', '女', '3', '1', '17'),
       (NULL, 'f1121', '女', '2', '1', '17'),
       (NULL, 'f1122', '女', '2', '0', '17'),
       (NULL, 'f1123', '女', '3', '0', '17'),
       (NULL, 'f1124', '女', '5', '4', '17'),
       (NULL, 'f1125', '女', '3', '1', '17'),
       (NULL, 'f1211', '女', '2', '1', '17'),
       (NULL, 'f1212', '女', '2', '2', '17'),
       (NULL, 'f1213', '女', '3', '2', '17'),
       (NULL, 'f1214', '女', '5', '4', '17'),
       (NULL, 'f1215', '女', '3', '1', '17'),
       (NULL, 'f1221', '女', '2', '0', '17'),
       (NULL, 'f1222', '女', '2', '0', '17'),
       (NULL, 'f1223', '女', '3', '0', '17'),
       (NULL, 'f1224', '女', '5', '4', '17'),
       (NULL, 'f1225', '女', '3', '1', '17'),
       (NULL, 'f2111', '女', '2', '1', '18'),
       (NULL, 'f2112', '女', '2', '2', '18'),
       (NULL, 'f2113', '女', '3', '2', '18'),
       (NULL, 'f2114', '女', '5', '4', '18'),
       (NULL, 'f2115', '女', '3', '1', '18'),
       (NULL, 'f2121', '女', '2', '0', '18'),
       (NULL, 'f2122', '女', '2', '2', '18'),
       (NULL, 'f2123', '女', '3', '2', '18'),
       (NULL, 'f2124', '女', '5', '4', '18'),
       (NULL, 'f2125', '女', '3', '0', '18'),
       (NULL, 'f2211', '女', '2', '1', '18'),
       (NULL, 'f2212', '女', '2', '2', '18'),
       (NULL, 'f2213', '女', '3', '2', '18'),
       (NULL, 'f2214', '女', '5', '4', '18'),
       (NULL, 'f2215', '女', '3', '1', '18'),
       (NULL, 'f2221', '女', '2', '0', '18'),
       (NULL, 'f2222', '女', '2', '2', '18'),
       (NULL, 'f2223', '女', '3', '2', '18'),
       (NULL, 'f2224', '女', '5', '4', '18'),
       (NULL, 'f2225', '女', '3', '1', '18'),
       (NULL, '5201', '女', '7', '5', '2'),
       (NULL, '5202', '女', '4', '3', '2'),
       (NULL, '5203', '女', '4', '2', '2'),
       (NULL, '5204', '女', '4', '3', '2'),
       (NULL, '5205', '女', '4', '1', '2'),
       (NULL, '5206', '女', '4', '4', '2'),
       (NULL, '5207', '女', '4', '0', '2'),
       (NULL, '5208', '女', '4', '2', '2'),
       (NULL, '5209', '女', '4', '3', '2'),
       (NULL, '5210', '女', '4', '3', '2'),
       (NULL, '5211', '女', '7', '5', '2'),
       (NULL, '5212', '女', '4', '3', '2'),
       (NULL, '5213', '女', '4', '2', '2'),
       (NULL, '5214', '女', '4', '3', '2'),
       (NULL, '5215', '女', '4', '1', '2'),
       (NULL, '5216', '女', '4', '4', '2'),
       (NULL, '5217', '女', '4', '0', '2'),
       (NULL, '5218', '女', '4', '2', '2'),
       (NULL, '5219', '女', '4', '3', '2'),
       (NULL, '5220', '女', '7', '0', '2'),
       (NULL, '5501', '女', '7', '5', '5'),
       (NULL, '5502', '女', '4', '3', '5'),
       (NULL, '5503', '女', '4', '2', '5'),
       (NULL, '5504', '女', '4', '3', '5'),
       (NULL, '5505', '女', '4', '1', '5'),
       (NULL, '5506', '女', '4', '0', '5'),
       (NULL, '5507', '女', '4', '2', '5'),
       (NULL, '5508', '女', '4', '2', '5'),
       (NULL, '5509', '女', '4', '3', '5'),
       (NULL, '5510', '女', '4', '3', '5'),
       (NULL, '5511', '女', '4', '2', '5'),
       (NULL, '5512', '女', '4', '3', '5'),
       (NULL, '5513', '女', '4', '2', '5'),
       (NULL, '5514', '女', '4', '3', '5'),
       (NULL, '5515', '女', '4', '1', '5'),
       (NULL, '5516', '女', '4', '0', '5'),
       (NULL, '5517', '女', '4', '2', '5'),
       (NULL, '5518', '女', '4', '2', '5'),
       (NULL, '5519', '女', '4', '3', '5'),
       (NULL, '5520', '女', '7', '4', '5'),
       (NULL, '5101', '男', '7', '5', '1'),
       (NULL, '5102', '男', '4', '3', '1'),
       (NULL, '5103', '男', '4', '2', '1'),
       (NULL, '5104', '男', '4', '3', '1'),
       (NULL, '5105', '男', '4', '0', '1'),
       (NULL, '5106', '男', '4', '4', '1'),
       (NULL, '5107', '男', '4', '2', '1'),
       (NULL, '5108', '男', '4', '2', '1'),
       (NULL, '5109', '男', '4', '3', '1'),
       (NULL, '5110', '男', '4', '3', '1'),
       (NULL, '5111', '男', '4', '2', '1'),
       (NULL, '5112', '男', '4', '3', '1'),
       (NULL, '5113', '男', '4', '2', '1'),
       (NULL, '5114', '男', '4', '0', '1'),
       (NULL, '5115', '男', '4', '1', '1'),
       (NULL, '5116', '男', '4', '4', '1'),
       (NULL, '5117', '男', '4', '2', '1'),
       (NULL, '5118', '男', '4', '2', '1'),
       (NULL, '5119', '男', '4', '3', '1'),
       (NULL, '5120', '男', '7', '4', '1'),
       (NULL, '5301', '男', '7', '5', '3'),
       (NULL, '5302', '男', '4', '3', '3'),
       (NULL, '5303', '男', '4', '2', '3'),
       (NULL, '5304', '男', '4', '3', '3'),
       (NULL, '5305', '男', '4', '0', '3'),
       (NULL, '5306', '男', '4', '4', '3'),
       (NULL, '5307', '男', '4', '2', '3'),
       (NULL, '5308', '男', '4', '2', '3'),
       (NULL, '5309', '男', '4', '3', '3'),
       (NULL, '5310', '男', '4', '3', '3'),
       (NULL, '5311', '男', '4', '2', '3'),
       (NULL, '5312', '男', '4', '3', '3'),
       (NULL, '5313', '男', '4', '2', '3'),
       (NULL, '5314', '男', '4', '0', '3'),
       (NULL, '5315', '男', '4', '1', '3'),
       (NULL, '5316', '男', '4', '4', '3'),
       (NULL, '5317', '男', '4', '2', '3'),
       (NULL, '5318', '男', '4', '2', '3'),
       (NULL, '5319', '男', '4', '3', '3'),
       (NULL, '5320', '男', '7', '4', '3'),
       (NULL, '5401', '男', '7', '5', '4'),
       (NULL, '5402', '男', '4', '3', '4'),
       (NULL, '5403', '男', '4', '2', '4'),
       (NULL, '5404', '男', '4', '3', '4'),
       (NULL, '5405', '男', '4', '0', '4'),
       (NULL, '5406', '男', '4', '4', '4'),
       (NULL, '5407', '男', '4', '2', '4'),
       (NULL, '5408', '男', '4', '2', '4'),
       (NULL, '5409', '男', '4', '3', '4'),
       (NULL, '5410', '男', '4', '3', '4'),
       (NULL, '5411', '男', '4', '2', '4'),
       (NULL, '5412', '男', '4', '3', '4'),
       (NULL, '5413', '男', '4', '2', '4'),
       (NULL, '5414', '男', '4', '0', '4'),
       (NULL, '5415', '男', '4', '1', '4'),
       (NULL, '5416', '男', '4', '4', '4'),
       (NULL, '5417', '男', '4', '2', '4'),
       (NULL, '5418', '男', '4', '2', '4'),
       (NULL, '5419', '男', '4', '3', '4'),
       (NULL, '5420', '男', '7', '4', '4'),
       (NULL, '8111', '男', '4', '3', '6'),
       (NULL, '8112', '男', '4', '1', '6'),
       (NULL, '8113', '男', '4', '2', '6'),
       (NULL, '8114', '男', '6', '2', '6'),
       (NULL, '8121', '男', '4', '3', '6'),
       (NULL, '8122', '男', '4', '1', '6'),
       (NULL, '8123', '男', '4', '2', '6'),
       (NULL, '8124', '男', '6', '2', '6'),
       (NULL, '8211', '男', '4', '3', '7'),
       (NULL, '8212', '男', '4', '1', '7'),
       (NULL, '8213', '男', '4', '3', '7'),
       (NULL, '8214', '男', '6', '5', '7'),
       (NULL, '8221', '男', '4', '3', '7'),
       (NULL, '8222', '男', '4', '1', '7'),
       (NULL, '8223', '男', '4', '2', '7'),
       (NULL, '8224', '男', '6', '2', '7'),
       (NULL, '8311', '男', '4', '4', '8'),
       (NULL, '8312', '男', '4', '1', '8'),
       (NULL, '8313', '男', '4', '0', '8'),
       (NULL, '8314', '男', '6', '2', '8'),
       (NULL, '8321', '男', '4', '3', '8'),
       (NULL, '8322', '男', '4', '1', '8'),
       (NULL, '8323', '男', '4', '2', '8'),
       (NULL, '8324', '男', '6', '1', '8'),
       (NULL, '8411', '男', '4', '3', '9'),
       (NULL, '8412', '男', '4', '1', '9'),
       (NULL, '8413', '男', '4', '4', '9'),
       (NULL, '8414', '男', '6', '2', '9'),
       (NULL, '8421', '男', '4', '3', '9'),
       (NULL, '8422', '男', '4', '1', '9'),
       (NULL, '8423', '男', '4', '2', '9'),
       (NULL, '8424', '男', '6', '4', '9'),
       (NULL, '9111', '男', '4', '3', '10'),
       (NULL, '9112', '男', '4', '1', '10'),
       (NULL, '9113', '男', '4', '2', '10'),
       (NULL, '9114', '男', '6', '2', '10'),
       (NULL, '9121', '男', '4', '3', '10'),
       (NULL, '9122', '男', '4', '1', '10'),
       (NULL, '9123', '男', '4', '2', '10'),
       (NULL, '9124', '男', '6', '2', '10'),
       (NULL, '9211', '男', '4', '3', '11'),
       (NULL, '9212', '男', '4', '1', '11'),
       (NULL, '9213', '男', '4', '3', '11'),
       (NULL, '9214', '男', '6', '5', '11'),
       (NULL, '9221', '男', '4', '3', '11'),
       (NULL, '9222', '男', '4', '1', '11'),
       (NULL, '9223', '男', '4', '2', '11'),
       (NULL, '9224', '男', '6', '2', '11'),
       (NULL, '9311', '男', '4', '4', '12'),
       (NULL, '9312', '男', '4', '1', '12'),
       (NULL, '9313', '男', '4', '0', '12'),
       (NULL, '9314', '男', '6', '2', '12'),
       (NULL, '9321', '男', '4', '3', '12'),
       (NULL, '9322', '男', '4', '1', '12'),
       (NULL, '9323', '男', '4', '2', '12'),
       (NULL, '9324', '男', '6', '1', '12'),
       (NULL, '9411', '男', '4', '3', '13'),
       (NULL, '9412', '男', '4', '1', '13'),
       (NULL, '9413', '男', '4', '4', '13'),
       (NULL, '9414', '男', '6', '2', '13'),
       (NULL, '9421', '男', '4', '3', '13'),
       (NULL, '9422', '男', '4', '1', '13'),
       (NULL, '9423', '男', '4', '2', '13'),
       (NULL, '9424', '男', '6', '4', '13'),
       (NULL, 'e2111', '男',  '2', '1', '15'),
       (NULL, 'e2112', '男',  '2', '2', '15'),
       (NULL, 'e2113', '男',  '3', '2', '15'),
       (NULL, 'e2114', '男',  '5', '4', '15'),
       (NULL, 'e2115', '男',  '3', '1', '15'),
       (NULL, 'e2121', '男',  '2', '1', '15'),
       (NULL, 'e2122', '男',  '2', '2', '15'),
       (NULL, 'e2123', '男',  '3', '2', '15'),
       (NULL, 'e2124', '男',  '5', '4', '15'),
       (NULL, 'e2125', '男',  '3', '0', '15'),
       (NULL, 'e2211', '男',  '2', '1', '15'),
       (NULL, 'e2212', '男',  '2', '2', '15'),
       (NULL, 'e2213', '男',  '3', '2', '15'),
       (NULL, 'e2214', '男',  '5', '4', '15'),
       (NULL, 'e2215', '男',  '3', '1', '15'),
       (NULL, 'e2221', '男',  '2', '1', '15'),
       (NULL, 'e2222', '男',  '2', '2', '15'),
       (NULL, 'e2223', '男',  '3', '0', '15'),
       (NULL, 'e2224', '男',  '5', '4', '15'),
       (NULL, 'e2225', '男',  '3', '1', '15'),
       (NULL, 'e3111', '男',  '2', '1', '16'),
       (NULL, 'e3112', '男',  '2', '2', '16'),
       (NULL, 'e3113', '男',  '3', '2', '16'),
       (NULL, 'e3114', '男',  '5', '4', '16'),
       (NULL, 'e3115', '男',  '3', '1', '16'),
       (NULL, 'e3121', '男',  '2', '1', '16'),
       (NULL, 'e3122', '男',  '2', '2', '16'),
       (NULL, 'e3123', '男',  '3', '2', '16'),
       (NULL, 'e3124', '男',  '5', '4', '16'),
       (NULL, 'e3125', '男',  '3', '0', '16'),
       (NULL, 'e3211', '男',  '2', '1', '16'),
       (NULL, 'e3212', '男',  '2', '2', '16'),
       (NULL, 'e3213', '男',  '3', '2', '16'),
       (NULL, 'e3214', '男',  '5', '4', '16'),
       (NULL, 'e3215', '男',  '3', '1', '16'),
       (NULL, 'e3221', '男',  '2', '1', '16'),
       (NULL, 'e3222', '男',  '2', '2', '16'),
       (NULL, 'e3223', '男',  '3', '0', '16'),
       (NULL, 'e3224', '男',  '5', '4', '16'),
       (NULL, 'e3225', '男',  '3', '1', '16'),
       (NULL, 'f3113', '男', '3', '2', '19'),
       (NULL, 'f3114', '男', '5', '4', '19'),
       (NULL, 'f3115', '男', '3', '1', '19'),
       (NULL, 'f3121', '男', '2', '0', '19'),
       (NULL, 'f3122', '男', '2', '2', '19'),
       (NULL, 'f3123', '男', '3', '2', '19'),
       (NULL, 'f3124', '男', '5', '4', '19'),
       (NULL, 'f3125', '男', '3', '0', '19'),
       (NULL, 'f3211', '男', '2', '1', '19'),
       (NULL, 'f3212', '男', '2', '2', '19'),
       (NULL, 'f3213', '男', '3', '2', '19'),
       (NULL, 'f3214', '男', '5', '4', '19'),
       (NULL, 'f3215', '男', '3', '1', '19'),
       (NULL, 'f3221', '男', '2', '0', '19'),
       (NULL, 'f3222', '男', '2', '2', '19'),
       (NULL, 'f3223', '男', '3', '2', '19'),
       (NULL, 'f3224', '男', '5', '4', '19'),
       (NULL, 'f3225', '男', '3', '1', '19');

INSERT INTO userdb.users (user_id, uid, phone, stu_id, name, gender, password)
VALUES (1, '12345678', '18312345678', '1234567890', '张三', '女', 'e10adc3949ba59abbe56e057f20f883e');