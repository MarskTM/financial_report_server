package rpc

import (
	"context"

	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

type BizService struct {

	pb.UnimplementedBizServiceServer
}

func (s *BizService) Authenticate(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {



	return nil, nil
}


func NewBizService() *BizService {
	return &BizService{

	}
}
