package rpc

import (
	"context"
	"fmt"

	"github.com/MarskTM/financial_report_server/infrastructure/database/dao"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/golang/glog"
)

type BizService struct {
	BizModel model.BizModel
	pb.UnimplementedBizServiceServer
}

func (s *BizService) Authenticate(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {
	if req.Username == "" {
		err := fmt.Errorf("Error not found username for Authenticate!")
		glog.V(1).Info(err)
		return nil, err
	}

	return &pb.AuthResponse{
		Session:   1,
		UserId:    35023,
		Usernames: "To Kim Manh",
		Roles:     []string{"admin"},
	}, nil
}

func (s *BizService) BasicQuery(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {
	commonDAO := dao.NewCommonDAO(s.BizModel.DB.Postgre)

	query := "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"

	ids := []int32{1, 2}
	_, err := commonDAO.BasicQuery(query, ids)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *BizService) AdvancedFilter(ctx context.Context, req *pb.Credentials) (*pb.AuthResponse, error) {
	return nil, nil
}

func NewBizService(model model.BizModel) *BizService {
	return &BizService{
		BizModel: model,
	}
}
