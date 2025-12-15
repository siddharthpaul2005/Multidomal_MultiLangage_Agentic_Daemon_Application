# Correct proto generator for Windows

$PROTO_DIR = "proto"
$OUT_DIR   = "generated"

# Ensure output folder exists
if (!(Test-Path $OUT_DIR)) {
    New-Item -ItemType Directory -Force -Path $OUT_DIR | Out-Null
}

# Run protoc via the Python grpc_tools wrapper (this provides the grpc_python plugin)
python -m grpc_tools.protoc `
    -I $PROTO_DIR `
    --python_out=$OUT_DIR `
    --grpc_python_out=$OUT_DIR `
    proto/common.proto `
    proto/agent.proto `
    proto/manager.proto