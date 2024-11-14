package main

import (
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/biz_server/internal"
)

func main() {
	bizServer := internal.NewBizService()
	system.RunAppService(bizServer)
}
