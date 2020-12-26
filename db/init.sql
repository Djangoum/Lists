create database if not exists lists;

use lists;

create table if not exists users 
(
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255),
    password varchar(255),
    first_name varchar(100), 
    last_name varchar(100), 
    created_at datetime, 
    updated_at datetime, 
    PRIMARY KEY (id)
) ENGINE = InnoDB;

create table if not exists lists 
(
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255), 
    created_at datetime, 
    updated_at datetime, 
    PRIMARY KEY (id)
) ENGINE = InnoDB;

create table if not exists list_items 
(
    id int NOT NULL AUTO_INCREMENT,
    list_id int, 
    information text, 
    created_at datetime null, 
    updated_at datetime null,
    PRIMARY KEY (id),
    FOREIGN KEY (list_id) REFERENCES lists(id)
) ENGINE = InnoDB;

create table if not exists lists_item_assignees
(
    list_item_id int,
    user_id int,
    PRIMARY KEY (list_item_id, user_id),
    FOREIGN KEY (list_item_id) REFERENCES list_items(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB;