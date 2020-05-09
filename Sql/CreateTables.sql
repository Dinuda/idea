DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS investors;

DROP TABLE IF EXISTS students;

CREATE TABLE `users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`username` varchar(255) NOT NULL,
	`password` varchar(255) NOT NULL,
	`firstname` varchar(255) NOT NULL,
	`lastname` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL UNIQUE,
	`phone_no` varchar(255) NOT NULL UNIQUE,
	`date_of_birth` DATE NOT NULL,
	`description` TEXT,
	`type` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `investors` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`linkedin` varchar(255) NOT NULL UNIQUE,
	`company` varchar(255) ,
	`user_id` INT NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `students` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`team_id` INT NOT NULL,
	`user_id` INT NOT NULL,
	`profession` INT NOT NULL,
	`university` TEXT NOT NULL,
	`cv` varchar(255) NOT NULL UNIQUE,
	`team_role` VARCHAR(255) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `profesions` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255) NOT NULL UNIQUE,
	PRIMARY KEY (`id`)
);

CREATE TABLE `projects` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL,
	`description` TEXT NOT NULL,
	`created_date` DATE NOT NULL,
	`closed_date` DATE,
	`investorteam_id` INT NOT NULL UNIQUE,
	`studentteam_id` INT,
	PRIMARY KEY (`id`)
);

CREATE TABLE `investorteam` (
	`id` INT NOT NULL,
	`investor_id` INT NOT NULL UNIQUE,
	PRIMARY KEY (`id`)
);

CREATE TABLE `studentteam` (
	`id` INT NOT NULL,
	`student_id` INT NOT NULL,
	PRIMARY KEY (`id`)
);

ALTER TABLE `investors` ADD CONSTRAINT `investors_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

ALTER TABLE `students` ADD CONSTRAINT `students_fk0` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);

ALTER TABLE `students` ADD CONSTRAINT `students_fk1` FOREIGN KEY (`profession`) REFERENCES `profesions`(`id`);

ALTER TABLE `projects` ADD CONSTRAINT `projects_fk0` FOREIGN KEY (`investorteam_id`) REFERENCES `investorteam`(`id`);

ALTER TABLE `projects` ADD CONSTRAINT `projects_fk1` FOREIGN KEY (`studentteam_id`) REFERENCES `studentteam`(`id`);

ALTER TABLE `investorteam` ADD CONSTRAINT `investorteam_fk0` FOREIGN KEY (`investor_id`) REFERENCES `investors`(`id`);

ALTER TABLE `studentteam` ADD CONSTRAINT `studentteam_fk0` FOREIGN KEY (`student_id`) REFERENCES `students`(`id`);
