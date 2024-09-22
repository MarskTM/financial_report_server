# Run in root directory (Backend folder)
python -m grpc_tools.protoc -I./baselib/proto/pb --python_out=./baselib/proto/code/python --grpc_python_out=./baselib/proto/code/python test.proto