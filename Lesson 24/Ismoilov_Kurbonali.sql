SELECT DISTINCT author
FROM books
WHERE publisher_id > 3;

SELECT *
FROM books
ORDER BY price
LIMIT 3;

SELECT *
FROM books
ORDER BY pubplished_at
OFFSET 2 LIMIT 3;

SELECT b.title,
       b.author,
       b.price
FROM books b,
     publishers p
WHERE b.publisher_id = p.id
  AND p.name in ('HarperCollins', 'Simon & Schuster');

SELECT *
FROM books
WHERE pubplished_at BETWEEN '2021-01-01' AND '2022-12-31'
  AND price > 20.00;

SELECT *
FROM books
WHERE author like '%J%';

SELECT b.title,
       p.name,
       b.pubplished_at
FROM books b,
     publishers p
WHERE p.address = 'New York';