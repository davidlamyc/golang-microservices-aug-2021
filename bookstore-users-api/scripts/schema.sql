CREATE SCHEMA IF NOT EXISTS users CHARACTER SET utf8;

CREATE TABLE IF NOT EXISTS users.users (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(45) NULL,
    last_name VARCHAR(45) NULL,
    email VARCHAR(45) NOT NULL,
    date_created DATETIME DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(45) NOT NULL,
    password VARCHAR(32) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX email_UNIQUE (email ASC)
);