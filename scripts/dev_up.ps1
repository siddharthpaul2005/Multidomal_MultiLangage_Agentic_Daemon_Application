Write-Host "[1/2] Starting manager..."
Start-Process powershell -ArgumentList "cd manager; go run ."

Write-Host "[2/2] Starting example agent..."
Start-Process powershell -ArgumentList "cd agents/example_agent; python agent.py"
