package main

import (
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/document/internal"
)

func main() {
	gatewayService := internal.NewDocumentService()
	system.RunAppService(gatewayService)
}
