CREATE TABLE user (
    ID int,
    Name varchar(255)
);

INSERT INTO user VALUE (1,"sckshuhari");


CREATE TABLE product (
    id int AUTO_INCREMENT,
    product_name varchar(255),
    product_brand varchar(255),
    quantity int,
    product_price double,
    image_url varchar(255),
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);