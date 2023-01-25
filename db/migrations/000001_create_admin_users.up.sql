CREATE TABLE IF NOT EXISTS `admin_users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '識別番号',
  `status` TINYINT NOT NULL COMMENT '状態',
  `name` VARCHAR(50) NOT NULL COMMENT '表示名',
  `username` VARCHAR(50) NOT NULL COMMENT 'ユーザー名',
  `email` VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
  `password` BINARY(64) NOT NULL COMMENT 'パスワード',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `deleted_at` DATETIME COMMENT '削除日時',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_admin_users_username` (`username`),
  UNIQUE INDEX `idx_admin_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT="管理ユーザー";
