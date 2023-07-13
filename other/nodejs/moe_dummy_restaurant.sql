-- Adminer 4.8.1 MySQL 11.0.2-MariaDB-1:11.0.2+maria~ubu2204 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `moe_dummy_restaurant`;
CREATE TABLE `moe_dummy_restaurant` (
  `order_number` int(11) NOT NULL AUTO_INCREMENT,
  `res_date` varchar(255) NOT NULL,
  `res_time` varchar(255) NOT NULL,
  `res_person_name` varchar(255) NOT NULL,
  `res_location` varchar(255) NOT NULL,
  `res_num_of_people` varchar(255) NOT NULL,
  `res_status` varchar(255) NOT NULL,
  PRIMARY KEY (`order_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `moe_dummy_restaurant` (`order_number`, `res_date`, `res_time`, `res_person_name`, `res_location`, `res_num_of_people`, `res_status`) VALUES
(6,	'2020-01-15',	'18:00',	'Shane',	'Mumbai',	'5',	'SUCCESS'),
(7,	'2020-01-15',	'18:00',	'Shane',	'Mumbai',	'5',	'SUCCESS'),
(8,	'2020-01-15',	'18:00',	'selman',	'Mumbai',	'5',	'SUCCESS'),
(9,	'2020-01-15',	'18:00',	'Shane',	'Mumbai',	'5',	'SUCCESS');

-- 2023-07-04 22:01:59
