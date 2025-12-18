Get-Process | Where-Object { .ProcessName -like "*go*" -or .ProcessName -like "python" } | Stop-Process -Force
Write-Host "Cleaned."

<#
Explanation:

- Get-Process
  Enumerates all processes currently running on the system. This returns objects that include process metadata such as `ProcessName`, `Id`, `Path`, etc.

- Where-Object { .ProcessName -like "*go*" -or .ProcessName -like "python" }
  Filters the list of processes to those whose `ProcessName` matches the given patterns.
  - `-like` performs wildcard pattern matching (supports `*`), and is case-insensitive by default in PowerShell.
  - The patterns used here match process names that contain the substring "go" (e.g., `go`, `gofmt`, etc.) or match "python".
  Note: This can match multiple processes and may be broader than intended; consider using explicit names (e.g., `go.exe`, `python.exe`) or adding additional checks (process `Id`, `Path`) for safety.

- Stop-Process -Force
  Terminates the filtered processes immediately. The `-Force` flag forces termination without prompting and may not allow processes to shut down cleanly.

- Write-Host "Cleaned."
  Prints a simple confirmation message to the console after the stop attempt.

Safety notes and recommendations:
- This script can terminate any Go or Python process (including unrelated ones). Run with caution.
- For safer cleanup, preview targets before killing, for example:
    Get-Process | Where-Object { $_.ProcessName -like "*go*" -or $_.ProcessName -like "python" } | Select-Object Id, ProcessName, Path
  and then selectively stop specific `Id` values.
- If you only want to stop processes started by `dev_up.ps1`, consider tracking and storing the PIDs when starting them and only stopping those PIDs.
- On non-Windows systems, behavior differs (process names and commands may vary).
#>
