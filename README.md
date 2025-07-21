# Go API Students

📚 REST API для управления списком студентов на языке Go с использованием PostgreSQL.

## 🚀 Функционал

- Получить список студентов (`GET /students`)
- Добавить студента (`POST /students`)
- Обновить данные студента (`PUT /students/{id}`)
- Удалить студента (`DELETE /students/{id}`)

## 🛠️ Стек

- Язык: Go
- БД: PostgreSQL
- Роутинг: [Chi](https://github.com/go-chi/chi)
- Работа с БД: [pgx](https://github.com/jackc/pgx)
- .env: [godotenv](https://github.com/joho/godotenv)

## 📦 Установка и запуск

1.Клонировать репозиторий 

```bash
git clone https://github.com/yourname/go-api-students.git
cd go-api-students


2. Создать .env файл

DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=testdb
DB_HOST=localhost
DB_PORT=5432

3. Поднять PostgreSQL через Docker(если не установлен)

docker run --name pg-go-api \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=testdb \
  -p 5432:5432 \
  -d postgres

4. Создать таблицу 

CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  age INT,
  enrolled BOOLEAN DEFAULT true
);

5. Запустить сервер 

go run ./cmd/api




