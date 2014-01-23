-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.34 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win32
-- HeidiSQL 版本:                  8.1.0.4545
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- 导出 desk 的数据库结构
CREATE DATABASE IF NOT EXISTS `desk` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `desk`;


-- 导出  表 desk.basetasks 结构
CREATE TABLE IF NOT EXISTS `basetasks` (
  `bid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`bid`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- 正在导出表  desk.basetasks 的数据：~9 rows (大约)
/*!40000 ALTER TABLE `basetasks` DISABLE KEYS */;
INSERT INTO `basetasks` (`bid`, `name`) VALUES
	(1, '打印机硒鼓更换'),
	(2, 'AD加域'),
	(3, '装防病毒杀毒软件'),
	(4, '桌面软件问题'),
	(5, '网络连接问题'),
	(6, 'OA使用问题'),
	(7, '邮箱使用问题'),
	(8, '设备故障申报'),
	(9, '其他服务');
/*!40000 ALTER TABLE `basetasks` ENABLE KEYS */;


-- 导出  表 desk.state 结构
CREATE TABLE IF NOT EXISTS `state` (
  `wid` varchar(100) NOT NULL,
  `state` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  desk.state 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `state` DISABLE KEYS */;
INSERT INTO `state` (`wid`, `state`) VALUES
	('oU7acuKn1S0jLyoDrds4Idl63Pg4', 'timeout');
/*!40000 ALTER TABLE `state` ENABLE KEYS */;


-- 导出  表 desk.task 结构
CREATE TABLE IF NOT EXISTS `task` (
  `tid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `taker` varchar(100) NOT NULL DEFAULT '' COMMENT 'maintainer ID',
  `asker` varchar(100) NOT NULL DEFAULT '',
  `taken` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`tid`),
  KEY `taken` (`taken`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='待处理任务';

-- 正在导出表  desk.task 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `task` DISABLE KEYS */;
/*!40000 ALTER TABLE `task` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
