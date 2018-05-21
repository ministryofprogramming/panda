CREATE TABLE `companies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `logo` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `vat_number` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `city` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `country_id` int(11) NOT NULL,
  `monday` TINYINT NOT NULL DEFAULT 1,
  `tuesday` TINYINT NOT NULL DEFAULT 1,
  `wednesday` TINYINT NOT NULL DEFAULT 1,
  `thursday` TINYINT NOT NULL DEFAULT 1,
  `friday` TINYINT NOT NULL DEFAULT 1,
  `saturday` TINYINT NOT NULL DEFAULT 0,
  `sunday` TINYINT NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`),

  FOREIGN KEY (`country_id`) REFERENCES countries(`id`) 
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

