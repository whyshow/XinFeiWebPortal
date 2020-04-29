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

 Date: 29/04/2020 12:39:38
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
  `article_display` int(1) NOT NULL DEFAULT '1' COMMENT '是否显示',
  `article_category` varchar(10) DEFAULT NULL COMMENT '文章类别',
  PRIMARY KEY (`article_id`),
  UNIQUE KEY `article_id` (`article_id`) USING HASH COMMENT 'id索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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

SET FOREIGN_KEY_CHECKS = 1;
