--Задача 1
CREATE TABLE inventory
(
    id           SERIAL PRIMARY KEY,
    product_name VARCHAR(30),
    category     VARCHAR(30),
    quantity     INTEGER   DEFAULT 0,
    price        FLOAT     DEFAULT 0.00,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO inventory (product_name, category, quantity, price, last_updated)
VALUES ('Ipad', 'technologies', 12, 980.99, DEFAULT),
       ('Air Max', 'shoes', 45, 80.50, DEFAULT),
       ('Jabulani', 'ball', 120, 5.99, DEFAULT);

SELECT *
FROM inventory AS i;

SELECT *
FROM inventory AS i
WHERE i.category = 'technologies';

SELECT *
FROM inventory AS i
WHERE i.product_name = 'Iphone';


--Задача 2.
CREATE TABLE books
(
    id               SERIAL PRIMARY KEY NOT NULL,
    title            VARCHAR(30),
    author           VARCHAR(30),
    publication_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    genre            VARCHAR(20),
    is_available     BOOLEAN   DEFAULT TRUE
);

INSERT INTO books (title, author, publication_date, genre, is_available)
VALUES ('Война и мир', 'Лев Толстой', DEFAULT, 'Исторический роман', DEFAULT),
       ('Дубровский', 'А.С. Пушкин', DEFAULT, 'Роман', DEFAULT),
       ('Капитанская дочка', 'А.С. Пушкин', DEFAULT, 'Роман', DEFAULT),
       ('Таинственный остров', 'Жюль Верн', DEFAULT, 'Приключения', DEFAULT),
       ('Дети каптина Гранта', 'Жюль Верн', DEFAULT, 'Приключения', DEFAULT);

SELECT *
FROM books;

SELECT *
FROM books
WHERE author = 'Жюль Верн';

SELECT *
FROM books
WHERE genre = 'Роман';

SELECT *
FROM books
WHERE is_available = TRUE;


--Задача 3.
CREATE TABLE students
(
    id         SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(20),
    last_name  VARCHAR(20),
    email      VARCHAR UNIQUE,
    course     INTEGER DEFAULT 1,
    grade      INTEGER DEFAULT 0
);

INSERT INTO students (first_name, last_name, email, course, grade)
VALUES ('Эшмат', 'Тошматов', 'eshtosh@mail.com', 2, 45);

SELECT *
FROM students;

SELECT *
FROM students
WHERE course BETWEEN '1' AND '3';

SELECT *
FROM students
WHERE grade > 80;


