Долгополов Андрей Дмитриевич, группа 221331, Веб-разработка: FastAPI (Python) vs Gin (Go), Лабораторная работа №10, Вариант 6

## Benchmarks

Примеры сравнительных бенчмарков под нагрузкой для `FastAPI` и `Gin` находятся в папке `benchmarks/`.

См. `benchmarks/README.md`.

## API Docs

Пример с документацией API (Swagger для FastAPI и OpenAPI endpoint для Gin) находится в папке `api_docs/`.

См. `api_docs/README.md`.

## Go Logging Middleware

Пример с middleware для логирования запросов в Go (Gin) находится в папке `go_logging/`.

См. `go_logging/README.md`.

## Repository checks

Добавлены автоматические проверки в CI (`.github/workflows/repo-check.yml`) и локальный скрипт:

- `python scripts/repo_check.py`

Скрипт проверяет:
- размер tracked-файлов < 1 MB;
- отсутствие запрещенных артефактов в git (`.exe`, `.dll`, `.so`, `.venv`, `target`, `__pycache__`);
- наличие `go.mod` для Go-проектов;
- Rust edition (`2021` или `2024`);
- `pyproject.toml` для maturin-проектов;
- заполненность `PROMPT_LOG.md`.

## Unit-test checks

Добавлены локальные скрипты для проверки работоспособности (`.scripts\...`), которые можно проверить с помощью команды в терминале (`powershell -ExecutionPolicy Bypass -File .\scripts\...`):

- `.\scripts\run_benchmarks_tests.ps1`
- `.\scripts\run_api_docs_tests.ps1`
- `.\scripts\run_go_logging_tests.ps1`
