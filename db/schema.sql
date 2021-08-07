DROP DATABASE IF EXISTS account;
CREATE DATABASE account;
USE account;

CREATE TABLE users (
  user_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  email VARCHAR(256) NOT NULL,
  account_type ENUM('student', 'teacher', 'general') NOT NULL,
  student_number INT,
  password VARCHAR(200) NOT NULL
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

create table admins (
  admin_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(10) NOT NULL,
  password VARCHAR(200) NOT NULL
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;
