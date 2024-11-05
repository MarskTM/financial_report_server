package internal

import (
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	config     env.DocumentConfig
	managerDao database.ManagerDAO
)

// ----------------------------------------------------------------
type DocumentService struct {
	server *grpc.Server

	// ----------------------------------------------------------------
}

// Constructor creates a new GatewayServer
func NewDocumentService() system.ServicesInterface {
	return &DocumentService{
		server: grpc.NewServer(),
	}
}

// -----------------------------------------------------------------------------
// service interface
func (s *DocumentService) Install() error {

	glog.V(3).Infof("gateway::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		glog.V(1).Infof("(-) gateway::Initialize - Error: %+v", err)
		return err
	}

	glog.V(1).Infof("(+) load configuration successfully!")
	// 2. Install DAO
	managerDao.ConnectDB(config.DBConfig, env.PostgresType)

	// 3. Install gRPC client

	// 4. Install gRPC server
	// pb.RegisterDocumentServer(s.server,)

	return nil
}

func (s *DocumentService) Start() {
	go func() {
		lis, err := net.Listen("tcp", config.Addr)
		if err != nil {
			glog.Errorf("Failed to listen: %v", err)
		}

		if err := s.server.Serve(lis); err != nil {
			glog.Errorf("Failed to serve gRPC Document: %v", err)
		}
	}()
}

func (s *DocumentService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
