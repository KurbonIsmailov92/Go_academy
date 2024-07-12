CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50)            NOT NULL,
    login      VARCHAR(20)            NOT NULL,
    password   VARCHAR(30)            NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active  BOOLEAN   DEFAULT TRUE NOT NULL
);

CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER REFERENCES users (id) NOT NULL,
    title       VARCHAR                       NOT NULL,
    description TEXT,
    is_done     BOOLEAN   DEFAULT FALSE       NOT NULL,
    is_deleted  BOOLEAN   DEFAULT FALSE       NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP,
    done_at     TIMESTAMP
);

SELECT u.name    AS "Имя пользователя",
       t.title   AS "Название задачи",
       t.is_done AS "Статус задачи"
FROM tasks t,
     users u
WHERE t.user_id = u.id
ORDER BY u.name;
