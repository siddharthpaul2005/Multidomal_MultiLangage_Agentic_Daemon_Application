Write-Host "[Proto] Generating protobuf code..."

$PROTO_DIR = "proto"
$PY_OUT    = "generated"
$GO_OUT    = "manager"

# Ensure output dirs exist
New-Item -ItemType Directory -Force -Path $PY_OUT | Out-Null

# ------------------------
# Python code generation
# ------------------------
python -m grpc_tools.protoc `
  -I $PROTO_DIR `
  --python_out=$PY_OUT `
  --grpc_python_out=$PY_OUT `
  $PROTO_DIR/common.proto `
  $PROTO_DIR/manager.proto `
  $PROTO_DIR/agent.proto

# ------------------------
# Go code generation
# ------------------------
if ($env:GOPATH) {
    $env:Path = "$($env:GOPATH)\bin;$($env:Path)"
}

protoc `
  -I $PROTO_DIR `
  --go_out=$GO_OUT `
  --go-grpc_out=$GO_OUT `
  --go_opt=module=hyperagent/manager `
  --go-grpc_opt=module=hyperagent/manager `
  $PROTO_DIR/common.proto `
  $PROTO_DIR/manager.proto `
  $PROTO_DIR/agent.proto

Write-Host "[Proto] ✅ Generation complete."
