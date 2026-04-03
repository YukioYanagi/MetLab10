from fastapi import FastAPI

app = FastAPI(
    title="MetLab10 FastAPI Docs Demo",
    version="1.0.0",
    description="Demo API with Swagger UI and OpenAPI schema.",
)


@app.get("/ping", tags=["health"])
def ping() -> dict[str, bool]:
    return {"ok": True}

