-- Создание таблицы users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,              -- Уникальный идентификатор пользователя
    username VARCHAR(50) NOT NULL,      -- Имя пользователя
    password VARCHAR(100) NOT NULL      -- Пароль пользователя
);

-- Создание таблицы samples
CREATE TABLE samples (
    id SERIAL PRIMARY KEY,              -- Уникальный идентификатор образца
    sample_name VARCHAR(100) NOT NULL,  -- Название образца
    sample_body TEXT NOT NULL,          -- Содержимое образца
    user_id INT NOT NULL,               -- Связь с пользователем
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);