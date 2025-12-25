# ============================================
# Proto generator (Windows / PowerShell)
# Generates:
#  - Go gRPC code for Manager
#  - Python gRPC code for Agents
# ============================================

$ErrorActionPreference = "Stop"

# Directories
$PROTO_DIR  = "proto"
$PY_OUT_DIR = "proto"        # Python code lives alongside proto (importable)
$GO_OUT_DIR = "manager"      # Go code generated inside manager module

Write-Host "[Proto] Generating protobuf code..."

# ------------------------------------------------
# Sanity checks
# ------------------------------------------------

if (-not (Test-Path $PROTO_DIR)) {
    throw "Proto directory '$PROTO_DIR' not found."
}

if (-not (Get-Command protoc -ErrorAction SilentlyContinue)) {
    throw "protoc not found on PATH. Install Protocol Buffers compiler."
}

# Ensure Go plugins are available
if ($env:GOPATH) {
    $env:Path = "$($env:GOPATH)\bin;$($env:Path)"
}

if (-not (Get-Command protoc-gen-go -ErrorAction SilentlyContinue)) {
    throw "protoc-gen-go not found. Install with:
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
}

if (-not (Get-Command protoc-gen-go-grpc -ErrorAction SilentlyContinue)) {
    throw "protoc-gen-go-grpc not found. Install with:
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
}

# ------------------------------------------------
# Go code generation (Manager)
# ------------------------------------------------

Write-Host "[Proto] Generating Go code..."

protoc `
  -I $PROTO_DIR `
  --go_out=$GO_OUT_DIR `
  --go-grpc_out=$GO_OUT_DIR `
  --go_opt=module=hyperagent/manager `
  --go-grpc_opt=module=hyperagent/manager `
  $PROTO_DIR/*.proto

# ------------------------------------------------
# Python code generation (Agents)
# ------------------------------------------------

Write-Host "[Proto] Generating Python code..."

python -m grpc_tools.protoc `
  -I proto `
  --python_out=. `
  --grpc_python_out=. `
  proto/*.proto

Write-Host "[Proto] ✅ Generation complete."
