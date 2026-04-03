from fastapi.testclient import TestClient

from benchmarks.fastapi_server.app import STATIC_JSON, app


def test_ping_returns_static_json() -> None:
    client = TestClient(app)
    response = client.get("/ping")

    assert response.status_code == 200
    assert response.content == STATIC_JSON
    assert response.headers["content-type"].startswith("application/json")


def test_docs_and_openapi_are_disabled() -> None:
    # Для чистоты бенчмарка служебные endpoint'ы отключены.
    assert app.openapi_url is None
    assert app.docs_url is None
    assert app.redoc_url is None

