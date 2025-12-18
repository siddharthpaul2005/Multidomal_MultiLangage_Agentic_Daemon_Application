Write-Host "[1/2] Starting manager..." //prints message to the console
Start-Process powershell -ArgumentList "cd manager; go run ."

Write-Host "[2/2] Starting example agent..."
Start-Process powershell -ArgumentList "cd agents/example_agent; python agent.py"

//Launches PowerShell and tells it to execute the script file at dev_up.ps1. This is equivalent to opening PowerShell and running dev_up.ps1.
//Starts a new PowerShell process (detached by default). The new process runs two commands in sequence: cd manager (change to the manager directory) and go run . (build & run the manager Go program). Because Start-Process returns immediately, the manager runs in the background.

//Starts another detached PowerShell process. That process changes into example_agent and runs the example Python agent via python agent.py. It also runs in the background.