CREATE TABLE `client` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `name` VARCHAR(30) NOT NULL,
        UNIQUE KEY `name_unique_index`(`enterprise_id`, `workshop_id`, `name`),
    `phone_number` VARCHAR(20),
        -- INDEX `phone_number_index`(`phone_number`),
    `email` VARCHAR(254),
        -- INDEX `email_index`(`email`),
    `address` VARCHAR(100),
        -- INDEX `addrass_index`(`address`),
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300),
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;