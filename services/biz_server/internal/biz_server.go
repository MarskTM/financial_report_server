package internal

import (
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/biz_server/internal/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"

	doc_client "github.com/MarskTM/financial_report_server/services/document/client"
)

var (
	config     env.BizServerConfig
	managerDao database.ManagerDBDao
)

// ----------------------------------------------------------------
type BizService struct {
	server *grpc.Server

	// ----------------------------------------------------------------
	clientConnection map[string]interface{}
}

// Constructor creates a new GatewayServer
func NeBizService() system.ServicesInterface {
	return &BizService{
		server: grpc.NewServer(),
	}
}

// -----------------------------------------------------------------------------
// service interface
func (s *BizService) Install() error {

	glog.V(3).Infof("gateway::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		glog.V(1).Infof("(-) gateway::Initialize - Error: %+v", err)
		return err
	}
	glog.V(1).Infof("(+) load configuration successfully!")

	// 2. Install DAO
	managerDao.ConnectDB(*config.DB, system.PostgresDB)

	// 3. Install gRPC client
	s.clientConnection["document"] = doc_client.NewDocumentClient()

	// 4. Install gRPC server
	pb.RegisterBizServiceServer(s.server, rpc.NewBizService())

	return nil
}

func (s *BizService) Start() {
	go func() {
		lis, err := net.Listen("tcp", config.Addr)
		if err != nil {
			glog.Errorf("Failed to listen: %v", err)
		}

		if err := s.server.Serve(lis); err != nil {
			glog.Errorf("Failed to serve gRPC Biz server: %v", err)
		}
	}()
}

func (s *BizService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
