/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : xinfei

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 10/06/2020 15:40:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for xinfei_ statistics
-- ----------------------------
DROP TABLE IF EXISTS `xinfei_ statistics`;
CREATE TABLE `xinfei_ statistics` (
  `date` date NOT NULL COMMENT '日期',
  `pv_day` int(11) DEFAULT NULL COMMENT '日浏览量',
  `pv_month` int(11) DEFAULT NULL COMMENT '月浏览量',
  `pv_year` int(11) DEFAULT NULL COMMENT '年浏览量',
  `uv_day` int(11) DEFAULT NULL COMMENT '日访问量',
  `uv_month` int(11) DEFAULT NULL COMMENT '月访问量',
  `uv_year` int(11) DEFAULT NULL COMMENT '年访问量',
  PRIMARY KEY (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for xinfei_admin
-- ----------------------------
DROP TABLE IF EXISTS `xinfei_admin`;
CREATE TABLE `xinfei_admin` (
  `admin_id` varchar(11) NOT NULL COMMENT '管理员账号ID',
  `admin_name` varchar(25) NOT NULL COMMENT '管理员名称',
  `admin_password` varchar(255) NOT NULL COMMENT '管理员密码',
  `admin_permission` int(1) NOT NULL COMMENT '管理员权限',
  `admin_icon` varchar(255) DEFAULT NULL COMMENT '管理员头像',
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of xinfei_admin
-- ----------------------------
BEGIN;
INSERT INTO `xinfei_admin` VALUES ('10001', 'admin', 'e10adc3949ba59abbe56e057f20f883e', 9, NULL);
COMMIT;

-- ----------------------------
-- Table structure for xinfei_article
-- ----------------------------
DROP TABLE IF EXISTS `xinfei_article`;
CREATE TABLE `xinfei_article` (
  `article_id` varchar(11) NOT NULL COMMENT '文章ID',
  `article_title` varchar(50) NOT NULL COMMENT '文章标题',
  `article_date` varchar(25) NOT NULL COMMENT '发布日期',
  `article_text` longtext NOT NULL COMMENT '文章内容',
  `article_hot` int(11) NOT NULL DEFAULT '1' COMMENT '热度值',
  `article_user` varchar(25) NOT NULL COMMENT '发布者',
  `article_category` varchar(10) DEFAULT NULL COMMENT '文章类别',
  PRIMARY KEY (`article_id`),
  UNIQUE KEY `article_id` (`article_id`) USING HASH COMMENT 'id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of xinfei_article
-- ----------------------------
BEGIN;
INSERT INTO `xinfei_article` VALUES ('102359', '123', '2020-06-09 08:32:51', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('104163', '123', '2020-06-09 08:33:04', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('121864', '123', '2020-06-09 08:32:51', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('131551', '123', '2020-06-02 09:18:13', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('142405', '123', '2020-06-02 09:18:07', 'fdfjsdbfjsk试试bf', 4, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('157421', '123', '2020-06-08 14:45:29', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('159669', '123', '2020-06-08 14:45:28', 'fdfjsdbfjsk试试bf', 7, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('163717', '123', '2020-06-08 14:45:31', 'fdfjsdbfjsk试试bf', 4, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('173901', '123', '2020-06-09 08:33:01', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('248324', '123', '2020-06-09 08:33:01', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('278733', '123', '2020-06-08 16:18:12', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('301456', '123', '2020-06-08 16:18:14', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('317050', '123', '2020-06-08 16:18:16', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('345630', '123', '2020-06-09 08:32:59', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('372149', '123', '2020-06-09 08:32:57', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('41035', '123', '2020-06-09 08:32:54', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('419910', '123', '2020-06-08 16:18:15', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('435682', '123', '2020-06-08 16:18:10', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('467262', '123', '2020-06-09 08:32:55', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('468853', '123', '2020-06-09 08:32:50', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('484264', '123', '2020-06-09 08:32:58', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('49240', '123', '2020-06-09 08:32:59', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('495107', '123', '2020-06-09 08:32:56', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('504802', '123', '2020-06-09 08:33:02', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('609545', '123', '2020-06-08 16:18:13', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('647959', '123', '2020-06-09 08:32:52', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('701487', '123', '2020-06-09 08:32:55', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('764142', '123', '2020-06-08 16:18:15', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('838085', '123', '2020-06-09 08:32:54', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('907559', '123', '2020-06-09 08:32:53', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('926583', '123', '2020-06-09 08:32:57', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('933881', '123', '2020-06-09 08:32:53', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
INSERT INTO `xinfei_article` VALUES ('991658', '123', '2020-06-09 08:33:00', 'fdfjsdbfjsk试试bf', 1, 'admin', '小说');
COMMIT;

-- ----------------------------
-- Table structure for xinfei_day
-- ----------------------------
DROP TABLE IF EXISTS `xinfei_day`;
CREATE TABLE `xinfei_day` (
  `date` date NOT NULL COMMENT '日期',
  `pv_day` int(11) DEFAULT '0' COMMENT '日访问量',
  `uv_day` int(11) DEFAULT '0' COMMENT '日访客量',
  PRIMARY KEY (`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of xinfei_day
-- ----------------------------
BEGIN;
INSERT INTO `xinfei_day` VALUES ('2020-01-07', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for xinfei_user
-- ----------------------------
DROP TABLE IF EXISTS `xinfei_user`;
CREATE TABLE `xinfei_user` (
  `user_account` varchar(11) NOT NULL COMMENT '登录账号',
  `user_name` varchar(25) NOT NULL COMMENT '用户名',
  `user_password` varchar(255) NOT NULL COMMENT '登录密码',
  `user_grade` varchar(4) NOT NULL COMMENT '年级',
  `user_class` varchar(25) NOT NULL COMMENT '班级',
  `user_direction` varchar(25) NOT NULL COMMENT '专业方向',
  `user_permission` int(1) DEFAULT '0' COMMENT '用户权限',
  `user_phone` varchar(11) DEFAULT NULL COMMENT '用户手机号',
  `user_qq` varchar(25) DEFAULT NULL COMMENT '用户QQ号码',
  `user_motto` varchar(255) DEFAULT NULL COMMENT '用户座右铭',
  `user_icon` varchar(255) DEFAULT NULL COMMENT '用户头像地址',
  `user_activate` int(1) DEFAULT '0' COMMENT '0/1 未激活、激活',
  `register_date` varchar(25) DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`user_account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of xinfei_user
-- ----------------------------
BEGIN;
INSERT INTO `xinfei_user` VALUES ('1', '11', 'c81e728d9d4c2f636f067f89cc14862c', '2019', '11', '11', 0, '11', '11', '111', '', 0, '');
INSERT INTO `xinfei_user` VALUES ('16091231043', '张帅威', 'c4ca4238a0b923820dcc509a6f75849b', '2016', '软件1610', '移动互联', 1, '13151289178', '919624650', '学如逆水行舟，不进则退。', '', 1, '');
INSERT INTO `xinfei_user` VALUES ('19091231001', '小明', 'c4ca4238a0b923820dcc509a6f75849b', '2019', '软件1911', '大数据', 0, '1111111111', '1111111111', '黑色的闪电', '', 1, '');
INSERT INTO `xinfei_user` VALUES ('2', '11', 'c81e728d9d4c2f636f067f89cc14862c', '2019', '2', '2', 0, '2', '2', '2', '', 0, '');
INSERT INTO `xinfei_user` VALUES ('3', '11', 'eccbc87e4b5ce2fe28308fd9f2a7baf3', '2019', '3', '3', 0, '3', '3', '3', '', 0, '');
INSERT INTO `xinfei_user` VALUES ('4', '4', 'a87ff679a2f3e71d9181a67b7542122c', '2019', '4', '4', 0, '4', '4', '4', '', 0, '');
INSERT INTO `xinfei_user` VALUES ('5', '5', 'e4da3b7fbbce2345d7772b0674a318d5', '2019', '5', '5', 0, '5', '5', '5', '', 1, '');
INSERT INTO `xinfei_user` VALUES ('6', '6', '1679091c5a880faf6fb5e6087eb1b2dc', '2015', '6', '6', 0, '6', '6', '6', '', 1, '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
