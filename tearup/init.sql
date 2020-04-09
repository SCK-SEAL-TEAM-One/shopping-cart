DROP DATABASE IF EXISTS toy;
CREATE DATABASE IF NOT EXISTS toy CHARACTER SET utf8 COLLATE utf8_general_ci;
USE toy;

CREATE TABLE user (
    id int,
    name varchar(255)
) CHARACTER SET utf8 COLLATE utf8_general_ci;

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
) CHARACTER SET utf8 COLLATE utf8_general_ci;

INSERT INTO products (id,product_name,product_brand,quantity,product_price,image_url) VALUE (2,"43 Piece dinner Set","CoolKidz",10,12.95,"/43_Piece_dinner_Set.png");
INSERT INTO products (id,product_name,product_brand,quantity,product_price,image_url) VALUE (1,"Balance Training Bicycle","SportsFun",5,119.95,"/Balance_Training_Bicycle.png");

CREATE TABLE orders (
    id BIGINT AUTO_INCREMENT,
    shipping_method varchar(50),
    total_price double,
    transaction_id varchar(255) DEFAULT '',
    status ENUM('created','completed','cancle','fail') DEFAULT 'created',
    authorized timestamp DEFAULT current_timestamp,
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
) CHARACTER SET utf8 COLLATE utf8_general_ci;

INSERT INTO orders (id, shipping_method) VALUE (1,"Kerry");
INSERT INTO orders (id, total_price) 
VALUE (8004359103,14.59);

CREATE TABLE order_product (
    order_id BIGINT,
    product_id BIGINT,
    quantity int,
    product_price double
) CHARACTER SET utf8 COLLATE utf8_general_ci;

INSERT INTO order_product (order_id, product_id, quantity, product_price) 
VALUE (1, 2, 10, 1199.5);
INSERT INTO order_product (order_id, product_id, quantity, product_price) 
VALUE (1, 1, 10, 129.5);

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
) CHARACTER SET utf8 COLLATE utf8_general_ci;
INSERT INTO shipping (id,order_id,address,sub_district,district,province,zip_code,recipient,phone_number) 
VALUE (1,1,"405/37 ถ.มหิดล", "ท่าศาลา", "เมือง", "เชียงใหม่", "50000", "ณัฐญา ชุติบุตร", "0970809292");