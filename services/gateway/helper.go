package main

import (
	"github.com/MarskTM/financial_report_server/infrastructure/server"
	"github.com/MarskTM/financial_report_server/services/gateway/internal"
)

func main() {
	gatewayService := internal.NewGatewayService()
	go server.RunAppService(gatewayService)
}
