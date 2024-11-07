package rpc

import (
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

type BizService struct {
	BizModel model.BizModel
	pb.UnimplementedBizServiceServer
}

func NewBizService(model model.BizModel) *BizService {
	return &BizService{
		BizModel: model,
	}
}
