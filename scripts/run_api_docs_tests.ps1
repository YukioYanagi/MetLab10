param(
    [switch]$SkipInstall
)

$ErrorActionPreference = "Stop"

Write-Host "==> Running API Docs unit tests"

if (-not $SkipInstall) {
    Write-Host "==> Installing Python test dependencies (pytest, httpx)"
    python -m pip install pytest httpx
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
}

Write-Host "==> Running FastAPI API Docs tests (pytest)"
pytest -q api_docs/fastapi_server/test_app.py
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "==> Running Gin API Docs tests (go test)"
go -C api_docs/gin_server test ./...
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "==> All API Docs unit tests passed"
exit 0

