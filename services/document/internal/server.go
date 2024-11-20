package internal

import (
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	"github.com/MarskTM/financial_report_server/infrastructure/proto/pb"
	"github.com/MarskTM/financial_report_server/infrastructure/system"

	// biz_client "github.com/MarskTM/financial_report_server/services/biz_server/client"
	"github.com/MarskTM/financial_report_server/services/document/internal/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var docsModel model.DocumentModel

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

	glog.V(3).Infof("Document_server::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	if _, err := toml.DecodeFile("./config.toml", &docsModel.Config); err != nil {
		glog.V(1).Infof("(-) gateway::Initialize - Error: %+v", err)
		return err
	}

	glog.V(1).Infof("(+) load configuration successfully!")
	// 2. Install DAO
	docsModel.DB.ConnectDB(docsModel.Config.DBConfig, env.PostgresType)

	// 3. Install gRPC client
	// docsModel.BizClient = biz_client.NewBizClient()
	glog.V(1).Infof("(+) RegisterClient successfully!")

	// 4. Install gRPC server
	pb.RegisterDocumentServer(s.server, rpc.NewDocsService(docsModel))
	glog.V(1).Infof("(+) RegisterServer successfully!")

	return nil
}

func (s *DocumentService) Start() {
	glog.V(1).Infof("Document_server::Start listening on %s", docsModel.Config.Addr)
	go func() {
		lis, err := net.Listen("tcp", docsModel.Config.Addr)
		if err != nil {
			glog.Errorf("Failed to listen: %v", err)
		}

		if err := s.server.Serve(lis); err != nil {
			glog.Errorf("Failed to serve gRPC Document: %v", err)
		}
	}()
	
	glog.V(1).Infof("Document_server::Started cron job")
	go func ()  {
		
	}()
}

func (s *DocumentService) Shutdown(signals chan os.Signal) {
	<-signals
	s.server.GracefulStop()
}
