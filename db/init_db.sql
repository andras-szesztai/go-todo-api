CREATE DATABASE IF NOT EXISTS todos DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE todos;

CREATE TABLE IF NOT EXISTS todos (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	status ENUM('TODO', 'IN_PROGRESS', 'DONE') NOT NULL DEFAULT 'TODO',
	createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

	PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO todos (name, status) VALUES ('Task 1', 'TODO');
INSERT INTO todos (name, status) VALUES ('Task 2', 'IN_PROGRESS');
INSERT INTO todos (name, status) VALUES ('Task 3', 'DONE');

