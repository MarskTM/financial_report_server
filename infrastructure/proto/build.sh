# Run in root directory (Backend folder)
python -m grpc_tools.protoc -I./infrastructure/proto/pb --python_out=./infrastructure/proto/code/python --grpc_python_out=./infrastructure/proto/code/python test.proto

# For golang gencode
protoc --go_out=. --go-grpc_out=. services_analist.proto
protoc --go_out=. --go-grpc_out=. services_document.proto
protoc --go_out=. --go-grpc_out=. services.proto