ALTER TABLE employees
    ADD COLUMN region VARCHAR (30)
;
SELECT position,
       AVG(salary)
FROM employees
GROUP BY position;

INSERT INTO employees (name, department, salary, age, position, region)
VALUES 	   ('John Doe', 'Marketing', 50000.00, 35, 'Back-End','US'),
             ('Jane Smith', 'Marketing', 55000.00, 42, 'Designer','UK'),
             ('Bob Johnson', 'IT', 60000.00, 28, 'Back-End','US'),
             ('Sara Lee', 'IT', 65000.00, 52, 'Front-End','UK'),
             ('Mike Williams', 'HR', 45000.00, 31, 'Q&A','US'),
             ('Emily Davis', 'HR', 48000.00, 27, 'Q&A','China'),
             ('David Brown', 'Finance', 70000.00, 47, 'Back-End','China'),
             ('Samantha Wilson', 'Finance', 75000.00, 55, 'Back-End','US'),
             ('Tom Garcia', 'Marketing', 52000.00, 29, 'Back-End','Russia'),
             ('Olivia Hernandez', 'IT', 62000.00, 38, 'Front-End','Russia');


SELECT position,
       SUM(salary)
FROM employees
GROUP BY position
HAVING AVG(salary) > 50000;

SELECT name,
       salary
FROM employees
WHERE salary > (SELECT AVG(salary) FROM employees);