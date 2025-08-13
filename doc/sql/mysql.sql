
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `agriculture` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `agriculture`;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
                              `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                              `name` varchar(50) NOT NULL COMMENT '分类名称',
                              `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;

--
-- Table structure for table `news`
--

DROP TABLE IF EXISTS `news`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `news` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '新闻ID',
                        `title` varchar(255) NOT NULL COMMENT '新闻标题',
                        `abstract` text COMMENT '新闻摘要',
                        `keyword` json DEFAULT NULL COMMENT '关键词列表',
                        `source` varchar(100) DEFAULT '' COMMENT '新闻来源',
                        `content` longtext COMMENT '新闻内容',
                        `status` varchar(20) DEFAULT 'draft' COMMENT '新闻状态',
                        `comment` varchar(255) DEFAULT '' COMMENT '审核批注',
                        `files_url` json DEFAULT NULL COMMENT '新闻图片地址组',
                        `cover_url` varchar(512) DEFAULT '' COMMENT '封面图地址',
                        `user_id` bigint unsigned DEFAULT NULL COMMENT '用户id',
                        `category_id` bigint unsigned DEFAULT NULL COMMENT '分类id',
                        `published_at` datetime DEFAULT NULL COMMENT '发布时间',
                        `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                        `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                        `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                        `type` varchar(50) NOT NULL COMMENT '类型',
                        PRIMARY KEY (`id`),
                        KEY `idx_title` (`title`),
                        KEY `idx_status` (`status`),
                        KEY `idx_user_id` (`user_id`),
                        KEY `idx_category_id` (`category_id`),
                        KEY `idx_published_at` (`published_at`),
                        KEY `idx_deleted_at` (`deleted_at`),
                        KEY `idx_news_title` (`title`),
                        KEY `idx_news_status` (`status`),
                        KEY `idx_news_user_id` (`user_id`),
                        KEY `idx_news_category_id` (`category_id`),
                        KEY `idx_news_published_at` (`published_at`),
                        KEY `idx_news_deleted_at` (`deleted_at`),
                        CONSTRAINT `news_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
                        CONSTRAINT `news_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `news_categories` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='新闻表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news`
--

/*!40000 ALTER TABLE `news` DISABLE KEYS */;
INSERT INTO `news` VALUES (9,'夏季蔬菜价格回落，农民收入稳定增长','据农业部最新数据显示，今年夏季蔬菜供应充足，价格环比下降15%，但农民收入仍保持稳定增长...','[\"蔬菜\", \"价格\", \"农民收入\", \"农业政策\"]','农业日报','<p>在政府稳产保供政策的支持下，今年夏季蔬菜产量同比增长12%...详细分析市场供需变化...</p>','draft','','[\"/news/images/veg1.jpg\", \"/news/images/veg2.jpg\", \"/news/images/veg3.jpg\"]','/news/covers/vegetable_cover.jpg',1,3,'2025-06-10 09:30:00','2025-06-07 10:12:56','2025-06-23 11:17:27',NULL,'新闻'),(10,'智能灌溉系统助力节水农业','新一代物联网智能灌溉系统在华北平原试点成功，节水率达40%以上...','[\"智能农业\", \"节水灌溉\", \"物联网\", \"技术创新\"]','农业科技周刊','<h2>技术突破</h2><p>该系统通过土壤湿度传感器实时监测...详细介绍技术原理和应用案例...</p>','published','技术内容审核通过','[\"/news/images/tech1.jpg\", \"/news/images/tech2.jpg\"]','/news/covers/irrigation_cover.jpg',3,4,'2025-06-05 14:20:00','2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(11,'有机肥料国家标准正式实施','新版有机肥料国家标准于本月正式实施，对促进绿色农业发展具有重要意义...','[\"有机肥\", \"国家标准\", \"绿色农业\", \"政策法规\"]','农业部官网','<section>新标准明确了有机肥料的分类、技术要求、检测方法等内容...</section>','approved','政策文件需补充解读文章','[\"/news/images/policy1.jpg\"]','/news/covers/fertilizer_cover.jpg',4,1,NULL,'2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(12,'苹果套袋技术优化方案','农科院发布最新苹果套袋技术指南，可降低病虫害发生率30%...','[\"苹果种植\", \"套袋技术\", \"病虫害防治\"]','果树种植月刊','详细图文介绍套袋时间选择、材料选择和操作要点...','reviewing','请补充实际应用案例',NULL,'',2,2,NULL,'2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(13,'大豆深加工产业链升级','东北地区大豆深加工产业升级，新增10条高蛋白提取生产线...','[\"大豆\", \"农产品加工\", \"产业链\"]','农产品加工报','报道加工技术升级和市场需求变化...','draft','','[]','',1,5,NULL,'2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(14,'台风预警：农业防灾指南','气象局发布台风蓝色预警，农业部门紧急发布防灾减灾措施...','[\"台风\", \"防灾减灾\", \"农业生产\"]','气象与农业','<div class=\"warning\">重点防护措施：1. 加固大棚 2. 疏通沟渠 3. 抢收成熟作物...</div>','published','','[\"/news/images/typhoon1.jpg\", \"/news/images/typhoon2.jpg\"]','/news/covers/typhoon_cover.jpg',3,1,'2025-06-15 16:45:00','2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(15,'数字农业园区建设指南发布','农业农村部印发《数字农业园区建设指南》，明确三年建设目标...','[\"数字农业\", \"智慧农场\", \"政策指导\"]','数字农业观察','指南包含基础设施建设标准、智能装备配置要求和数据平台规范...','unpublished','等待配套资金文件下达','[\"/news/images/digital1.jpg\"]','/news/covers/digital_agri.jpg',4,4,NULL,'2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(16,'2025年小麦收购价公布','国家发改委公布2025年小麦最低收购价格，每50公斤提高3元...','[\"小麦\", \"收购价\", \"粮食政策\"]','粮食经济报','<table><tr><th>品种</th><th>价格(元/50kg)</th></tr><tr><td>三等小麦</td><td>118</td></tr>...</table>','offline','新价格文件即将发布','[]','',3,1,'2024-09-01 00:00:00','2025-06-07 10:12:56','2025-06-07 10:12:56',NULL,''),(26,'测试','haha','[\"李\", \"佳\", \"朗\"]','111','zhsddf666','draft','','[\"73e87ea7-6ed4-4300-9b5e-176ea32c65ab.png\"]','ff3c07f5-9f1f-4f95-937e-cf8deb83d364.jpg',4,2,NULL,'2025-07-17 13:27:22','2025-07-18 09:33:08','2025-07-18 09:40:01','');
/*!40000 ALTER TABLE `news` ENABLE KEYS */;

--
-- Table structure for table `news_categories`
--

DROP TABLE IF EXISTS `news_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `news_categories` (
                                   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
                                   `name` varchar(50) NOT NULL COMMENT '分类名称',
                                   `description` text COMMENT '分类描述',
                                   `sort_order` int unsigned NOT NULL DEFAULT '0' COMMENT '排序权重',
                                   `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                                   `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_news_categories_name` (`name`),
                                   KEY `idx_sort_order` (`sort_order`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='新闻分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `news_categories`
--

/*!40000 ALTER TABLE `news_categories` DISABLE KEYS */;
INSERT INTO `news_categories` VALUES (1,'科学技术','关于农业科技的新闻',1,'2023-10-10 10:10:10','2025-06-07 10:10:41'),(2,'市场分析','农产品市场行情分析',2,'2023-10-11 11:11:11','2025-06-07 10:10:41'),(3,'政策法规','农业相关政策和法规解读',3,'2023-10-12 12:12:12','2025-06-07 10:10:41'),(4,'种植方法','农作物种植技术和方法',4,'2023-10-13 13:13:13','2025-06-07 10:10:41'),(5,'行业动态','农业行业最新发展动态',5,'2023-10-14 14:14:14','2025-06-07 10:10:41'),(6,'农产品加工','农产品深加工技术与方法',6,'2025-06-07 10:10:41','2025-06-07 10:10:41'),(7,'有机农业','有机种植和生态农业技术',7,'2025-06-07 10:10:41','2025-06-07 10:10:41'),(8,'农业金融','农业贷款和金融服务信息',8,'2025-06-07 10:10:41','2025-06-07 10:10:41');
/*!40000 ALTER TABLE `news_categories` ENABLE KEYS */;

--
-- Table structure for table `supply_demand`
--

DROP TABLE IF EXISTS `supply_demand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `supply_demand` (
                                 `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '供需ID',
                                 `title` varchar(255) NOT NULL COMMENT '标题',
                                 `content` text COMMENT '内容',
                                 `tag_weigh` varchar(255) DEFAULT '0' COMMENT '标签重量',
                                 `tag_name` varchar(255) DEFAULT '' COMMENT '标签名称',
                                 `tag_price` varchar(255) DEFAULT '' COMMENT '标签价格',
                                 `cover_url` varchar(512) DEFAULT '' COMMENT '封面图地址',
                                 `files_url` json DEFAULT NULL COMMENT '附件地址列表',
                                 `likes` int DEFAULT '0' COMMENT '点赞数',
                                 `user_id` int unsigned DEFAULT NULL COMMENT '用户id',
                                 `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                 PRIMARY KEY (`id`),
                                 KEY `user_id` (`user_id`),
                                 KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='供需表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `supply_demand`
--

/*!40000 ALTER TABLE `supply_demand` DISABLE KEYS */;
INSERT INTO `supply_demand` VALUES (1,'优质苹果出售','新鲜采摘的红富士苹果，味甜多汁。','10斤装','红富士','80元','https://example.com/covers/apple.jpg','[\"https://example.com/files/apple1.jpg\", \"https://example.com/files/apple2.jpg\"]',12,1,'2025-07-19 10:43:15','2025-07-19 10:43:15',NULL),(2,'求购玉米','需要大量饲料玉米，长期合作。','50斤装','饲料玉米','120元','https://example.com/covers/corn.jpg','[\"https://example.com/files/corn1.jpg\"]',5,2,'2025-07-19 10:43:15','2025-07-19 10:43:15',NULL),(3,'土鸡蛋出售','自家散养土鸡蛋，营养丰富，口感香浓。','30枚装','土鸡蛋','65元','https://example.com/covers/eggs.jpg','[\"https://example.com/files/egg1.jpg\"]',20,3,'2025-07-19 10:43:15','2025-07-19 10:43:15',NULL),(4,'求购辣椒','收购大量红辣椒，要求干净无虫。','20斤装','红辣椒','100元','https://example.com/covers/chili.jpg','[\"https://example.com/files/chili1.jpg\"]',8,4,'2025-07-19 10:43:15','2025-07-19 10:43:15',NULL),(5,'自家花生出售','本地花生，颗粒饱满，味道香。','25斤装','本地花生','90元','https://example.com/covers/peanut.jpg','[\"https://example.com/files/peanut1.jpg\"]',15,5,'2025-07-19 10:43:15','2025-07-19 10:43:15',NULL),(6,'学目少断山置','学目少断山置','111','111','10元','4f219581-354c-47dd-b615-1b119534960d.jpg','[\"14aec626-aac1-4fc6-aea2-29fbf104f42a.jpg\", \"9a7dcfc1-5309-4ee3-b2c0-b4d997ff2f28.png\"]',0,4,'2025-07-19 14:31:01','2025-07-19 14:56:11','2025-07-19 14:56:11'),(7,'优质苹果出售','新鲜采摘的红富士苹果，味甜多汁。','10斤装','红富士','80元','https://example.com/covers/apple.jpg','[\"https://example.com/files/apple1.jpg\", \"https://example.com/files/apple2.jpg\"]',12,1,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(8,'求购玉米','需要大量饲料玉米，长期合作。','50斤装','饲料玉米','120元','https://example.com/covers/corn.jpg','[\"https://example.com/files/corn1.jpg\"]',5,2,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(9,'土鸡蛋出售','自家散养土鸡蛋，营养丰富，口感香浓。','30枚装','土鸡蛋','65元','https://example.com/covers/eggs.jpg','[\"https://example.com/files/egg1.jpg\"]',20,3,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(10,'求购辣椒','收购大量红辣椒，要求干净无虫。','20斤装','红辣椒','100元','https://example.com/covers/chili.jpg','[\"https://example.com/files/chili1.jpg\"]',8,4,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(11,'自家花生出售','本地花生，颗粒饱满，味道香。','25斤装','本地花生','90元','https://example.com/covers/peanut.jpg','[\"https://example.com/files/peanut1.jpg\"]',15,5,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(12,'优质苹果出售','新鲜采摘的红富士苹果，味甜多汁。','10斤装','红富士','80元','https://example.com/covers/apple.jpg','[\"https://example.com/files/apple1.jpg\", \"https://example.com/files/apple2.jpg\"]',12,1,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(13,'求购玉米','需要大量饲料玉米，长期合作。','50斤装','饲料玉米','120元','https://example.com/covers/corn.jpg','[\"https://example.com/files/corn1.jpg\"]',5,2,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(14,'土鸡蛋出售','自家散养土鸡蛋，营养丰富，口感香浓。','30枚装','土鸡蛋','65元','https://example.com/covers/eggs.jpg','[\"https://example.com/files/egg1.jpg\"]',20,3,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(15,'求购辣椒','收购大量红辣椒，要求干净无虫。','20斤装','红辣椒','100元','https://example.com/covers/chili.jpg','[\"https://example.com/files/chili1.jpg\"]',8,4,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(16,'自家花生出售','本地花生，颗粒饱满，味道香。','25斤装','本地花生','90元','https://example.com/covers/peanut.jpg','[\"https://example.com/files/peanut1.jpg\"]',15,5,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(17,'学目少断山置','学目少断山置','111','111','10元','82593353-ec17-4976-a4a0-1c35b2cc9e06.jpg','[\"65dab38a-8267-41f1-a6fa-0805cb9a3e75.jpg\", \"9540570e-32d3-4248-83f1-4b7846ec5513.png\"]',0,4,'2025-08-04 17:11:19','2025-08-04 17:12:01','2025-08-04 17:12:02'),(18,'学目少断山置','学目少断山置','111','111','10元','7d6d17f7-be8c-412f-84ed-176420be2d32.jpg','[\"7d6d17f7-be8c-412f-84ed-176420be2d32.jpg\"]',0,4,'2025-08-07 11:02:19','2025-08-07 11:02:19',NULL);
/*!40000 ALTER TABLE `supply_demand` ENABLE KEYS */;

--
-- Table structure for table `supply_demand_comment`
--

DROP TABLE IF EXISTS `supply_demand_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `supply_demand_comment` (
                                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论ID',
                                         `supply_demand_id` bigint NOT NULL COMMENT '供需ID',
                                         `user_id` bigint NOT NULL COMMENT '用户ID',
                                         `comment_content` text NOT NULL COMMENT '评论内容',
                                         `like_count` int DEFAULT '0' COMMENT '点赞数',
                                         `reply_id` bigint DEFAULT '0' COMMENT '父评论ID（0表示一级评论）',
                                         `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                         `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                         PRIMARY KEY (`id`),
                                         KEY `supply_demand_id` (`supply_demand_id`),
                                         KEY `reply_id` (`reply_id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='供需评论表（支持回复）';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `supply_demand_comment`
--

/*!40000 ALTER TABLE `supply_demand_comment` DISABLE KEYS */;
INSERT INTO `supply_demand_comment` VALUES (1,1,1,'苹果看起来不错，有联系方式吗？',3,0,'2025-07-18 17:52:49','2025-07-19 15:19:39',NULL),(2,1,1,'请问发货地是哪里？',1,0,'2025-07-18 17:52:49','2025-07-19 15:19:39',NULL),(3,2,2,'我这边有玉米资源，稍后联系。',2,0,'2025-07-18 17:52:49','2025-07-19 15:28:06',NULL),(4,1,3,'可以私聊我，电话xxx',0,1,'2025-07-18 17:52:49','2025-07-19 15:28:06',NULL),(5,1,202,'在山东发货。',0,2,'2025-07-18 17:52:49','2025-07-18 17:52:49',NULL),(41,1,4,'666',0,-1,'2025-07-19 15:32:08','2025-07-19 15:38:41','2025-07-19 15:38:41'),(42,1,101,'苹果看起来不错，有联系方式吗？',3,0,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(43,1,102,'请问发货地是哪里？',1,0,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(44,2,103,'我这边有玉米资源，稍后联系。',2,0,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(45,1,201,'可以私聊我，电话xxx',0,1,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(46,1,202,'在山东发货。',0,2,'2025-07-20 10:30:19','2025-07-20 10:30:19',NULL),(47,1,101,'苹果看起来不错，有联系方式吗？',3,0,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(48,1,102,'请问发货地是哪里？',1,0,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(49,2,103,'我这边有玉米资源，稍后联系。',2,0,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(50,1,201,'可以私聊我，电话xxx',0,1,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL),(51,1,202,'在山东发货。',0,2,'2025-07-20 10:31:21','2025-07-20 10:31:21',NULL);
/*!40000 ALTER TABLE `supply_demand_comment` ENABLE KEYS */;


# 给评论表的回复id加索引
ALTER TABLE `supply_demand_comment` ADD INDEX `idx_reply_id` (`reply_id`);
--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                         `password` varchar(255) NOT NULL COMMENT '用户密码',
                         `avatar_path` varchar(255) DEFAULT '' COMMENT '头像地址',
                         `nickname` varchar(32) NOT NULL COMMENT '用户昵称',
                         `role` tinyint NOT NULL DEFAULT '0' COMMENT '用户角色',
                         `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用户状态(0启用 1禁用)',
                         `phone` varchar(20) NOT NULL COMMENT '手机号',
                         `created_at` datetime DEFAULT NULL COMMENT '创建时间',
                         `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
                         `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uni_users_phone` (`phone`),
                         KEY `idx_status` (`status`),
                         KEY `idx_deleted_at` (`deleted_at`),
                         KEY `idx_role` (`role`),
                         KEY `idx_users_status` (`status`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'$2a$10$Trn.3XpZ8bAe1s5aVK0zN.JkzW4dFgY7HjvC','','农家小张',0,1,'13800138001','2025-06-06 21:46:18','2025-06-06 21:46:18',NULL),(2,'$2a$10$kLm9nT7sBvG3yHpXfR5uXe','/images/avatar2.jpg','果农小李',0,1,'13900139002','2025-06-06 21:46:18','2025-06-06 21:46:18',NULL),(3,'$2a$10$RmodSsu.v7Lu8yP7jXa9VeUUd8lOWEpXS62RlPv2bvtfIoePrQn4S','/images/admin_avatar.jpg','农业管理员',1,1,'15012345678','2025-06-06 21:46:18','2025-06-10 16:54:18',NULL),(4,'$2a$10$cHDQzXE1sy67yy2cU/Uqz.u6GpeFTWAN0ELSEMBrBrdAuOG3MpAoq','bcea8677-9f30-4ab8-8fda-845d6ac181c1.jpg','小王',2,1,'19820795808','2025-06-06 21:46:18','2025-07-23 22:56:01',NULL),(5,'$2a$10$vW9xYzAbCdEfGhIjKlOpqR','/images/blocked_user.jpg','违规用户',0,0,'15211111111','2025-06-06 21:46:18','2025-06-06 21:46:18',NULL),(6,'$2a$10$wXyZAbCdEfGhIjKlMnOpqR','','新农人',0,1,'15322222222','2025-06-06 21:46:18','2025-06-06 21:46:18',NULL),(19,'$2a$10$BrpGRXwvIxfCpzkZb.3mx.YMh1ZXYRiPKMoDzIoYq1Ik10FCfShZu','','',0,1,'19820795801','2025-06-09 17:20:11','2025-06-09 17:20:11','2025-06-10 16:40:16');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

