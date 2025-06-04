#  Todo API на Go + PostgreSQL + Gin

Полноценное API для задач (ToDo), написанное на Go. Проект реализует CRUD-операции, подключается к PostgreSQL.

---

##  Возможности

 ● Получить список всех задач
 ● Получить одну задачу по ID
 ● Создать задачу
 ● Обновить задачу
 ● Удалить задачу
 ● Валидация входных данных
 ● Работа с БД

---

##  Стек технологий

● Язык: Go 1.24.3
● Фреймворк: Gin
● База данных: PostgreSQL 17.5 x86-64 windows
● Менеджер переменных окружения: godotenv

---

## Установка

```bash
git clone https://github.com/твоя-ссылка.git
cd todoAPI
go mod tidy
```
## Запуск
1) Настроить .env файл под себя
2) Запустить PostgresSQL и создать БД(
    CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL
);
)
3) Запустить сервер -> go run main.go
4) Сервер будет доступен по адресу http://localhost:8080
5) Протестируй через POSTMAN или CURL
Пример запроса:
POST http://localhost:8080/todos
Content-Type: application/json
{
  "title": "Сходить в зал",
  "completed": false
}
