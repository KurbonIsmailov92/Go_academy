CREATE TABLE customers
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(30),
    last_name  VARCHAR(30),
    email      VARCHAR UNIQUE,
    age        INTEGER
);

INSERT INTO customers (first_name, last_name, email, age)
VALUES ('Vasya', 'Pupkin', 'vasya@mail.com', 23),
       ('Petya', 'Petrov', 'pipetya@mail.com', 25);

ALTER TABLE customers
    ADD email VARCHAR UNIQUE;

UPDATE customers
SET email = 'vasya@mail.com'
WHERE id = 1;

ALTER TABLE customers
    DROP COLUMN email;


ALTER TABLE customers
    ALTER COLUMN phone TYPE VARCHAR(30);

ALTER TABLE customers
    ALTER COLUMN first_name
        SET NOT NULl;

ALTER TABLE customers
    ALTER COLUMN first_name
        DROP NOT NULL;

ALTER TABLE clients
    RENAME TO customers;

ALTER TABLE customers
    RENAME COLUMN email TO post;

CREATE TABLE otps
(
    otp_code     VARCHAR(4),
    phone_number VARCHAR(22),
    sent_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (otp_code, phone_number)
);

INSERT INTO otps (otp_code, phone_number)
VALUES ('5565', '992333333333');





