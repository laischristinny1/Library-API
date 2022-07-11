CREATE DATABASE IF NOT EXISTS LibraryAPI;
USE LibraryAPI;

DROP TABLE IF EXISTS books;

CREATE TABLE books(
    id int auto_increment primary key,
    title varchar(50) not null,
    author varchar(50) not null,
    quantity int not null
) ENGINE=INNODB;

