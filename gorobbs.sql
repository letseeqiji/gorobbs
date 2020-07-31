/*
Navicat MySQL Data Transfer  1.0.4

Source Server         : myaliyun-xg
Source Server Version : 80015
Source Host           : 47.91.212.4:3306
Source Database       : gorobbs

Target Server Type    : MYSQL
Target Server Version : 80015
File Encoding         : 65001
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for bbs_attach
-- ----------------------------
DROP TABLE IF EXISTS `bbs_attach`;
CREATE TABLE `bbs_attach` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件id',
  `thread_id` int(11) NOT NULL DEFAULT '0' COMMENT '主题id',
  `post_id` int(11) NOT NULL DEFAULT '0' COMMENT '帖子id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `filesize` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '文件尺寸，单位字节',
  `width` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT 'width > 0 则为图片',
  `height` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `filename` char(120) NOT NULL DEFAULT '' COMMENT '文件名称，会过滤，并且截断，保存后的文件名，不包含URL前缀 upload_url',
  `orgfilename` char(120) NOT NULL DEFAULT '' COMMENT '上传的原文件名',
  `filetype` char(7) NOT NULL DEFAULT '' COMMENT 'image/txt/zip，小图标显示',
  `comment` char(100) NOT NULL DEFAULT '' COMMENT '文件注释 方便于搜索',
  `downloads_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '下载次数',
  `credits_num` int(11) NOT NULL DEFAULT '0' COMMENT '需要的积分',
  `golds_num` int(11) NOT NULL DEFAULT '0' COMMENT '需要的金币',
  `rmbs_num` int(11) NOT NULL DEFAULT '0' COMMENT '需要的人民币',
  `isimage` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为图片',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `downloads_num` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `pid` (`post_id`),
  KEY `uid` (`user_id`),
  KEY `idx_bbs_attach_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛附件表  只能按照从上往下的方式查找和删除！ 此表如果大，可以考虑通过 aid 分区。';


-- ----------------------------
-- Table structure for bbs_cache
-- ----------------------------
DROP TABLE IF EXISTS `bbs_cache`;
CREATE TABLE `bbs_cache` (
  `k` char(32) NOT NULL DEFAULT '',
  `v` mediumtext NOT NULL,
  `expiry` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='缓存表，用来保存临时数据';

-- ----------------------------
-- Records of bbs_cache
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_forum
-- ----------------------------
DROP TABLE IF EXISTS `bbs_forum`;
CREATE TABLE `bbs_forum` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(16) NOT NULL DEFAULT '',
  `rank` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '显示，倒序，数字越大越靠前',
  `threads_cnt` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '主题数',
  `todayposts_cnt` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '今日发帖，计划任务每日凌晨０点清空为０',
  `todaythreads_cnt` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '今日发主题，计划任务每日凌晨０点清空为０',
  `brief` text NOT NULL COMMENT '版块简介 允许HTML',
  `announcement` text NOT NULL COMMENT '版块公告 允许HTML',
  `accesson` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '是否开启权限控制',
  `orderby` tinyint(11) NOT NULL DEFAULT '0' COMMENT '默认列表排序，0: 顶贴时间 last_date， 1: 发帖时间 tid',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0',
  `moduids` char(120) NOT NULL DEFAULT '' COMMENT '每个版块有多个版主，最多10个： 10*12 = 120，删除用户的时候，如果是版主，则调整后再删除。逗号分隔',
  `seo_title` char(64) NOT NULL DEFAULT '' COMMENT 'SEO 标题，如果设置会代替版块名称',
  `seo_keywords` char(64) NOT NULL DEFAULT '',
  `digests_cnt` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `digests_num` int(11) DEFAULT '0',
  `icon` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_bbs_forum_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='板块表';

-- ----------------------------
-- Records of bbs_forum
-- ----------------------------
INSERT INTO `bbs_forum` VALUES ('1', '默认版块', '0', '0', '0', '0', '默认版块介绍', '默认版块公告', '0', '0', '0', '', '', '', null, '2019-07-29 10:30:52', '2019-08-14 03:49:36', null, '0', '/static/img/forum.png');

-- ----------------------------
-- Table structure for bbs_forum_access
-- ----------------------------
DROP TABLE IF EXISTS `bbs_forum_access`;
CREATE TABLE `bbs_forum_access` (
  `id` int(11) unsigned NOT NULL DEFAULT '0',
  `group_id` int(11) unsigned NOT NULL DEFAULT '0',
  `allowread` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `allowthread` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `allowpost` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '允许回复',
  `allowattach` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '允许上传附件',
  `allowdown` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '允许下载附件',
  PRIMARY KEY (`id`,`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='版块访问规则, forum.accesson 开启时生效';

-- ----------------------------
-- Records of bbs_forum_access
-- ----------------------------
INSERT INTO `bbs_forum_access` VALUES ('2', '0', '1', '0', '1', '0', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '1', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '2', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '4', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '5', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '6', '1', '0', '1', '0', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '7', '0', '0', '0', '0', '0');
INSERT INTO `bbs_forum_access` VALUES ('2', '101', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '102', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '103', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '104', '1', '1', '1', '1', '1');
INSERT INTO `bbs_forum_access` VALUES ('2', '105', '1', '1', '1', '1', '1');

-- ----------------------------
-- Table structure for bbs_group
-- ----------------------------
DROP TABLE IF EXISTS `bbs_group`;
CREATE TABLE `bbs_group` (
  `id` smallint(6) unsigned NOT NULL,
  `name` char(20) NOT NULL DEFAULT '' COMMENT '用户组名称',
  `creditsfrom` int(11) NOT NULL DEFAULT '0' COMMENT '积分从',
  `creditsto` int(11) NOT NULL DEFAULT '0' COMMENT '积分到',
  `allowread` int(11) NOT NULL DEFAULT '0' COMMENT '允许访问',
  `allowthread` int(11) NOT NULL DEFAULT '0' COMMENT '允许发主题',
  `allowpost` int(11) NOT NULL DEFAULT '0' COMMENT '允许回帖',
  `allowattach` int(11) NOT NULL DEFAULT '0' COMMENT '允许上传文件',
  `allowdown` int(11) NOT NULL DEFAULT '0' COMMENT '允许下载文件',
  `allowtop` int(11) NOT NULL DEFAULT '0' COMMENT '允许置顶',
  `allowupdate` int(11) NOT NULL DEFAULT '0' COMMENT '允许编辑',
  `allowdelete` int(11) NOT NULL DEFAULT '0',
  `allowmove` int(11) NOT NULL DEFAULT '0',
  `allowbanuser` int(11) NOT NULL DEFAULT '0' COMMENT '允许禁止用户',
  `allowdeleteuser` int(11) NOT NULL DEFAULT '0',
  `allowviewip` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '允许查看用户敏感信息',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_bbs_group_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户组';

-- ----------------------------
-- Records of bbs_group
-- ----------------------------
INSERT INTO `bbs_group` VALUES ('0', '游客组', '0', '0', '1', '0', '1', '0', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('1', '管理员组', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', null, null, null);
INSERT INTO `bbs_group` VALUES ('2', '超级版主组', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', null, null, null);
INSERT INTO `bbs_group` VALUES ('4', '版主组', '0', '0', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '0', '1', null, null, null);
INSERT INTO `bbs_group` VALUES ('5', '实习版主组', '0', '0', '1', '1', '1', '1', '1', '1', '1', '0', '1', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('6', '待验证用户组', '0', '0', '1', '0', '1', '0', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('7', '禁止用户组', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('101', '一级用户组', '0', '50', '1', '1', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('102', '二级用户组', '50', '200', '1', '1', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('103', '三级用户组', '200', '1000', '1', '1', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('104', '四级用户组', '1000', '10000', '1', '1', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);
INSERT INTO `bbs_group` VALUES ('105', '五级用户组', '10000', '10000000', '1', '1', '1', '1', '1', '0', '0', '0', '0', '0', '0', '0', null, null, null);


-- ----------------------------
-- Table structure for bbs_kv
-- ----------------------------
DROP TABLE IF EXISTS `bbs_kv`;
CREATE TABLE `bbs_kv` (
  `k` char(32) NOT NULL DEFAULT '',
  `v` mediumtext NOT NULL,
  `expiry` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='持久的 key value 数据存储, ttserver, mysql';


-- ----------------------------
-- Table structure for bbs_modlog
-- ----------------------------
DROP TABLE IF EXISTS `bbs_modlog`;
CREATE TABLE `bbs_modlog` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0',
  `post_id` int(11) unsigned NOT NULL DEFAULT '0',
  `subject` char(32) NOT NULL DEFAULT '',
  `comment` char(64) NOT NULL DEFAULT '',
  `rmbs` int(11) NOT NULL DEFAULT '0',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0',
  `action` char(16) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `uid` (`user_id`,`id`),
  KEY `tid` (`thread_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COMMENT='版主操作日志';

-- ----------------------------
-- Records of bbs_modlog
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_my_favourite
-- ----------------------------
DROP TABLE IF EXISTS `bbs_my_favourite`;
CREATE TABLE `bbs_my_favourite` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int(11) DEFAULT '0',
  `thread_id` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_bbs_my_favourite_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- ----------------------------
-- Table structure for bbs_mypost
-- ----------------------------
DROP TABLE IF EXISTS `bbs_mypost`;
CREATE TABLE `bbs_mypost` (
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0',
  `post_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`,`post_id`),
  KEY `tid` (`thread_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='我的回帖';


-- ----------------------------
-- Table structure for bbs_mythread
-- ----------------------------
DROP TABLE IF EXISTS `bbs_mythread`;
CREATE TABLE `bbs_mythread` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_bbs_mythread_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='我的主题，每个主题不管回复多少次，只记录一次。大表，需要分区';


-- ----------------------------
-- Table structure for bbs_post
-- ----------------------------
DROP TABLE IF EXISTS `bbs_post`;
CREATE TABLE `bbs_post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '帖子id',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主题id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `isfirst` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '是否为首帖，与 thread.firstpid 呼应',
  `userip` varchar(50) NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()',
  `images_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的图片数',
  `files_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的文件数',
  `doctype` tinyint(3) NOT NULL DEFAULT '0' COMMENT '类型，0: html, 1: txt; 2: markdown; 3: ubb',
  `quote_post_id` int(11) NOT NULL DEFAULT '0' COMMENT '引用哪个 pid，可能不存在',
  `message` longtext NOT NULL COMMENT '内容，用户提示的原始数据',
  `message_fmt` longtext NOT NULL COMMENT '内容，存放的过滤后的html内容，可以定期清理，减肥',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `audited` int(11) DEFAULT '1',
  `likes_cnt` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `tid` (`thread_id`,`id`),
  KEY `uid` (`user_id`),
  KEY `idx_bbs_post_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛帖子数据';

-- ----------------------------

-- ----------------------------
-- Table structure for bbs_post_update_log
-- ----------------------------
DROP TABLE IF EXISTS `bbs_post_update_log`;
CREATE TABLE `bbs_post_update_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `post_id` int(11) DEFAULT '0',
  `user_id` int(11) DEFAULT '0',
  `reason` varchar(255) DEFAULT '',
  `message` varchar(255) DEFAULT '',
  `old_message` varchar(255) DEFAULT '',
  `audited` int(11) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_bbs_post_update_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of bbs_post_update_log
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_queue
-- ----------------------------
DROP TABLE IF EXISTS `bbs_queue`;
CREATE TABLE `bbs_queue` (
  `id` int(11) unsigned NOT NULL DEFAULT '0',
  `value` int(11) NOT NULL DEFAULT '0',
  `expiry` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  UNIQUE KEY `queueid` (`id`,`value`),
  KEY `expiry` (`expiry`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='临时队列，用来保存临时数据';

-- ----------------------------
-- Records of bbs_queue
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_session
-- ----------------------------
DROP TABLE IF EXISTS `bbs_session`;
CREATE TABLE `bbs_session` (
  `id` char(32) NOT NULL DEFAULT '0' COMMENT '随机生成 id 不能重复 uniqueid() 13 位',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户id 未登录为 0，可以重复',
  `forum_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '所在的版块',
  `url` char(32) NOT NULL DEFAULT '' COMMENT '当前访问 url',
  `ip` int(11) unsigned NOT NULL DEFAULT '0',
  `useragent` char(128) NOT NULL DEFAULT '',
  `data` char(255) NOT NULL DEFAULT '',
  `bigdata` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有大数据',
  `last_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上次活动时间',
  PRIMARY KEY (`id`),
  KEY `ip` (`ip`),
  KEY `fid` (`forum_id`),
  KEY `uid_last_date` (`user_id`,`last_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='session 表 缓存到 runtime 表。 online_0 全局 online_fid 版块。提高遍历效率\r\n';

-- ----------------------------
-- Records of bbs_session
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_session_data
-- ----------------------------
DROP TABLE IF EXISTS `bbs_session_data`;
CREATE TABLE `bbs_session_data` (
  `session_id` char(32) NOT NULL DEFAULT '0',
  `last_date` int(11) unsigned NOT NULL DEFAULT '0',
  `data` text NOT NULL,
  PRIMARY KEY (`session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bbs_session_data
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_table_day
-- ----------------------------
DROP TABLE IF EXISTS `bbs_table_day`;
CREATE TABLE `bbs_table_day` (
  `year` smallint(11) unsigned NOT NULL DEFAULT '0' COMMENT '年',
  `month` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '月',
  `day` tinyint(11) unsigned NOT NULL DEFAULT '0' COMMENT '日',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '时间戳',
  `table` char(16) NOT NULL DEFAULT '' COMMENT '表名',
  `maxid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最大ID',
  `count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总数',
  PRIMARY KEY (`year`,`month`,`day`,`table`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统表';

-- ----------------------------
-- Records of bbs_table_day
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_thread
-- ----------------------------
DROP TABLE IF EXISTS `bbs_thread`;
CREATE TABLE `bbs_thread` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主题id',
  `forum_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '版块 id',
  `top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '置顶级别: 0: 普通主题, 1-3 置顶的顺序',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `userip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '发帖时用户ip ip2long()，主要用来清理',
  `subject` char(128) NOT NULL DEFAULT '' COMMENT ' 主题',
  `last_date` timestamp NULL DEFAULT NULL COMMENT '最后回复时间',
  `views_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '查看次数, 剥离出去，单独的服务，避免 cache 失效',
  `posts_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '回帖数',
  `images_num` tinyint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的图片数',
  `files_num` tinyint(6) NOT NULL DEFAULT '0' COMMENT '附件中包含的文件数',
  `mods_cnt` tinyint(6) NOT NULL DEFAULT '0' COMMENT '预留：版主操作次数，如果 > 0, 则查询 modlog，显示斑竹的评分',
  `isclosed` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '预留：是否关闭，关闭以后不能再回帖、编辑',
  `first_post_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '首贴 pid',
  `last_user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最近参与的 uid',
  `last_post_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '最后回复的 pid',
  `digest` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `favourite_cnt` int(11) DEFAULT '0',
  `audited` int(11) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `lastpid` (`last_post_id`),
  KEY `fid` (`forum_id`,`id`),
  KEY `fid_2` (`forum_id`,`last_post_id`),
  KEY `idx_bbs_thread_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='论坛主题';


-- ----------------------------
-- Table structure for bbs_thread_digest
-- ----------------------------
DROP TABLE IF EXISTS `bbs_thread_digest`;
CREATE TABLE `bbs_thread_digest` (
  `forum_id` smallint(6) NOT NULL DEFAULT '0',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0',
  `digest` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`thread_id`),
  KEY `top` (`user_id`,`thread_id`),
  KEY `fid` (`forum_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='精华表  -- 插件表';

-- ----------------------------
-- Records of bbs_thread_digest
-- ----------------------------

-- ----------------------------
-- Table structure for bbs_thread_top
-- ----------------------------
DROP TABLE IF EXISTS `bbs_thread_top`;
CREATE TABLE `bbs_thread_top` (
  `forum_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '查找板块置顶',
  `thread_id` int(11) unsigned NOT NULL DEFAULT '0',
  `top` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'top: 0 是普通最新贴，> 0 置顶贴',
  PRIMARY KEY (`thread_id`),
  KEY `top` (`top`,`thread_id`),
  KEY `fid` (`forum_id`,`top`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='置顶主题';


-- ----------------------------
-- Table structure for bbs_user
-- ----------------------------
DROP TABLE IF EXISTS `bbs_user`;
CREATE TABLE `bbs_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `group_id` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '用户组编号',
  `email` char(40) NOT NULL DEFAULT '' COMMENT '邮箱',
  `email_checked` tinyint(3) NOT NULL DEFAULT 0, 
  `username` char(32) NOT NULL DEFAULT '' COMMENT '用户名',
  `realname` char(16) NOT NULL DEFAULT '' COMMENT '用户名',
  `id_number` char(19) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(100) NOT NULL DEFAULT '' COMMENT '密码',
  `password_sms` char(16) NOT NULL DEFAULT '' COMMENT '密码',
  `phone` char(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `qq` char(15) NOT NULL DEFAULT '' COMMENT 'QQ',
  `threads_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '发帖数',
  `posts_cnt` int(11) NOT NULL DEFAULT '0' COMMENT '回帖数',
  `credits_num` int(11) NOT NULL DEFAULT '0' COMMENT '积分',
  `golds_num` int(11) NOT NULL DEFAULT '0' COMMENT '金币',
  `rmbs_num` int(11) NOT NULL DEFAULT '0' COMMENT '人民币',
  `create_ip` varchar(20) NOT NULL DEFAULT '0' COMMENT '创建时IP',
  `create_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `login_ip` varchar(20) NOT NULL DEFAULT '0' COMMENT '登录时IP',
  `login_date` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `logins_cnt` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `avatar` varchar(200) NOT NULL DEFAULT '/static/img/avatar.png' COMMENT '用户最后更新图像时间',
  `digests_num` int(11) DEFAULT NULL COMMENT '精华数',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `digests_cnt` int(11) DEFAULT '0',
  `state` int(11) DEFAULT NULL,
  `favourite_cnt` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  KEY `gid` (`group_id`),
  KEY `idx_bbs_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bbs_user
-- ----------------------------
INSERT INTO `bbs_user` (`id`, `group_id`, `email`, `username`, `realname`, `id_number`, `password`, `password_sms`, `phone`, `qq`, `threads_cnt`, `posts_cnt`, `credits_num`, `golds_num`, `rmbs_num`, `create_ip`, `create_date`, `login_ip`, `login_date`, `logins_cnt`, `avatar`, `digests_num`, `created_at`, `updated_at`, `deleted_at`, `digests_cnt`, `state`, `favourite_cnt`, `email_checked`) VALUES ('25', '1', 'admin@local.com', 'admin', '', '', '$2a$10$zzjAmJrsR0hk8UBbL9P3OOTLBBNEjtME1G5s3Vl2./.TwHrroDwkm', '', '', '', '0', '0', '0', '0', '0', '0', '0', '0', '2019-08-21 13:49:12', '0', '/static/img/avatar.png', NULL, '2019-08-21 13:49:12', '2019-08-21 13:49:12', NULL, '0', '0', '0', '1');

CREATE TABLE `bbs_tag_cate` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` char(32) NOT NULL DEFAULT '',
  `rank` int(11) unsigned NOT NULL DEFAULT '0',
  `enable` int(11) unsigned NOT NULL DEFAULT '0',
  `default_tag_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '默认值,如果没有，设为全部',
  `isforce` int(11) unsigned NOT NULL DEFAULT '0',
  `style` char(32) NOT NULL DEFAULT '',
  `comment` varchar(500) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;


CREATE TABLE `bbs_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag_cate_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'cate所属分类',
  `name` char(32) NOT NULL DEFAULT '',
  `rank` int(11) unsigned NOT NULL DEFAULT '0',
  `enable` int(11) unsigned NOT NULL DEFAULT '0',
  `style` char(32) NOT NULL DEFAULT '',
  `comment` varchar(500) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `cate_id` (`tag_cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;


CREATE TABLE `bbs_tag_forum` (
  `tag_id` int(11) unsigned NOT NULL DEFAULT '0',
  `forum_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`tag_id`,`forum_id`),
  KEY `tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='forum分类启用的tag';

CREATE TABLE `bbs_user_banned` (
  `user_id` int(11) NOT NULL,
  `from_date` timestamp NULL DEFAULT NULL,
  `to_date` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `bbs_ip_banned` (
  `ip` varchar(20) NOT NULL,
  `from_date` timestamp NULL DEFAULT NULL,
  `to_date` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
