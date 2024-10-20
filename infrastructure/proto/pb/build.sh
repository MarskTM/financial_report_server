# Run in root directory (Backend folder)
python -m grpc_tools.protoc -I./infrastructure/proto/pb --python_out=./infrastructure/proto/code/python --grpc_python_out=./infrastructure/proto/code/python test.proto

