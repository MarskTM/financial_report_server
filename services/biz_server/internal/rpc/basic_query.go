package rpc

import (
	"context"

	"github.com/MarskTM/financial_report_server/infrastructure/database/dao"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
)

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