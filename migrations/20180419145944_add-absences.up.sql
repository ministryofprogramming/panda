CREATE TABLE `absences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company_id` int(11) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 NOT NULL,
  `max_days` int NOT NULL DEFAULT 0,
  `is_paid` TINYINT NOT NULL DEFAULT 0,
  `period_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `unique` (`name`, `company_id`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`),

  FOREIGN KEY (`company_id`) REFERENCES companies(`id`),
  FOREIGN KEY (`period_id`) REFERENCES periods(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


