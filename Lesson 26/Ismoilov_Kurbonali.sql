--Задача 1
SELECT e.id,
       e.name,
       e.department,
       e.age
FROM employees e
WHERE e.age > (SELECT AVG(e2.age)
               FROM employees e2
               WHERE e2.department = e.department);

--Задача 2.
SELECT e.id,
       e.name,
       e.department,
       e.salary
FROM employees e
WHERE e.department = (SELECT department
                      FROM employees
                      ORDER BY salary DESC
                      LIMIT 1);

--Задача 3.
SELECT e.id,
       e.name,
       e.department,
       e.salary
FROM employees e
WHERE e.salary = (SELECT MAX(e2.salary)
                  FROM employees e2
                  WHERE e2.department = e.department);
