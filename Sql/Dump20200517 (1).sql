CREATE DATABASE  IF NOT EXISTS `idea` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */;
USE `idea`;
-- MySQL dump 10.13  Distrib 8.0.12, for Win64 (x86_64)
--
-- Host: localhost    Database: idea
-- ------------------------------------------------------
-- Server version	8.0.12

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Advertising and Pr'),(2,'Agri Tech'),(3,'Agriculture'),(5,'Arts and Design'),(6,'Banking'),(7,'Business Consulting'),(8,'Charity and Voluntary Work'),(9,'Education'),(10,'Energy and Utility'),(11,'Engineering'),(12,'Finance'),(13,'Food and Agriculture'),(14,'Hardware'),(15,'Healthcare'),(16,'Hospitality'),(17,'IT'),(18,'Law'),(19,'Leisure'),(20,'Management'),(21,'Manufacturing'),(22,'Marketing'),(23,'Media'),(24,'Medicine'),(25,'Pharmaceuticals'),(26,'Property and Construction'),(27,'Retail'),(28,'Sales'),(29,'Security'),(30,'Sport and Tourisim'),(31,'Transport and Logistics'),(32,'Travel');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `investorteam`
--

DROP TABLE IF EXISTS `investorteam`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `investorteam` (
  `id` int(40) NOT NULL,
  `investor_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `investor_id` (`investor_id`),
  CONSTRAINT `investorteam_fk0` FOREIGN KEY (`investor_id`) REFERENCES `investors` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `investorteam`
--

LOCK TABLES `investorteam` WRITE;
/*!40000 ALTER TABLE `investorteam` DISABLE KEYS */;
/*!40000 ALTER TABLE `investorteam` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `professions`
--

DROP TABLE IF EXISTS `professions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `professions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=243 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `professions`
--

LOCK TABLES `professions` WRITE;
/*!40000 ALTER TABLE `professions` DISABLE KEYS */;
INSERT INTO `professions` VALUES (1,'Software Engineer'),(2,'Mobile Developer'),(3,'Frontend Engineer'),(4,'Backend Engineer'),(5,'Full-Stack Engineer'),(6,'Engineering Manager'),(7,'QA Engineer'),(8,'DevOps'),(9,'Software Architect'),(10,'Embedded Engineer'),(11,'Data Engineer'),(12,'Designer'),(13,'UI/UX Designer'),(14,'User Researcher'),(15,'Visual Designer'),(16,'Creative Director'),(17,'Operations'),(18,'Finance/Accounting'),(19,'H.R.'),(20,'Office Manager'),(21,'Recruiter'),(22,'Customer Service'),(23,'Operations Manager'),(24,'Sales'),(25,'Business Development'),(26,'Sales Development Representative'),(27,'Account Executive'),(28,'BD Manager'),(29,'Account Manager'),(30,'Sales Manager'),(31,'Marketing'),(32,'Growth Hacker'),(33,'Marketing Manager'),(34,'Content Creator'),(35,'CEO'),(36,'CFO'),(37,'CMO'),(38,'COO'),(39,'CTO'),(40,'Hardware Engineer'),(41,'Mechanical Engineer'),(42,'Systems Engineer'),(43,'Business Analyst'),(44,'Data Scientist'),(45,'Product Manager'),(46,'Project Manager'),(47,'Attorney');
/*!40000 ALTER TABLE `professions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `students`
--

DROP TABLE IF EXISTS `students`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `students` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `profession` int(11) NOT NULL,
  `cv` varchar(255) NOT NULL,
  `team_role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `cv` (`cv`),
  KEY `students_fk0` (`user_id`),
  KEY `students_fk2` (`profession`),
  CONSTRAINT `students_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `students_fk2` FOREIGN KEY (`profession`) REFERENCES `professions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `students`
--

LOCK TABLES `students` WRITE;
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
INSERT INTO `students` VALUES (1,7,34,'https://test.test.com',NULL);
/*!40000 ALTER TABLE `students` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `studentteam`
--

DROP TABLE IF EXISTS `studentteam`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `studentteam` (
  `id` char(40) NOT NULL,
  `student_id` int(11) NOT NULL,
  `title` text,
  PRIMARY KEY (`id`),
  KEY `studentteam_fk0` (`student_id`),
  CONSTRAINT `studentteam_fk0` FOREIGN KEY (`student_id`) REFERENCES `students` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `studentteam`
--

LOCK TABLES `studentteam` WRITE;
/*!40000 ALTER TABLE `studentteam` DISABLE KEYS */;
/*!40000 ALTER TABLE `studentteam` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone_no` varchar(255) NOT NULL,
  `description` text,
  `type` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone_no` (`phone_no`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (7,'test','$2a$10$fREKV.BSPlQWH4sOltFhsuku9L7g7YuVj4AVGmF3AUxsL.E5Z0SIS','test','test','test@test.com','9400000000','Example Description Example Description Example Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example DescriptionExample Description Example Descriptio','Student');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-17 13:47:59
