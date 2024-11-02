package main

import (
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/gateway/internal"
)

func main() {
	gatewayService := internal.NewGatewayService()
	system.RunAppService(gatewayService)
}
