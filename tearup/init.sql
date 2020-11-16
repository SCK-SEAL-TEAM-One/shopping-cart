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

INSERT INTO products (id,product_name,product_brand,quantity,product_price,image_url) VALUE 
(1,"Balance Training Bicycle","SportsFun",5,119.95,"/Balance_Training_Bicycle.png"),
(2,"43 Piece dinner Set","CoolKidz",10,12.95,"/43_Piece_dinner_Set.png"),
(3,"Horses and Unicorns Set","CoolKidZ",3,24.95,"/Horses_and_Unicorns_Set.png"),
(4,"Hoppity Ball 26 inch","SportsFun",12,29.95,"/Hoppity_Ball_26_inch.png"),
(5,"Sleeping Queens Board Game","CoolKidZ",0,12.95,"/Sleeping_Queens_Board_Game.png"),
(6,"Princess Palace","CoolKidZ",7,24.95,"/Princess_Palace.png"),
(7,"Best Friends Forever Magnetic Dress Up","CoolKidZ",15,24.95,"/Best_Friends_Forever_Magnetic_Dress_Up.png"),
(8,"City Gargage Truck Lego","Lego",0,19.95,"/City_Gargage_Truck_Lego.png"),
(9,"Kettrike Tricycle","SportsFun",0,249.95,"/Kettrike_Tricycle.png"),
(10,"Princess Training Bicycle","SportsFun",0,119.95,"/Princess_Training_Bicycle.png"),
(11,"Earth DVD Game","VideoVroom",2,34.99,"/Earth_DVD_Game.png"),
(12,"Twilight Board Game","GeekToys",7,24.95,"/Twilight_Board_Game.png"),
(13,"Settlers of Catan Board Game","GeekToys",11,44.95,"/Settlers_of_Catan_Board_Game.png"),
(14,"OMG - Gossip Girl Board Game","GeekToys",14,24.95,"/OMG_Gossip_Girl_Board_Game.png"),
(15,"Sailboat","CoolKidZ",10,24.95,"/Sailboat.png"),
(16,"Scrabble","GeekToys",10,19.95,"/Scrabble.png"),
(17,"Start Wars Darth Vader Lego","GeekToys",9,39.95,"/Start_Wars_Darth_Vader_Lego.png"),
(18,"Snoopy Sno-Cone Machine","Modelz",12,24.95,"/Snoopy_Sno_Cone_Machine.png"),
(19,"Gourmet Cupcake Maker","CoolKidZ",15,39.95,"/Gourmet_Cupcake_Maker.png"),
(20,"Creator Beach House Lego","Lego",0,39.95,"/Creator_Beach_House_Lego.png"),
(21,"Jacques the Peacock Play and Grow","CoolKidZ",7,12.95,"/Jacques_the_Peacock_Play_and_Grow.png"),
(22,"Nutbrown Hare","CoolKidZ",0,12.99,"/Nutbrown_Hare.png"),
(23,"Dancing Aligator","CoolKidZ",4,19.95,"/Dancing_Aligator.png"),
(24,"Mashaka the Monkey","BarnyardBlast",5,36.95,"/Mashaka_the_Monkey.png"),
(25,"Sleep Sheep","BarnyardBlast",23,39.00,"/Sleep_Sheep.png"),
(26,"Les Dollie Toffee Apple","CoolKidZ",2,24.95,"/Les_Dollie_Toffee_Apple.png"),
(27,"Sand Play Set","ModelZ",32,24.95,"/Sand_Play_Set.png"),
(28,"Melody Express Musical Train","ModelZ",6,42.95,"/Melody_Express_Musical_Train.png"),
(29,"My First LEGO DUPLO Set","Lego",0,19.95,"/My_First_LEGO_DUPLO_Set.png"),
(30,"Fisher-Price stroller","CoolKidZ",5,25.99,"/Fisher_Price_stroller.png"),
(31,"Mortimer the Moose Play and Grow","CoolKidZ",5,12.95,"/Mortimer_the_Moose_Play_and_Grow.png");

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