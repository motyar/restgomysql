-- phpMyAdmin SQL Dump
-- version 3.4.10.1deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost:3306
-- Generation Time: Sep 12, 2014 at 02:52 AM
-- Server version: 5.5.31
-- PHP Version: 5.3.10-1ubuntu3.9

SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `farm`
--

-- --------------------------------------------------------

--
-- Table structure for table `pandas`
--

CREATE TABLE IF NOT EXISTS `pandas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(500) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=14 ;

--
-- Dumping data for table `pandas`
--

INSERT INTO `pandas` (`id`, `name`) VALUES
(1, 'Name changed via API'),
(2, 'A new cute panda came via API door'),
(3, 'Another cute panda'),
(4, 'Pandas sister');
