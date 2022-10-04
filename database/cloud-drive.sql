-- -------------------------------------------------------------
-- Database Type: MySQL
-- Database: cloud-drive
-- -------------------------------------------------------------


DROP TABLE IF EXISTS `repository_pool`;
CREATE TABLE `repository_pool` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '记录的唯一标识',
  `filehash` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文件唯一标识hash',
  `filename` varchar(255) DEFAULT NULL COMMENT '文件名称',
  `fileext` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'Filename Extension 文件扩展名',
  `filesize` double DEFAULT NULL COMMENT 'file size文件大小',
  `filepath` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='中心存储池';

DROP TABLE IF EXISTS `share`;
CREATE TABLE `share` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL COMMENT '记录的唯一标识',
  `user_identity` varchar(36) DEFAULT NULL,
  `repository_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '公共存储池中的唯一标识',
  `user_repository_identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户存储池中的唯一标识',
  `expired_time` int DEFAULT NULL COMMENT '分享的文件失效时间（秒），0的话用不失效',
  `click_num` int DEFAULT NULL COMMENT '分享点击次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

DROP TABLE IF EXISTS `user_repository`;
CREATE TABLE `user_repository` (
  `id` int NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) DEFAULT NULL COMMENT '记录的唯一标识',
  `user_identity` varchar(36) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `repository_identity` varchar(36) DEFAULT NULL,
  `fileext` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文件或文件夹类型',
  `filename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '当前文件或文件夹名称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
