param(
    [switch]$SkipInstall
)

$ErrorActionPreference = "Stop"

Write-Host "==> Running Benchmarks unit tests"

if (-not $SkipInstall) {
    Write-Host "==> Installing Python test dependencies (pytest, httpx)"
    python -m pip install pytest httpx
    if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }
}

Write-Host "==> Running FastAPI tests (pytest)"
pytest -q benchmarks/fastapi_server/test_app.py
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "==> Running Gin tests (go test)"
go -C benchmarks/gin_server test ./...
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "==> All Benchmarks unit tests passed"
exit 0

