CREATE TABLE user (
    id int,
    name varchar(255)
);

INSERT INTO user VALUE (1,"sckshuhari");


CREATE TABLE products (
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

INSERT INTO products (id,product_name,product_brand,quantity,product_price,image_url) VALUE (2,"43 Piece dinner Set","CoolKidz",10,12.95,"/43_Piece_dinner_Set.png");

CREATE TABLE orders (
    id int AUTO_INCREMENT,
    total_price double,
    transaction_id varchar(255),
    completed smallint(1) DEFAULT 0,
    authorized timestamp,
    updated timestamp DEFAULT current_timestamp,
    created timestamp DEFAULT current_timestamp ON UPDATE current_timestamp,
    PRIMARY KEY (id)
);