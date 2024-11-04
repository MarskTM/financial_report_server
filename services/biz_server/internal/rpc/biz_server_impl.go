package rpc

import (
	"context"

	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

type BizService struct {
	BizModel model.BizModel
	pb.UnimplementedBizServiceServer
}

func (s *BizService) Authenticate(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {
	
	return nil, nil
}

func NewBizService(model model.BizModel) *BizService {
	return &BizService{
		BizModel: model,
	}
}
