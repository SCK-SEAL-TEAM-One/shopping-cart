DROP DATABASE IF EXISTS toy;
CREATE DATABASE IF NOT EXISTS toy CHARACTER SET utf8 COLLATE utf8_general_ci;
USE toy

CREATE TABLE user (
    id int,
    name varchar(255)
);

INSERT INTO user VALUE (1,"sckshuhari");


CREATE TABLE products (
    id BIGINT AUTO_INCREMENT,
    product_name varchar(255),
    product_brand varchar(255),
    quantity int,
    product_price double,
    image_url varchar(255),
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);

INSERT INTO products (id,product_name,product_brand,quantity,product_price,image_url) VALUE (2,"43 Piece dinner Set","CoolKidz",10,12.95,"/43_Piece_dinner_Set.png");

CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT,
    total_price double,
    transaction_id varchar(255) DEFAULT '',
    completed smallint(1) DEFAULT 0,
    authorized timestamp DEFAULT current_timestamp,
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE order_product (
    order_id BIGINT,
    product_id BIGINT,
    quantity int,
    product_price double
);

CREATE TABLE shipping (
    id int AUTO_INCREMENT,
    order_id BIGINT,
    address varchar(255),
    sub_district varchar(255),
    district varchar(255),
    province varchar(255),
    zip_code varchar(5),
    recipient varchar(255),
    phone_number varchar(13),
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);