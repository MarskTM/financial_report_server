package client

import (
	"log"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"google.golang.org/grpc"
)

func NewDocumentClient() pb.DocumentClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Không thể kết nối đến gRPC server: %v", err)
	}
	defer conn.Close()

	// Tạo client từ kết nối
	DocumentClient := pb.NewDocumentClient(conn)

	return DocumentClient
}
