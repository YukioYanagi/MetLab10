# Benchmarks (FastAPI vs Gin)

Цель: сравнить **латентность/пропускную способность** под нагрузкой для одинакового endpoint.

Сравнивается обработчик `/ping`, который возвращает статический JSON-ответ.

## Быстрый старт

### FastAPI

1. Установите зависимости:
   - `pip install -r fastapi_server/requirements.txt` (в папке `benchmarks`)
2. Запустите сервер (порт `8000`):
   - `cd benchmarks`
   - `uvicorn fastapi_server.app:app --host 127.0.0.1 --port 8000 --workers 1 --log-level warning --no-access-log`

### Gin

1. Запустите сервер (порт `8001`):
   - `cd benchmarks/gin_server`
   - `go run main.go`

## Нагрузочное тестирование

Рекомендуется запускать **в двух отдельных терминалах**.

### `wrk` (предпочтительно для latency/throughput)

FastAPI:
```bash
wrk -t4 -c200 -d30s --latency http://127.0.0.1:8000/ping
```

Gin:
```bash
wrk -t4 -c200 -d30s --latency http://127.0.0.1:8001/ping
```

Чтобы найти точку насыщения, повторяйте с разными `-c` (например, `50/100/200/400`).

### `ab` (проще, но менее информативно)

FastAPI:
```bash
ab -n 200000 -c 200 -k http://127.0.0.1:8000/ping
```

Gin:
```bash
ab -n 200000 -c 200 -k http://127.0.0.1:8001/ping
```

## Важно для честного сравнения

- Сравнивайте одинаковые настройки: `workers 1` у FastAPI и одиночный Gin-процесс.
- Уберите лишнее: логирование/трейсинг/тяжелые middleware.
- На Windows часто разная настройка сети/антивирус/фоновые процессы влияет сильнее, чем сам фреймворк.

## Unit-тесты (работоспособность endpoint)

Тесты добавлены для обоих серверов и запускаются в CI автоматически.

Локальный запуск:

- Go (Gin):
  - `cd benchmarks/gin_server`
  - `go test ./...`

- Python (FastAPI):
  - `python -m pip install pytest httpx`
  - `pytest -q`

Единый скрипт из корня репозитория:

- `powershell -ExecutionPolicy Bypass -File .\scripts\run_benchmarks_tests.ps1`
- быстрый повторный запуск без установки зависимостей:
  - `powershell -ExecutionPolicy Bypass -File .\scripts\run_benchmarks_tests.ps1 -SkipInstall`
