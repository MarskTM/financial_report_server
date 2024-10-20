package internal

import (
	"net/http"
	"os"

	"github.com/MarskTM/financial_report_server/baselib/server"
	"github.com/MarskTM/financial_report_server/services/gateway/internal/config"
	"github.com/MarskTM/financial_report_server/services/gateway/internal/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type GatewayService struct {
	apiServer *http.Server
	server    *grpc.Server
	impl      *rpc.ServiceImpl
	config    *config.ConfigService
}

// Constructor creates a new GatewayServer
func NewGatewayService() server.ServicesInterface {
	return &GatewayService{}
}

// -----------------------------------------------------------------------------
// service interface
func (s *GatewayService) Install() error {

	glog.V(3).Infof("server::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	err := config.LoadConfig(s.config)
	if err != nil {
		glog.V(1).Infof("server::Initialize - Error: %+v", err)
		return err
	}

	// 2. Install DAO

	// 3. Install gRPC client

	// 4. Install gRPC server

	return nil
}

func (s *GatewayService) Start() {
	go func() {
        err := s.apiServer.ListenAndServe()
        if err!= nil {
            glog.Fatalf("apiServer::ListenAndServe - Error: %+v", err)
        }
    }()
}

func (s *GatewayService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
