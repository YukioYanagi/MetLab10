# API Docs Example (FastAPI + Gin)

Отдельный пример (вне `benchmarks`) для задачи:
- Swagger-документация для FastAPI;
- OpenAPI endpoint для Gin;
- unit-тесты на работоспособность.

## Структура

- `fastapi_server/` — FastAPI с `/docs` и `/openapi.json` (встроенные возможности).
- `gin_server/` — Gin с endpoint `/openapi.json` (ручной OpenAPI-документ).

## Запуск

### FastAPI

```bash
cd api_docs
pip install -r fastapi_server/requirements.txt
uvicorn fastapi_server.app:app --host 127.0.0.1 --port 8100 --workers 1 --log-level warning --no-access-log
```

Swagger UI:
- [http://127.0.0.1:8100/docs](http://127.0.0.1:8100/docs)

### Gin

```bash
cd api_docs/gin_server
go run main.go
```

OpenAPI JSON:
- [http://127.0.0.1:8101/openapi.json](http://127.0.0.1:8101/openapi.json)

## Unit-тесты

Из корня репозитория:

```bash
python -m pip install pytest httpx
pytest -q api_docs/fastapi_server/test_app.py
go -C api_docs/gin_server test ./...
```

Единый скрипт из корня репозитория:

```bash
powershell -ExecutionPolicy Bypass -File .\scripts\run_api_docs_tests.ps1
```

Быстрый повторный запуск (без установки зависимостей):

```bash
powershell -ExecutionPolicy Bypass -File .\scripts\run_api_docs_tests.ps1 -SkipInstall
```

