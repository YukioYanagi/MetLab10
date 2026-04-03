from fastapi import FastAPI
from fastapi.responses import Response

# Отключаем OpenAPI/Docs, чтобы не влиять на измерение производительности.
app = FastAPI(openapi_url=None, docs_url=None, redoc_url=None)

STATIC_JSON = b'{"ok":true}'


@app.get("/ping")
def ping() -> Response:
    # Возвращаем заранее подготовленные байты, чтобы минимизировать накладные расходы
    # на сериализацию/валидацию.
    return Response(content=STATIC_JSON, media_type="application/json")

