package config

import (
	"github.com/MarskTM/financial_report_server/infrastructure/database/config"
)

// ----------------------------------------------------------------------------
// config models
type ConfigService struct {
	Gateway  GatewayConfig     `json:"gateway"`
	DBConfig config.PostConfig `json:"dbconfig"`
}

type GatewayConfig struct {
	Addr           string `json:"addr"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	MaxHeaderBytes int    `json:"max_header_bytes"`
}
