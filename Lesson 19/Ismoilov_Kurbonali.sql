---------------------------------------------------
--Задача 4.
CREATE TABLE books
(
    id               SERIAL PRIMARY KEY NOT NULL,
    title            VARCHAR(30),
    author           VARCHAR(30),
    publication_date TIMESTAMP,
    genre            VARCHAR(20),
    is_available     BOOLEAN DEFAULT TRUE
);
---------------------------------------------------

---------------------------------------------------
--Задач 5.
CREATE TABLE students_academic_performances
(
    id         SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(20),
    last_name  VARCHAR(20),
    email      VARCHAR UNIQUE,
    course     INTEGER DEFAULT '1',
    grade      INTEGER
);
---------------------------------------------------