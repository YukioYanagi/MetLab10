# Go Logging Middleware Example

Отдельный пример для задачи: добавить middleware для логирования в Go (Gin).

## Что реализовано

- кастомный middleware `requestLoggingMiddleware`;
- логируются метод, путь, HTTP-статус и длительность запроса;
- endpoint `/ping` для проверки работоспособности.

## Запуск

```bash
cd go_logging/server
go run main.go
```

Сервис слушает `127.0.0.1:8201`.

## Unit-тесты

Из корня репозитория:

```bash
go -C go_logging/server test ./...
```

Или единым скриптом:

```bash
powershell -ExecutionPolicy Bypass -File .\scripts\run_go_logging_tests.ps1
```

