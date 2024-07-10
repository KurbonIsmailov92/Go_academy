--Задача 1
CREATE DATABASE computer_firm;
CREATE TABLE product
(
    maker VARCHAR(10) NOT NULL,
    model VARCHAR(50) NOT NULL,
    type  VARCHAR(50) NOT NULL,
    PRIMARY KEY (model)
);
CREATE TABLE pc
(
    code  SERIAL PRIMARY KEY,
    model VARCHAR(50) NOT NULL,
    speed SMALLINT    NOT NULL,
    ram   SMALLINT    NOT NULL,
    hd    DECIMAL     NOT NULL,
    cd    VARCHAR(10) NOT NULL,
    price DECIMAL,
    FOREIGN KEY (model) REFERENCES product (model)
);
CREATE TABLE laptop
(
    code   SERIAL PRIMARY KEY,
    model  VARCHAR(50) NOT NULL,
    speed  SMALLINT    NOT NULL,
    ram    SMALLINT    NOT NULL,
    hd     DECIMAL     NOT NULL,
    price  DECIMAL,
    screen SMALLINT    NOT NULL,
    FOREIGN KEY (model) REFERENCES product (model)
);
CREATE TABLE printer
(
    code  SERIAL PRIMARY KEY,
    model VARCHAR(50)           NOT NULL,
    color BOOLEAN DEFAULT FALSE NOT NULL,
    type  VARCHAR(10)           NOT NULL,
    price DECIMAL,
    FOREIGN KEY (model) REFERENCES product (model)
);
------------------------------------------------------------------------------------------------------------------------
--Задача 2
INSERT INTO product (maker, model, type)
VALUES ('B', '1121', 'PC'),
       ('A', '1232', 'PC'),
       ('A', '1233', 'PC'),
       ('E', '1260', 'PC'),
       ('A', '1276', 'Printer'),
       ('D', '1288', 'Printer'),
       ('A', '1298', 'Laptop'),
       ('C', '1321', 'Laptop'),
       ('A', '1401', 'Printer'),
       ('A', '1408', 'Printer'),
       ('D', '1433', 'Printer'),
       ('E', '1434', 'Printer'),
       ('B', '1750', 'Laptop'),
       ('A', '1752', 'Laptop'),
       ('E', '2112', 'PC'),
       ('E', '2113', 'PC');
INSERT INTO pc (model, speed, ram, hd, cd, price)
VALUES ('1232', '500', '64', '5.0', '12x', 600.00),
       ('1121', '750', '128', '14.0', '40x', 850.00),
       ('1233', '500', '64', '5.0', '12x', 600.00),
       ('1121', '600', '128', '14.0', '40x', 850.00),
       ('1121', '600', '128', '8.0', '40x', 850.00),
       ('1233', '750', '128', '20.0', '50x', 950.00),
       ('1232', '500', '32', '10.0', '12x', 400.00),
       ('1232', '450', '64', '8.0', '24x', 350.00),
       ('1232', '450', '32', '10.0', '24x', 350.00),
       ('1260', '500', '32', '10.0', '12x', 350.00),
       ('1233', '900', '128', '40.0', '40x', 980.00),
       ('1233', '800', '128', '20.0', '50x', 970.00);
INSERT INTO laptop (model, speed, ram, hd, price, screen)
VALUES ('1298', '350', '32', '4', '700.00', '11'),
       ('1321', '500', '64', '8', '970.00', '12'),
       ('1750', '750', '128', '12', '1200.00', '14'),
       ('1298', '600', '64', '10', '1050.00', '15'),
       ('1752', '750', '128', '10', '1150.00', '14'),
       ('1298', '450', '64', '10', '950.00', '12');
INSERT INTO printer (model, color, type, price)
VALUES ('1288', FALSE, 'Laser', '400.00'),
       ('1408', FALSE, 'Matrix', '270.00'),
       ('1401', FALSE, 'Matrix', '150.00'),
       ('1434', TRUE, 'Jet', '290.00'),
       ('1433', TRUE, 'Jet', '270.00'),
       ('1276', FALSE, 'Laser', '400.00');
------------------------------------------------------------------------------------------------------------------------
--Задача 3
SELECT l.model,
       l.ram,
       l.screen
FROM laptop l
WHERE l.price > '1000';
------------------------------------------------------------------------------------------------------------------------
--Задача 4
SELECT *
FROM printer p
WHERE p.color = TRUE;
------------------------------------------------------------------------------------------------------------------------
--Задача 5
SELECT model,
       speed,
       hd
FROM pc
WHERE cd in ('12x', '24x')
  AND price < 600;
------------------------------------------------------------------------------------------------------------------------