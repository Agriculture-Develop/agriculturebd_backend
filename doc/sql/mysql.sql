CREATE DATABASE /*!32312 IF NOT EXISTS*/ `agriculture` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `agriculture`;

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `news`;
DROP TABLE IF EXISTS `supply_demand_comment`;
DROP TABLE IF EXISTS `supply_demand`;
DROP TABLE IF EXISTS `news_categories`;
DROP TABLE IF EXISTS `users`;
-- 先建 users 表，后续外键引用
DROP TABLE IF EXISTS `users`; -- 再次确保
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `avatar_path` varchar(255) DEFAULT '' COMMENT '头像地址',
  `nickname` varchar(32) NOT NULL COMMENT '用户昵称',
  `role` tinyint NOT NULL DEFAULT '0' COMMENT '用户角色',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用户状态(0启用 1禁用)',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
INSERT INTO `users` (`id`,`password`,`avatar_path`,`nickname`,`role`,`status`,`phone`,`created_at`) VALUES
(1,'$2a$10$demo','','农家小张',0,1,'13800138001',NOW()),
(2,'$2a$10$demo','/images/avatar2.jpg','果农小李',0,1,'13900139002',NOW()),
(3,'$2a$10$demo','/images/admin.jpg','农业管理员',1,1,'15012345678',NOW());
-- 再建分类表
DROP TABLE IF EXISTS `news_categories`; -- 用户要求显式删除
CREATE TABLE `news_categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(50) NOT NULL COMMENT '分类名称',
  `slug` varchar(50) NOT NULL COMMENT '唯一标识',
  `description` text COMMENT '分类描述',
  `sort_order` int unsigned NOT NULL DEFAULT 0 COMMENT '排序权重',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_news_categories_name` (`name`),
  UNIQUE KEY `uk_news_categories_slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='新闻分类表';
INSERT INTO `news_categories` (`id`,`name`,`slug`,`description`,`sort_order`,`created_at`) VALUES
(1,'政策法规','policy','农业政策法规',10,NOW()),
(2,'科技创新','tech','农业科技与设备',20,NOW()),
(3,'市场行情','market','价格与市场分析',30,NOW());
-- 创建 news 表（外键已存在）
DROP TABLE IF EXISTS `news`; -- 用户要求显式删除
CREATE TABLE `news` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `abstract` text,
  `keyword` json DEFAULT NULL,
  `source` varchar(100) NOT NULL DEFAULT '',
  `content` longtext NOT NULL,
  `status` enum('draft','reviewing','approved','published','offline','unpublished') NOT NULL DEFAULT 'draft',
  `comment` varchar(255) NOT NULL DEFAULT '',
  `files_url` json DEFAULT NULL,
  `cover_url` varchar(512) NOT NULL DEFAULT '',
  `type` enum('news','policy','tech','market') NOT NULL DEFAULT 'news',
  `user_id` bigint unsigned DEFAULT NULL,
  `category_id` bigint unsigned DEFAULT NULL,
  `published_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_news_status_deleted_at` (`status`,`deleted_at`),
  CONSTRAINT `fk_news_category` FOREIGN KEY (`category_id`) REFERENCES `news_categories` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_news_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='新闻表';
INSERT INTO `news` (`title`,`abstract`,`keyword`,`source`,`content`,`status`,`files_url`,`cover_url`,`type`,`user_id`,`category_id`,`published_at`) VALUES
('夏季蔬菜价格回落，农民收入稳定增长','蔬菜供应充足','["蔬菜","价格"]','农业日报','<p>内容...</p>','published','["/news/images/veg1.jpg"]','/news/covers/vegetable_cover.jpg','market',1,3,NOW());
-- 创建 supply_demand 表（依赖 users）
DROP TABLE IF EXISTS `supply_demand`; -- 用户要求显式删除
CREATE TABLE `supply_demand` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `content` text,
  `category` varchar(50) NOT NULL DEFAULT '',
  `tag_weigh` varchar(255) DEFAULT '0',
  `tag_name` varchar(255) DEFAULT '',
  `tag_price` varchar(255) DEFAULT '',
  `cover_url` varchar(512) DEFAULT '',
  `files_url` json DEFAULT NULL,
  `likes` int DEFAULT 0,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_sd_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='供需表';
INSERT INTO `supply_demand` (`title`,`content`,`category`,`tag_weigh`,`tag_name`,`tag_price`,`user_id`) VALUES
('优质苹果出售','果园现采红富士苹果','水果','10斤装','红富士','80元',1);
-- 恢复外键检查
SET FOREIGN_KEY_CHECKS=1;
