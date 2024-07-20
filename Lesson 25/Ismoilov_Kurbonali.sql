--Задача 1
SELECT COUNT(e.id)                                          AS "Общее количество сотрудников",
       AVG(e.age)                                           AS "Средний возраст сотрудников",
       MAX(e.salary)                                        AS "Максимальная зарплата",
       MIN(e.salary)                                        AS "минимальная зарплата",
       SUM(e.salary)                                        AS "Общая сумма всех зарплат",
       (SELECT COUNT(id)
        FROM employees
        WHERE age > 45)                                     AS "Кол-во сотрудников старше 45 лет",
       (SELECT COUNT(id)
        FROM employees
        WHERE salary > (SELECT AVG(salary) FROM employees)) AS "Кол-во сотрудников с з/п выше среднего"
FROM employees e;

--Задача 2
SELECT e.department  AS "Должность",
       SUM(e.salary) AS "Сумма зарплат"
FROM employees e
GROUP BY e.department
HAVING AVG(e.salary) > 50000;


