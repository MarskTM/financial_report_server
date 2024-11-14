package internal

import (
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/biz_server/internal/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	doc_client "github.com/MarskTM/financial_report_server/services/document/client"
)

var bizModel model.BizModel

// ----------------------------------------------------------------
type BizService struct {
	server *grpc.Server
}

// -----------------------------------------------------------------------------
// service interface
func (s *BizService) Install() error {

	glog.V(3).Infof("Biz_server::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &bizModel.Config); err != nil {
		glog.V(1).Infof("(-) Biz_server::Initialize - Error: %+v", err)
		return err
	}
	glog.V(1).Infof("(+) load configuration successfully!")

	// 2. Install DAO
	bizModel.DB.ConnectDB(bizModel.Config.DBConfig, env.PostgresType)
	glog.V(1).Infof("(+) Connected Database successfully!")

	// 3. Install gRPC client
	bizModel.DocsClient = doc_client.NewDocumentClient()
	glog.V(1).Infof("(+) RegisterClient successfully!")

	// 4. Install gRPC server
	pb.RegisterBizServiceServer(s.server, rpc.NewBizService(bizModel))
	glog.V(1).Infof("(+) RegisterServer successfully!")

	return nil
}

func (s *BizService) Start() {
	glog.V(1).Infof("Biz_server::Start listening on %s", bizModel.Config.Addr)
	go func() {
		lis, err := net.Listen("tcp", bizModel.Config.Addr)
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

// Constructor creates a new GatewayServer
func NewBizService() system.ServicesInterface {
	return &BizService{
		server: grpc.NewServer(),
	}
}
