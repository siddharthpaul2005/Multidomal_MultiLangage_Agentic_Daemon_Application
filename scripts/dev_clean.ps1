Get-Process | Where-Object { .ProcessName -like "*go*" -or .ProcessName -like "python" } | Stop-Process -Force
Write-Host "Cleaned."
