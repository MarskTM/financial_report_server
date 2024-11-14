//  File này định nghĩa cho các công nghệ đươc truyền xuống tầng internal để xử lý dũ liệu.
//  Các mô hình sẽ được khai báo trên cùng cúa các file helper, dùng khi chạy service.

package model

import (
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/go-chi/jwtauth"
)

type BizModel struct {
	Config env.BizServerConfig
	DB     database.ManagerDAO

	DocsClient pb.DocumentClient
	// AnalystClient pb.AnalystClient
}

type GatewayModel struct {
	Config     env.GatewayConfig
	ManagerDao database.ManagerDAO
	DecodeAuth *jwtauth.JWTAuth
	EncodeAuth *jwtauth.JWTAuth

	DocsClient pb.DocumentClient
	BizClient  pb.BizServiceClient
}

type DocumentModel struct {
	Config env.DocumentConfig
	DB     database.ManagerDAO

	BizClient pb.BizServiceClient
}
