CREATE DATABASE agriculture ;
USE agriculture;

-- 用户表
CREATE TABLE users (
                       id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
                       password VARCHAR(255) NOT NULL COMMENT '用户密码',
                       avatar_path VARCHAR(255) DEFAULT '' COMMENT '头像地址',
                       nickname VARCHAR(32) NOT NULL COMMENT '用户昵称',
                       role TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户角色(0普通用户 1管理员 2超级管理员)',
                       status VARCHAR(10) NOT NULL DEFAULT 'enabled' COMMENT '用户状态',
                       phone VARCHAR(20) DEFAULT '' COMMENT '手机号',
                       created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                       updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                       deleted_at DATETIME DEFAULT NULL COMMENT '删除时间',
                       INDEX idx_status (status),
                       INDEX idx_deleted_at (deleted_at),
                       INDEX idx_role (role)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

ALTER TABLE users MODIFY COLUMN phone VARCHAR(20) NOT NULL;
ALTER TABLE users ADD UNIQUE INDEX idx_phone_unique (phone);

-- 新闻表
CREATE TABLE news (
                      id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '新闻ID',
                      title VARCHAR(255) NOT NULL COMMENT '新闻标题',
                      abstract TEXT COMMENT '新闻摘要',
                      keyword JSON COMMENT '关键词列表',
                      source VARCHAR(100) DEFAULT '' COMMENT '新闻来源',
                      content LONGTEXT COMMENT '新闻内容',
                      status VARCHAR(20) NOT NULL DEFAULT 'draft' COMMENT '状态',
                      comment VARCHAR(255) DEFAULT '' COMMENT '审核批注',
                      files_url JSON COMMENT '新闻图片数组地址',
                      cover_url VARCHAR(512) DEFAULT '' COMMENT '封面图地址',
                      user_id BIGINT UNSIGNED COMMENT '用户ID',

                      category_id BIGINT UNSIGNED COMMENT '分类ID',
                      published_at DATETIME COMMENT '发布时间',
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                      deleted_at DATETIME DEFAULT NULL COMMENT '删除时间（软删除）',

    -- 添加外键约束
                      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
                      FOREIGN KEY (category_id) REFERENCES news_categories(id) ON DELETE SET NULL,

    -- 添加索引
                      INDEX idx_title (title),
                      INDEX idx_status (status),
                      INDEX idx_user_id (user_id),
                      INDEX idx_category_id (category_id),
                      INDEX idx_published_at (published_at),
                      INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='新闻表';

CREATE TABLE news_categories (
                                 id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '分类ID',
                                 name VARCHAR(50) NOT NULL COMMENT '分类名称',
                                 description TEXT COMMENT '分类描述',
                                 sort_order INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序权重',
                                 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

                                 INDEX idx_sort_order (sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='新闻分类表';


-- 插入用户数据 (6条记录)
INSERT INTO users (password, avatar_path, nickname, role, status, phone)
VALUES
    -- 普通用户
    ('$2a$10$Trn.3XpZ8bAe1s5aVK0zN.JkzW4dFgY7HjvC', '/images/avatar1.jpg', '农家小张', 0, 'enabled', '13800138001'),
    ('$2a$10$kLm9nT7sBvG3yHpXfR5uXe', '/images/avatar2.jpg', '果农小李', 0, 'enabled', '13900139002'),
    -- 管理员
    ('$2a$10$pQr4vJwXfHjKlMnOpqRsLe', '/images/admin_avatar.jpg', '农业管理员', 1, 'enabled', '15012345678'),
    -- 超级管理员
    ('$2a$10$sT6uW8xYzAbCdEfGhIjKlO', '/images/super_admin.jpg', '系统超管', 2, 'enabled', '15100000000'),
    -- 禁用用户
    ('$2a$10$vW9xYzAbCdEfGhIjKlOpqR', '/images/blocked_user.jpg', '违规用户', 0, 'disabled', '15211111111'),
    -- 新注册用户
    ('$2a$10$wXyZAbCdEfGhIjKlMnOpqR', '', '新农人', 0, 'enabled', '15322222222');

-- 插入新闻分类数据 (需先创建分类表)
CREATE TABLE IF NOT EXISTS categories (
                                          id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                          name VARCHAR(50) NOT NULL COMMENT '分类名称',
                                          created_at DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- 插入新闻数据 (8条记录，关联用户和分类)
INSERT INTO news (
    title, abstract, keyword, source, content, status, comment,
    files_url, cover_url, user_id, category_id, published_at
)
VALUES
    (
        '夏季蔬菜价格回落，农民收入稳定增长',
        '据农业部最新数据显示，今年夏季蔬菜供应充足，价格环比下降15%，但农民收入仍保持稳定增长...',
        '["蔬菜", "价格", "农民收入", "农业政策"]',
        '农业日报',
        '<p>在政府稳产保供政策的支持下，今年夏季蔬菜产量同比增长12%...详细分析市场供需变化...</p>',
        'published',
        '',
        '["/news/images/veg1.jpg", "/news/images/veg2.jpg", "/news/images/veg3.jpg"]',
        '/news/covers/vegetable_cover.jpg',
        1,  -- 用户ID: 农家小张
        3,  -- 分类ID: 市场行情
        '2025-06-10 09:30:00'
    ),
    (
        '智能灌溉系统助力节水农业',
        '新一代物联网智能灌溉系统在华北平原试点成功，节水率达40%以上...',
        '["智能农业", "节水灌溉", "物联网", "技术创新"]',
        '农业科技周刊',
        '<h2>技术突破</h2><p>该系统通过土壤湿度传感器实时监测...详细介绍技术原理和应用案例...</p>',
        'published',
        '技术内容审核通过',
        '["/news/images/tech1.jpg", "/news/images/tech2.jpg"]',
        '/news/covers/irrigation_cover.jpg',
        3,  -- 用户ID: 农业管理员
        4,  -- 分类ID: 农业科技
        '2025-06-05 14:20:00'
    ),
    (
        '有机肥料国家标准正式实施',
        '新版有机肥料国家标准于本月正式实施，对促进绿色农业发展具有重要意义...',
        '["有机肥", "国家标准", "绿色农业", "政策法规"]',
        '农业部官网',
        '<section>新标准明确了有机肥料的分类、技术要求、检测方法等内容...</section>',
        'approved',
        '政策文件需补充解读文章',
        '["/news/images/policy1.jpg"]',
        '/news/covers/fertilizer_cover.jpg',
        4,  -- 用户ID: 系统超管
        1,  -- 分类ID: 政策法规
        NULL  -- 尚未发布
    ),
    (
        '苹果套袋技术优化方案',
        '农科院发布最新苹果套袋技术指南，可降低病虫害发生率30%...',
        '["苹果种植", "套袋技术", "病虫害防治"]',
        '果树种植月刊',
        '详细图文介绍套袋时间选择、材料选择和操作要点...',
        'reviewing',
        '请补充实际应用案例',
        NULL,
        '',
        2,  -- 用户ID: 果农小李
        2,  -- 分类ID: 种植技术
        NULL
    ),
    (
        '大豆深加工产业链升级',
        '东北地区大豆深加工产业升级，新增10条高蛋白提取生产线...',
        '["大豆", "农产品加工", "产业链"]',
        '农产品加工报',
        '报道加工技术升级和市场需求变化...',
        'draft',
        '',
        '[]',
        '',
        1,  -- 用户ID: 农家小张
        5,  -- 分类ID: 农产品加工
        NULL
    ),
    (
        '台风预警：农业防灾指南',
        '气象局发布台风蓝色预警，农业部门紧急发布防灾减灾措施...',
        '["台风", "防灾减灾", "农业生产"]',
        '气象与农业',
        '<div class="warning">重点防护措施：1. 加固大棚 2. 疏通沟渠 3. 抢收成熟作物...</div>',
        'published',
        '',
        '["/news/images/typhoon1.jpg", "/news/images/typhoon2.jpg"]',
        '/news/covers/typhoon_cover.jpg',
        3,  -- 用户ID: 农业管理员
        1,  -- 分类ID: 政策法规
        '2025-06-15 16:45:00'
    ),
    (
        '数字农业园区建设指南发布',
        '农业农村部印发《数字农业园区建设指南》，明确三年建设目标...',
        '["数字农业", "智慧农场", "政策指导"]',
        '数字农业观察',
        '指南包含基础设施建设标准、智能装备配置要求和数据平台规范...',
        'unpublished',
        '等待配套资金文件下达',
        '["/news/images/digital1.jpg"]',
        '/news/covers/digital_agri.jpg',
        4,  -- 用户ID: 系统超管
        4,  -- 分类ID: 农业科技
        NULL
    ),
    (
        '2025年小麦收购价公布',
        '国家发改委公布2025年小麦最低收购价格，每50公斤提高3元...',
        '["小麦", "收购价", "粮食政策"]',
        '粮食经济报',
        '<table><tr><th>品种</th><th>价格(元/50kg)</th></tr><tr><td>三等小麦</td><td>118</td></tr>...</table>',
        'offline',
        '新价格文件即将发布',
        '[]',
        '',
        3,  -- 用户ID: 农业管理员
        1,  -- 分类ID: 政策法规
        '2024-09-01 00:00:00'  -- 历史发布时间
    );


INSERT INTO news_categories (id, name, description, sort_order, created_at)
VALUES
    (1, '科学技术', '关于农业科技的新闻', 1, '2023-10-10 10:10:10'),
    (2, '市场分析', '农产品市场行情分析', 2, '2023-10-11 11:11:11'),
    (3, '政策法规', '农业相关政策和法规解读', 3, '2023-10-12 12:12:12'),
    (4, '种植方法', '农作物种植技术和方法', 4, '2023-10-13 13:13:13'),
    (5, '行业动态', '农业行业最新发展动态', 5, '2023-10-14 14:14:14');

-- 添加额外的分类用于测试
INSERT INTO news_categories (name, description, sort_order)
VALUES
    ('农产品加工', '农产品深加工技术与方法', 6),
    ('有机农业', '有机种植和生态农业技术', 7),
    ('农业金融', '农业贷款和金融服务信息', 8);