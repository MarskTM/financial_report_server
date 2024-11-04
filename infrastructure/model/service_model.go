//  File này định nghĩa cho các công nghệ đươc truyền xuống tầng internal để xử lý dũ liệu.
//  Các mô hình sẽ được khai báo trên cùng cúa các file helper, dùng khi chạy service.

package model

import (
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
)

type BizModel struct {
	Config env.BizServerConfig
	DB     database.ManagerDBDao
}

type GatewayModel struct {}

type DocumentModel struct {}