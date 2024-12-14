# Backend for Parser

Backend API сервис для работы с парсером. Позволяет управлять пользователями и образцами данных, а также взаимодействовать с парсером через специальный эндпоинт.

## Установка

1. **Склонируйте репозиторий:**
   ```bash
   git clone https://github.com/mirdan2207/backend-for-parser.git
   cd backend-for-parser
   ```

2. **Установите зависимости:**
   Убедитесь, что у вас установлен Go (версии 1.20 или выше).
   ```bash
   go mod download
   ```

3. **Настройте базу данных:**
   Создайте PostgreSQL базу данных по скрипту db.sql

4. **Добавьте .env файл**
    Вот пример:
    ```json
    DB_CONN_STRING=...
    SERVER_PORT=...
    ```

5. **Запустите сервер:**
   ```bash
   go run cmd/main.go
   ```

## Использование

API предоставляет следующие эндпоинты:

### **Пользователи**
- **Регистрация пользователя**
  ```http
  POST /users/register
  ```
  **Параметры запроса:**
  ```json
  {
      "login": "user_login",
      "password": "user_password"
  }
  ```
  **Oтвет:**
  ```json
  {
    "user_id": int,
  }
  ```

- **Авторизация пользователя**
  ```http
  POST /users/login
  ```
  **Параметры запроса:**
  ```json
  {
      "login": "user_login",
      "password": "user_password"
  }
  ```
  **Oтвет:**
  ```json
  {
    "user_id": int,
  }
  ```

### **Образцы данных**
- **Создание образца**
  ```http
  POST /samples
  ```
  **Тело запроса:**
  ```json
  {
      "name": "sample_name",
      "data": "sample_data"
  }
  ```

- **Получение всех образцов**
  ```http
  GET /samples
  ```

- **Получение образца по ID**
  ```http
  GET /samples/{id}
  ```

- **Редактирование образца**
  ```http
  PUT /samples/{id}
  ```
  **Тело запроса:**
  ```json
  {
      "name": "updated_name",
      "data": "updated_data"
  }
  ```

- **Удаление образца**
  ```http
  DELETE /samples/{id}
  ```

### **Парсер**
- **Запуск парсера**
  ```http
  POST /execute
  ```
  **Тело запроса:**
  ```json
  {
      "sample_id": "id_of_sample"
  }
  ```
  **Ответ:** Результаты обработки от парсера.

## Технологии
- Язык программирования: Go
- Фреймворк для роутинга: `gorilla/mux`
- База данных: PostgreSQL

