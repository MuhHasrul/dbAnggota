-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jun 25, 2024 at 04:52 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.1.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `Sekertaris`
--

-- --------------------------------------------------------

--
-- Table structure for table `userAccount`
--

CREATE TABLE `userAccount` (
  `Uid` int(11) NOT NULL,
  `Email` varchar(50) NOT NULL,
  `Password` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `userAccount`
--

INSERT INTO `userAccount` (`Uid`, `Email`, `Password`) VALUES
(1, 'dayat@gmail.com', '123456789'),
(2, 'kerenjelek3@gmail.com', 'asdadadad'),
(3, '\' = \'\' or \'0\' = \'0\'-- -', ''),
(4, 'cnt@gmail.com', '123'),
(5, 'cnt12@gmail.com', '12345678'),
(6, 'nabila@gmail.com', '123123123'),
(7, 'cnt12@gmail.com', '123217323'),
(8, 'ccnt@gmail.com', 'ccntdotorg'),
(9, 'ccnt12@gmail.com', 'ccntdotorg'),
(10, 'ccnt012@gmail.com', 'asubanget');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `userAccount`
--
ALTER TABLE `userAccount`
  ADD PRIMARY KEY (`Uid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `userAccount`
--
ALTER TABLE `userAccount`
  MODIFY `Uid` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
