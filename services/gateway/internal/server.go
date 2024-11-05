package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	biz_client "github.com/MarskTM/financial_report_server/services/biz_server/client"
	doc_client "github.com/MarskTM/financial_report_server/services/document/client"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var gatewayModel model.GatewayModel

// ----------------------------------------------------------------
type GatewayService struct {
	apiServer *http.Server
	server    *grpc.Server

	// ----------------------------------------------------------------
	clientConnection map[string]interface{}
}

// Constructor creates a new GatewayServer
func NewGatewayService() system.ServicesInterface {
	return &GatewayService{
		clientConnection: make(map[string]interface{}),
	}
}

// -----------------------------------------------------------------------------
// service interface
func (s *GatewayService) Install() error {
	glog.V(1).Infof(">>> gateway::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &gatewayModel.Config); err != nil {
		glog.V(1).Infof("(-) gateway::Initialize - Error: %+v", err)
		return err
	}
	glog.V(1).Infoln("(+) load configuration for gateway successfully!")

	// 2. Install DAO
	gatewayModel.ManagerDao.ConnectDB(gatewayModel.Config.DBConfig, env.PostgresType)
	glog.V(1).Infoln("(+) Install Database successfully!")

	// 3. Install gRPC client
	gatewayModel.DocsClient = doc_client.NewDocumentClient()
	gatewayModel.BizClient = biz_client.NewBizClient()

	// 4. Install gRPC server

	// 5. Install HTTP server
	s.apiServer = &http.Server{
		Addr:         gatewayModel.Config.Addr,
		Handler:      Router(gatewayModel),
		ReadTimeout:  time.Duration(gatewayModel.Config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(gatewayModel.Config.WriteTimeout) * time.Second,
	}

	return nil
}

func (s *GatewayService) Start() {
	go func() {
		err := s.apiServer.ListenAndServe()
		if err != nil {
			glog.Fatalf("gateway::ListenAndServer - Error: %+v", err)
		}
	}()
}

func (s *GatewayService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
