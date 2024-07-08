--Задача 1. Таблица "Заказы"

CREATE TABLE sales
(
    id         SERIAL PRIMARY KEY,
    client_id  INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    quantity   INTEGER   DEFAULT 0,
    ordered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO sales (client_id, product_id, quantity)
VALUES (12, 23, 2);

SELECT *
FROM sales
WHERE client_id = '2';

UPDATE sales
SET quantity = quantity + 12
WHERE id = 1;

DELETE
FROM sales
WHERE id = 1;

--Задача 2. Таблица "События"

CREATE TABLE actions
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(30),
    date                TIMESTAMP,
    place               VARCHAR(30),
    participiants_count INTEGER
);

INSERT INTO actions (name, date, place, participiants_count)
VALUES ('Проведение операции', CURRENT_TIMESTAMP, 'Касса', 2);

SELECT *
FROM actions
WHERE place = 'Касса';

UPDATE actions
SET participiants_count = participiants_count + 1
WHERE id = '12';

DELETE
FROM actions
WHERE id = '2'