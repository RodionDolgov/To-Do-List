Это простое REST API для управления задачами (TODO-лист), разработанное на Go с использованием Fiber и PostgreSQL.

## Запуск

1. Убедитесь, что у вас установлены Go и PostgreSQL.
2. Создайте базу данных `ToDoList` в PostgreSQL.
3. Создайте таблицу `tasks` с помощью SQL-запроса из файла `migrations/001_create_tasks_table.sql`.
4. Запустите приложение:
   ```bash
   go run main.go