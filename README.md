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
