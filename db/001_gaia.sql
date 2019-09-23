-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.2.14-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for GOAPI
CREATE DATABASE IF NOT EXISTS `GOAPI` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `GOAPI`;

-- Dumping structure for table GOAPI.GOAPI_users
CREATE TABLE IF NOT EXISTS `GOAPI_users` (
  `id` varbinary(16) NOT NULL COMMENT 'User''s unique identity',
  `username` varchar(32) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varbinary(32) DEFAULT NULL,
  `salt` varbinary(16) DEFAULT NULL,
  PRIMARY KEY (`username`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='GOAPI GOAPI_users table';

-- Dumping structure for table GOAPI.GOAPI_profiles
CREATE TABLE IF NOT EXISTS `GOAPI_profiles` (
  `id` varbinary(16) NOT NULL COMMENT 'User''s unique identity',
  `field` varchar(64) NOT NULL,
  `value` varchar(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='GOAPI user''s profile table';

-- Dumping structure for table GOAPI.GOAPI_sessions
CREATE TABLE IF NOT EXISTS `GOAPI_sessions` (
  `id` varbinary(32) NOT NULL COMMENT 'Session''s unique identity',
  `uid` varbinary(16) NOT NULL COMMENT 'User''s id', 
  `last_active` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='GOAPI user''s session table';

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;