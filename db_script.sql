-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: exam
-- ------------------------------------------------------
-- Server version	9.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `it01`
--

DROP TABLE IF EXISTS `it01`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `it01` (
  `id` char(36) NOT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `age` int DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `it01`
--

LOCK TABLES `it01` WRITE;
/*!40000 ALTER TABLE `it01` DISABLE KEYS */;
INSERT INTO `it01` VALUES ('0413c516-361c-47bd-8bb7-326c08542eaa','2122323','2022-02-06',4,'233','2026-02-06 15:01:19'),('2f940d6a-16a2-4061-a433-a8f4ea80887b','Ekachai Saengsuksai','2000-01-01',26,'Bangkok','2026-02-06 14:27:54'),('5e508eba-6e92-465d-8a4a-516771312bd5','Ekachai Saengsuksai','2000-01-01',26,'Bangkok','2026-02-06 14:28:18'),('c7ea1a15-b492-46fd-8cac-96627f71f198','2323','2024-02-07',2,'233','2026-02-06 14:27:54'),('c897aa84-075a-4582-ba41-c2390dc0bb67','23233','2025-10-07',1,'2323','2026-02-06 14:27:54');
/*!40000 ALTER TABLE `it01` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `it02`
--

DROP TABLE IF EXISTS `it02`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `it02` (
  `id` char(36) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_date` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `it02`
--

LOCK TABLES `it02` WRITE;
/*!40000 ALTER TABLE `it02` DISABLE KEYS */;
INSERT INTO `it02` VALUES ('395f0f4e-6442-4699-987f-d352be62d3a9','hornokhutatototo','$2a$10$FkKUtGxNyLtGXyW4.813/esC3DjmatJS6dmLb651HuwfdPim84YNa','2026-02-06 15:41:56'),('c2663d89-108d-43a4-9697-470a59514e86','eiei','$2a$10$B.lE7eeGuPNfI/c93r40fuXHycL2xsBfCUHaW/rrhrGUTXvnPYg4a','2026-02-06 15:54:42'),('f83d849f-7ca7-4a7b-ac34-1042100e4dee','hornokhutato','$2a$10$eBB0JgfKhUgvIPqq6IHE7.NbgQmRpfG9LuN8xcldi1Je9U6ZerjNK','2026-02-06 15:37:09');
/*!40000 ALTER TABLE `it02` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `it03`
--

DROP TABLE IF EXISTS `it03`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `it03` (
  `id` char(36) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` text,
  `status` varchar(32) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `it03`
--

LOCK TABLES `it03` WRITE;
/*!40000 ALTER TABLE `it03` DISABLE KEYS */;
INSERT INTO `it03` VALUES ('a6f2j7h5-6i7f-9j8g-4h9i-6789012345fa','รายการที่ 6','เทส','อนุมัติ','2026-02-06 16:37:52','2026-02-06 18:13:53'),('b1a7e2c0-1f2a-4e3b-9c4d-1234567890ab','รายการที่ 1','เทส','อนุมัติ','2026-02-06 16:37:52','2026-02-06 18:13:53'),('b7g3k8i6-7j8g-0k9h-5i0j-7890123456ab','รายการที่ 7','เทส','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 17:26:50'),('c2b8f3d1-2e3b-5f4c-0d5e-2345678901bc','รายการที่ 2','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52'),('c8h4l9j7-8k9h-1l0i-6j1k-8901234567bc','รายการที่ 8','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52'),('d3c9g4e2-3f4c-6g5d-1e6f-3456789012cd','รายการที่ 3','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52'),('d9i5m0k8-9l0i-2m1j-7k2l-9012345678cd','รายการที่ 9','233','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 17:18:36'),('e0j6n1l9-0m1j-3n2k-8l3m-0123456789de','รายการที่ 10','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52'),('e4d0h5f3-4g5d-7h6e-2f7g-4567890123de','รายการที่ 4','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52'),('f5e1i6g4-5h6e-8i7f-3g8h-5678901234ef','รายการที่ 5','xxxxxx','รออนุมัติ','2026-02-06 16:37:52','2026-02-06 16:37:52');
/*!40000 ALTER TABLE `it03` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `it04`
--

DROP TABLE IF EXISTS `it04`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `it04` (
  `id` char(36) NOT NULL,
  `firstname` varchar(100) NOT NULL,
  `lastname` varchar(100) NOT NULL,
  `email` varchar(150) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `profile` longtext NOT NULL,
  `birthday` date NOT NULL,
  `occupation` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `it04`
--

LOCK TABLES `it04` WRITE;
/*!40000 ALTER TABLE `it04` DISABLE KEYS */;
/*!40000 ALTER TABLE `it04` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'exam'
--

--
-- Dumping routines for database 'exam'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-06 22:48:02
