package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/database"
	"github.com/MarskTM/financial_report_server/infrastructure/system"
	"github.com/MarskTM/financial_report_server/services/gateway/internal/rpc"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/go-chi/jwtauth"
)

var (
	config         env.GatewayConfig
	configDocument env.DocumentConfig
	managerDao     database.ManagerDBDao
	DecodeAuth     *jwtauth.JWTAuth
	EncodeAuth     *jwtauth.JWTAuth
)

// ----------------------------------------------------------------
type GatewayService struct {
	apiServer *http.Server
	server    *grpc.Server
	impl      *rpc.GatewayImpl

	// ----------------------------------------------------------------
	documentClient *grpc.ClientConn
}

// Constructor creates a new GatewayServer
func NewGatewayService() system.ServicesInterface {
	return &GatewayService{}
}

// -----------------------------------------------------------------------------
// service interface
func (s *GatewayService) Install() error {

	glog.V(3).Infof("gateway::Initialize ..!")
	// -----------------------------------------------------------------------------
	// 1. Install configuration
	err := utils.LoadConfig(&config)
	if err != nil {
		glog.V(1).Infof("gateway::Initialize - Error: %+v", err)
		return err
	}

	// 2. Install DAO
	managerDao.ConnectDB(*config.DB, system.PostgresDB)

	// 4. Install server
	s.apiServer = &http.Server{
		Addr:         config.Addr,
		Handler:      Router(),
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
	}

	// 3. Install gRPC client
	conn, err := grpc.Dial(configDocument.URL, grpc.WithInsecure())
	if err != nil {
		glog.Error("Failed to connect: %v", err)
	}
	defer conn.Close()
	// s.documentClient = pb.NewDocumentClient(document.)

	// 4. Install gRPC server


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
