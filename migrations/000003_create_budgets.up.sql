CREATE TABLE `budgets` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `total_income` bigint unsigned NOT NULL,
  `needs_percentage` tinyint unsigned NOT NULL DEFAULT '50',
  `wants_percentage` tinyint unsigned NOT NULL DEFAULT '30',
  `savings_percentage` tinyint unsigned NOT NULL DEFAULT '20',
  `year` smallint unsigned NOT NULL,
  `month` tinyint unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_budget_month` (`user_id`,`year`,`month`),
  CONSTRAINT `fk_budgets_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `chk_budget_month` CHECK ((`month` between 1 and 12)),
  CONSTRAINT `chk_budget_percentage` CHECK ((((`needs_percentage` + `wants_percentage`) + `savings_percentage`) = 100))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
