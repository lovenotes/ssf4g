CREATE DATABASE ssf4g_login_srv DEFAULT CHARACTER SET utf8;
USE ssf4g_login_srv;

CREATE TABLE `accounts` (
  `accnt_name` varchar(32) NOT NULL,
  `accnt_id` bigint(11) UNSIGNED NOT NULL,  
  `pass_hash` varchar(40) NOT NULL DEFAULT '',
  `last_ip` varchar(20) DEFAULT '',
  `platform` tinyint(3) unsigned NOT NULL DEFAULT '1', 
  `created_at` timestamp NOT NULL, 
  `updated_at` timestamp NULL,
  PRIMARY KEY (`accnt_name`),
  UNIQUE KEY `idx_accntid` (`accnt_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tickets` (
  `ticket_id` bigint(11) UNSIGNED NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`ticket_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `tickets` (`ticket_id`) VALUES (102260130);