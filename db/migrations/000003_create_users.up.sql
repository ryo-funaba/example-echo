CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '識別番号',
  `status` smallint(1) NOT NULL DEFAULT '1' COMMENT '状態',
  `name` VARCHAR(50) NOT NULL COMMENT '表示名',
  `username` VARCHAR(50) NOT NULL COMMENT 'ユーザー名',
  `email` VARCHAR(256) NOT NULL COMMENT 'メールアドレス',
  `password` VARCHAR(256) NOT NULL COMMENT 'パスワード',
  `user_type_id` int(11) NOT NULL COMMENT 'ユーザータイプID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` datetime NULL COMMENT '削除日時',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_users_username` (`username`),
  UNIQUE INDEX `idx_users_email` (`email`),
  INDEX `idx_users_user_type_id` (`user_type_id`),
  FOREIGN KEY `fk_users_user_type_id` (`user_type_id`) REFERENCES `user_types` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) DEFAULT CHARACTER SET utf8mb4 COMMENT="ユーザー";
