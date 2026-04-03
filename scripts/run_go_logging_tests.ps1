$ErrorActionPreference = "Stop"

Write-Host "==> Running Go Logging unit tests"
go -C go_logging/server test ./...
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "==> Go Logging unit tests passed"
exit 0

