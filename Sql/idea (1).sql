DROP DATABASE IF EXISTS `idea`;
CREATE DATABASE `idea`; 
USE `idea`;

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255),
  `password` varchar(255),
  `name` varchar(255),
  `email` varchar(255),
  `phone_no` int,
  `address` varchar(255),
  `description` longtext,
  `type` varchar(255)
);

CREATE TABLE `students` (
  `id` int PRIMARY KEY,
  `profession` int,
  `cv` varchar(255)
);

CREATE TABLE `investors` (
  `id` int PRIMARY KEY,
  `linkedin` varchar(255),
  `company` varchar(255)
);

CREATE TABLE `projects` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(255),
  `description` longtext,
  `created_at` datetime,
  `closed_at` datetime,
  `category` int,
  `host` int
);

CREATE TABLE `projectstudentteam` (
  `project_id` int,
  `user_id` int
);

CREATE TABLE `projectinvestorteam` (
  `project_id` int,
  `user_id` int
);

CREATE TABLE `professions` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY,
  `name` varchar(255)
);

ALTER TABLE `students` ADD FOREIGN KEY (`id`) REFERENCES `users` (`id`);

ALTER TABLE `students` ADD FOREIGN KEY (`profession`) REFERENCES `professions` (`id`);

ALTER TABLE `investors` ADD FOREIGN KEY (`id`) REFERENCES `users` (`id`);

ALTER TABLE `projects` ADD FOREIGN KEY (`category`) REFERENCES `categories` (`id`);

ALTER TABLE `projects` ADD FOREIGN KEY (`host`) REFERENCES `users` (`id`);

ALTER TABLE `projectstudentteam` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

ALTER TABLE `projectstudentteam` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `projectinvestorteam` ADD FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`);

ALTER TABLE `projectinvestorteam` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


INSERT INTO `professions` VALUES (1,'Software Engineer'),(2,'Mobile Developer'),(3,'Frontend Engineer'),(4,'Backend Engineer'),(5,'Full-Stack Engineer'),(6,'Engineering Manager'),(7,'QA Engineer'),(8,'DevOps'),(9,'Software Architect'),(10,'Embedded Engineer'),(11,'Data Engineer'),(12,'Designer'),(13,'UI/UX Designer'),(14,'User Researcher'),(15,'Visual Designer'),(16,'Creative Director'),(17,'Operations'),(18,'Finance/Accounting'),(19,'H.R.'),(20,'Office Manager'),(21,'Recruiter'),(22,'Customer Service'),(23,'Operations Manager'),(24,'Sales'),(25,'Business Development'),(26,'Sales Development Representative'),(27,'Account Executive'),(28,'BD Manager'),(29,'Account Manager'),(30,'Sales Manager'),(31,'Marketing'),(32,'Growth Hacker'),(33,'Marketing Manager'),(34,'Content Creator'),(35,'CEO'),(36,'CFO'),(37,'CMO'),(38,'COO'),(39,'CTO'),(40,'Hardware Engineer'),(41,'Mechanical Engineer'),(42,'Systems Engineer'),(43,'Business Analyst'),(44,'Data Scientist'),(45,'Product Manager'),(46,'Project Manager'),(47,'Attorney');

INSERT INTO `categories` VALUES (1,'Advertising and Pr'),(2,'Agri Tech'),(3,'Agriculture'),(5,'Arts and Design'),(6,'Banking'),(7,'Business Consulting'),(8,'Charity and Voluntary Work'),(9,'Education'),(10,'Energy and Utility'),(11,'Engineering'),(12,'Finance'),(13,'Food and Agriculture'),(14,'Hardware'),(15,'Healthcare'),(16,'Hospitality'),(17,'IT'),(18,'Law'),(19,'Leisure'),(20,'Management'),(21,'Manufacturing'),(22,'Marketing'),(23,'Media'),(24,'Medicine'),(25,'Pharmaceuticals'),(26,'Property and Construction'),(27,'Retail'),(28,'Sales'),(29,'Security'),(30,'Sport and Tourisim'),(31,'Transport and Logistics'),(32,'Travel');